# realescrape-go

Rewrite of my [Python based scraper](https://github.com/nronzel/realescrape)
in Go using the [Colly](https://github.com/gocolly/colly) web scraping framework.

> NOTE This is for educational purposes only, and serves as a way to learn scraping
> and practice creating an API that can be consumed by a front end.

I don't intend on building a package to distribute as the nature of a web scraper
is to break, and it will likely need to be updated constantly.

## Upcoming

- [x] ~lotsize conversion~
- [x] ~ratios~
- [x] ~split utility functions into separate modules~
- [x] ~export to JSON~
- [x] ~combine json files into a master file with all data~

more custom parameters:

- [ ] beds
- [ ] baths
- [ ] sqft
- [ ] max-price
- [ ] single/multi family
- [x] radius

TODO:
- [ ] better error handling

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

#### 3. Run the program

Run the program with the following command

```bash
go run . "Miami FL"
```

The necessary dependencies should be automatically installed based on the go.mod
file.

**Locations must be entered in the following formats:**

`"Miami FL"` - separate location and state, state must be capital

`"San-Francisco CA"` - use a hyphen for locations with spaces

`"90210"` - you can also just use a zip code

> Currently will not work for full state searches (e.g. "Florida"). The url can't
take a radius parameter. This may be remedied in the future by taking an argument
flag at the console.

When the program is complete, you will see some stats in the console on how many
listings were scraped, how long it took, etc.

A CSV and JSON file will be generated and saved in the `scans` folder, located
at the root of the project directory. This folder will be created if it doesn't
already exist.

## DevLog

~Decided to export to JSON instead of porting over the MongoDB functionality from
the Python version. For the purposes of just displaying the data and some
visualizations locally I think a simple [JSON server](https://github.com/typicode/json-server)
would suffice.~

I realized using JSON Server would prevent me from being able to make my own API.

So instead it is back to the original idea of using a local MongoDB instance.

### Issues

#### Known Issues

zip code search will pull listings nearby if there are a small number of
results for the searched zip. Not really a big problem though.

~will sometimes randomly not return data, but appear to scan each page.
I am not sure what causes this, however you can just run the program again
and it will work. A quick check at the console print of the number of listings
scanned should tell you if it worked or not.~

I believe the above issue is fixed by applying a min radius of 1 mile. I made
radius a global const so it can be easily modified.

---

If you run into any problems you can open an issue, or submit a pull request.
