# realescrape-go

Rewrite of my [Python based scraper](https://github.com/nronzel/realescrape)
in Go using the [Colly](https://github.com/gocolly/colly) web scraping framework.

> NOTE This is for educational purposes only, and serves as a way to learn scraping
> and create my own API to get the data on a local front end.

I don't intend on building a package to distribute as the nature of a web scraper
is to break, and it will likely need to be updated constantly.

## Upcoming

- [x] ~lotsize conversion~
- [x] ~ratios~
- [x] ~split utility functions into separate modules~
- [x] ~export to JSON~

more custom parameters:

- [ ] beds
- [ ] baths
- [ ] sqft
- [ ] max-price
- [ ] single/multi family

## Description

Simple web scraper for Realtor.com.

This program is intended to be run locally only and I do not intend to host
the scraper publicly as a service.

Currently has the following fixed search parameters:

- Single family homes
- Minimum Price - $100,000
- Minimum Bedrooms - 1
- Minimum Bathroom - 1
- Age - 3+
- Hides all houses pending sales
- Hides 55+ communities
- Sorts by newest listings

I will be implementing logic to make most of the search parameters dynamic.

Since the data is being stored in a CSV file you could always create filters
and sort data yourself once it is scraped.

A frontend will be developed as well that can be hosted locally to easily
view and filter through data.

## Installation

#### 1. Install Go

To install and run this scraper locally, you will first need to ensure you have
Go installed on your machine. You can download it and install following
the directions from the [official Golang website](https://go.dev/doc/install).

Be sure to follow the directions for your particular operating system.

#### 2. Clone this repository

Clone this repo to your local machine in whatever directory you choose.

```bash
git clone https://github.com/nronzel/realescrape-go
```

Navigate into the project directory

```bash
cd realescrape-go
```

#### 3. Install necessary dependencies

This project uses the Colly web scraping framework. To install the required
dependencies you can run the following command

```bash
go get -u
```

#### 4. Run the program

Run the program with the following command

```bash
go run . "Miami FL"
```

**Locations must be entered in the following formats:**

`"Miami FL"` - separate location and state, state must be capital

`"San-Francisco CA"` - use a hyphen for locations with spaces

`"90210"` - you can also just use a zip code

When the program is complete, you will see some stats in the console on how many
listings were scraped, how long it took, etc.

A CSV and JSON file will be generated and saved in the `scans` folder, located
at the root of the project directory. This folder will be created if it doesn't
already exist.

## Misc

Decided to export to JSON instead of porting over the MongoDB functionality from
the Python version. For the purposes of just displaying the data and some
visualizations locally I think a simple [JSON server](https://github.com/typicode/json-server)
would suffice.

I will have to add some logic to merge any JSON files present in the `scans`
directory to run the JSON server from that single file.

All filtering and visualizations will be done on the front end.

MongoDB functionality may be added to this re-write down the line once I get
some more of the basic features done.

## Issues

If you run into any problems you can open an issue, or submit a pull request.

Currently there are some times where the program will run, appear to scrape,
but the stats at the end will show 0 listings scraped. In this situation just
try running the same scrape again as it will work as long as the location entered
is valid.

I will likely implement some logic to check if the listings scraped is 0
and re-run the scrape if that is the case.
