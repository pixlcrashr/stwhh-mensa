package crawler

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pixlcrashr/stwhh-mensa/pkg/model"
	"github.com/pixlcrashr/stwhh-mensa/pkg/nullable"
	"net/http"
	"strconv"
	"strings"
)

const STWHHUrl = "https://www.stwhh.de/speiseplan?t=today"

type Crawler struct {
	httpClient *http.Client
}

func NewCrawler() *Crawler {
	return &Crawler{
		httpClient: &http.Client{},
	}
}

func (c *Crawler) Crawl(ctx context.Context) (Result, error) {
	var res Result

	resp, err := c.httpClient.Get(STWHHUrl)
	if err != nil {
		return res, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return res, errors.New("invalid status code received")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return res, err
	}

	ds := make([]model.Dish, 0)

	doc.Find(".menue-tile").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		d, err := parseDish(selection)
		if err != nil {
			return false
		}

		ds = append(ds, d)
		return true
	})

	res.Dishes = ds

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
	title := selection.Find(".singlemeal .singlemeal__top .row .col-12 .singlemeal__headline").Text()

	title = strings.Trim(title, " \n\r")

	return title, nil
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
