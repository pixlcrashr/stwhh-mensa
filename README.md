# stwhh-mensa

A small and simple program to crawl the daily menu data from https://www.stwhh.de/speiseplan.

The goal is to provide a good data over time and an API/website to retrieve data over longer periods of time.

## Getting started

### Download

Simply just go to the [Releases](https://www.stwhh.de/speiseplan?t=today) section and download your desired version. Atm, only Linux builds are available.

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
- [ ] Add automatic crawler (daily interval crawling)
- [ ] Update database structure (e.g. normalizing the DB schema)
- [ ] Add a GraphQL API
- [ ] Add a simple Frontend interface
