# stwhh-mensa

![GitHub Actions Status](https://github.com/pixlcrashr/stwhh-mensa/actions/workflows/build.yaml/badge.svg?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/pixlcrashr/stwhh-mensa.svg)](https://pkg.go.dev/github.com/pixlcrashr/stwhh-mensa)
[![Go Report Card](https://goreportcard.com/badge/github.com/pixlcrashr/stwhh-mensa)](https://goreportcard.com/report/github.com/pixlcrashr/stwhh-mensa)

A small and simple program to crawl the daily menu data from https://www.stwhh.de/speiseplan.

The goal is to provide STWHH food data over time and an API/website to retrieve current and past data over longer periods of time. Comparing prices over time is especially one topic of high interest.

## Getting started

### Download

Simply just go to the [Releases](https://github.com/pixlcrashr/stwhh-mensa/releases) section and download your desired version. Atm, only Linux builds are available.

### Docker

To use the Docker image, run:

```sh
docker run -it -v <local-data-folder>:/opt/app/data themysteriousvincent/stwhh-mensa:latest crawler --db-path /opt/app/data/db.sqlite
```

### Building

To build the program, simply install Go 1.22.5 and run the following:

```shell
go mod tidy
go build -o stwhh-mensa ./main.go
```

Then, you can run the built binary like so:

```shell
./stwhh-mensa
```

### CLI

The program provides a CLI interface, which should help you how to use the program.

Simply type

```shell
stwhh-mensa -h
```

and the CLI-Help will be shown.


## Roadmap

- [x] Add basic crawler
- [x] Add automatic crawler (periodically interval crawling)
- [x] Update database structure (e.g. normalizing the DB schema)
- [ ] Add price same-day/same-week price changes
- [ ] Add a GraphQL API
- [ ] Add a simple Frontend interface
