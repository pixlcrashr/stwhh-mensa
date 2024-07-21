package crawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pixlcrashr/stwhh-mensa/pkg/model"
	"github.com/pixlcrashr/stwhh-mensa/pkg/nullable"
	slices2 "github.com/pixlcrashr/stwhh-mensa/pkg/slices"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const STWHHUrl = "https://www.stwhh.de/speiseplan?t=this_week"

type Crawler struct {
	httpClient *http.Client
}

func NewCrawler() *Crawler {
	return &Crawler{
		httpClient: &http.Client{},
	}
}

func (c *Crawler) Crawl(ctx context.Context) ([]Day, error) {
	resp, err := c.httpClient.Get(STWHHUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code received")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	days, err := parseDays(doc)
	if err != nil {
		return nil, err
	}

	return days, nil
}

func parseDays(document *goquery.Document) ([]Day, error) {
	locSel := document.Find(".tx-epwerkmenu-content .tx-epwerkmenu-menu-location-container:not(.d-none)")

	var days = make([]Day, 0)

	locSel.Each(func(i int, selection *goquery.Selection) {
		loc, err := parseLocation(selection)
		if err != nil {
			return
		}

		categories, err := parseCategoriesWithDate(selection)
		if err != nil {
			return
		}

		for _, category := range categories {
			days = slices2.AddOrSet(
				days,
				func(d Day) bool {
					return d.Date == category.Date
				},
				func(d Day) Day {
					d.Gastronomies = slices2.AddOrSet(
						d.Gastronomies,
						func(g model.Gastronomy) bool {
							return g.ID == loc.ID
						},
						func(g model.Gastronomy) model.Gastronomy {
							g.Categories = slices2.AddOrSet(
								g.Categories,
								func(c model.Category) bool {
									return c.ID == category.Category.ID
								},
								func(c model.Category) model.Category {
									return c
								},
								func() model.Category {
									return category.Category
								},
							)
							return g
						},
						func() model.Gastronomy {
							return model.Gastronomy{
								ID:       loc.ID,
								Name:     loc.Name,
								Location: loc.Location,
								Categories: []model.Category{
									category.Category,
								},
							}
						},
					)
					return d
				},
				func() Day {
					return Day{
						Date: category.Date,
						Gastronomies: []model.Gastronomy{
							{
								ID:       loc.ID,
								Name:     loc.Name,
								Location: loc.Location,
								Categories: []model.Category{
									category.Category,
								},
							},
						},
					}
				},
			)
		}
	})

	return days, nil
}

func parseLocation(selection *goquery.Selection) (model.Gastronomy, error) {
	var g model.Gastronomy

	idStr, ok := selection.Attr("data-location-id")
	if !ok {
		return g, errors.New("gastronomy id could not be parsed")
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return g, err
	}

	nameSel := selection.Find(".col-10 .offset-1 .mensainfo .row .col-12 .mensainfo__title").First()
	locSel := selection.Find(".col-10 .offset-1 .mensainfo .row .col-12 .mensainfo__subtitle").First()

	name := nameSel.Text()
	loc := locSel.Text()

	g.ID = int(id)
	g.Name = name
	g.Location = loc
	g.Categories = []model.Category{}

	return g, nil
}

func parseCategoriesWithDate(selection *goquery.Selection) ([]struct {
	Date     time.Time
	Category model.Category
}, error) {
	catSel := selection.Find(".row .col-12 .tx-epwerkmenu-menu-location-wrapper .tx-epwerkmenu-menu-locationpart-wrapper .tx-epwerkmenu-menu-times-wrapper .tx-epwerkmenu-menu-timestamp-wrapper")

	res := make([]struct {
		Date     time.Time
		Category model.Category
	}, 0)

	catSel.Each(func(i int, selection *goquery.Selection) {
		r, err := parseCategoryWithDate(selection)
		if err != nil {
			return
		}

		res = append(res, r)
	})

	return res, nil
}

func parseCategoryWithDate(selection *goquery.Selection) (struct {
	Date     time.Time
	Category model.Category
}, error) {
	var res struct {
		Date     time.Time
		Category model.Category
	}

	dateStr, ok := selection.Attr("data-timestamp")
	if !ok {
		return res, errors.New("attribute data-timestamp does not exist")
	}

	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return res, err
	}

	res.Date = date

	subSel := selection.Find(".row .col-12").First()

	var category model.Category

	catNameSel := subSel.Find(".container-fluid .row .col-10 .menulist__categorytitle").First()
	name := strings.Trim(catNameSel.Text(), " \n\r")

	category.Name = name

	dishSel := subSel.Find(".menulist__mealswrapper .container-fluid .row .col-12 .row .menue-tile")

	dishes := make([]model.Dish, 0)

	dishSel.Each(func(i int, selection *goquery.Selection) {
		dish, err := parseDish(dishSel)
		if err != nil {
			return
		}

		category.ID = dish.CategoryIDs[0]

		dishes = append(dishes, dish)
	})

	category.Dishes = dishes
	res.Category = category

	return res, nil
}

func parseDish(selection *goquery.Selection) (model.Dish, error) {
	id, err := parseDishID(selection)
	if err != nil {
		return model.Dish{}, err
	}

	name, err := parseDishName(selection)
	if err != nil {
		return model.Dish{}, err
	}

	categoryIDs, err := parseDishCategoryIDs(selection)
	if err != nil {
		return model.Dish{}, err
	}

	prices, err := parseDishPrice(selection, id)
	if err != nil {
		return model.Dish{}, err
	}

	allergens, err := parseDishAllergens(selection)
	if err != nil {
		return model.Dish{}, err
	}

	symbolIDs, err := parseDishSymbolIDs(selection)
	if err != nil {
		return model.Dish{}, err
	}

	return model.Dish{
		ID:          id,
		Name:        name,
		CategoryIDs: categoryIDs,
		Prices:      prices,
		Allergens:   allergens,
		SymbolIDs:   symbolIDs,
	}, nil
}

func parseDishID(selection *goquery.Selection) (int, error) {
	idStr, ok := selection.Attr("data-uid")
	if !ok {
		return 0, errors.New("id attribute does not exist")
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func parseDishCategoryIDs(selection *goquery.Selection) ([]int, error) {
	categoryIDsStr, ok := selection.Attr("data-categories")
	if !ok {
		return nil, errors.New("categories attribute does not exist")
	}

	categoryIDSs := strings.Split(
		strings.Trim(categoryIDsStr, " "),
		" ",
	)
	categoryIDs := make([]int, 0)
	for _, s := range categoryIDSs {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}

		categoryIDs = append(categoryIDs, int(v))
	}

	return categoryIDs, nil
}

func parseDishName(selection *goquery.Selection) (string, error) {
	title := selection.Find(".singlemeal .singlemeal__top .row .col-12 .singlemeal__headline").First().Text()

	title = strings.Trim(title, " \n\r")
	parts := strings.Split(title, "\n")

	var resultParts []string
	for _, part := range parts {
		s := strings.Trim(part, " \n\r")
		if len(s) == 0 {
			continue
		}

		resultParts = append(resultParts, s)
	}

	return strings.Join(resultParts, ", "), nil
}

func parseDishAllergens(selection *goquery.Selection) ([]string, error) {
	allergensStr, ok := selection.Attr("data-allergens")
	if !ok {
		return nil, errors.New("attribute data-allergens does not exist")
	}

	return strings.Split(
		strings.Trim(allergensStr, " \n\r"),
		" ",
	), nil
}

func parseDishSymbolIDs(selection *goquery.Selection) ([]int, error) {
	symbolIDsStr, ok := selection.Attr("data-symbols")
	if !ok {
		return nil, errors.New("attribute data-symbols does not exist")
	}

	if len(symbolIDsStr) == 0 {
		return []int{}, nil
	}

	symbolIDStrs := strings.Split(
		strings.Trim(symbolIDsStr, " \n\r"),
		" ",
	)

	symbolIDs := make([]int, 0)
	for _, s := range symbolIDStrs {
		id, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, err
		}

		symbolIDs = append(symbolIDs, int(id))
	}

	return symbolIDs, nil
}

func parseDishPrice(selection *goquery.Selection, id int) (p model.Prices, err error) {
	selection.Find(fmt.Sprintf("#textCollapse%d .singlemeal__bottom .row-custom-2 .col-12 .dlist:not(.dlist--inline) .dlist__item .singlemeal__info", id)).EachWithBreak(func(i int, selection *goquery.Selection) bool {
		typeStr := selection.Contents().Not(".singlemeal__info--semibold").Text()
		typeStr = strings.Trim(typeStr, " \n\r")

		valueStr := selection.Find(".singlemeal__info--semibold").Text()
		valueStr = strings.Replace(
			strings.Trim(valueStr, " —€\n\r"),
			",",
			"",
			1,
		)

		var value nullable.Nullable[int]

		if len(valueStr) > 0 {
			v, err := strconv.ParseInt(valueStr, 10, 32)
			if err != nil {
				return true
			}

			value = nullable.Value(int(v))
		}

		switch typeStr {
		case "Gäste":
			p.Guests = value
			break
		case "Studierende":
			p.Students = value
			break
		case "Bedienstete":
			p.Employees = value
			break
		}

		return true
	})

	return p, err
}
