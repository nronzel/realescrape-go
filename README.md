# realescrape-go ![go](https://img.shields.io/github/languages/top/nronzel/realescrape-go?style=flat-square)

Rewrite of my [Python based scraper](https://github.com/nronzel/realescrape)
in Go using the [Colly](https://github.com/gocolly/colly) web scraping framework.

> NOTE This project is for educational purposes only. Please use with care.

## Upcoming

- [x] ~lotsize conversion~
- [x] ~ratios~
- [x] ~split utility functions into separate modules~
- [x] ~export to JSON~
- [x] ~combine json files into a master file with all data~

-- These will be implemented once I create a proper CLI for this program.

Custom parameters:
- [ ] beds
- [ ] baths
- [ ] sqft
- [ ] max-price
- [ ] single/multi family
- [x] radius

TODO:
- [x] ~MongoDB~
- [x] ~API endpoint~
- [ ] Unit tests for the DB and API
- [ ] Front End
- [x] ~Split code into separate packages for easier maintanability~
- [ ] Create a proper CLI to set the parameters rather than taking arguments

## Description

Simple web scraper for Realtor.com.

This program is intended to be run locally only.

It also requires a MongoDB instance to be running locally on the default port
with no credentials.

Currently has the following fixed search parameters:

- Single family homes
- Minimum Price - $100,000
- Minimum Bedrooms - 1
- Minimum Bathroom - 1
- Age - 3+
- Hides all houses pending sales
- Hides 55+ communities
- Hides foreclosures
- Sorts by newest listings

I will be implementing logic to make most of the search parameters dynamic.

Since the data is being stored in a CSV file you could always create filters
and sort data yourself once it is scraped.

A frontend will be developed as well that can be hosted locally to easily
view and filter through data.

##### API

The API is a very simple Echo server with a single GET endpoint to retrieve
all of the houses in the Mongo collection.

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

#### 3. Install dependencies

In the root directory of the project, run
```bash
go get .
```

#### 4. Run the program

Run the program with the following command, with the location you are interested
in as an argument.

```bash
go run main.go "Miami FL"
```
Running `main.go` will connect to the MongoDB collection, perform the data
extraction, merge all JSON documents in the /scans directory into a `master.json`
file, uploads all of the entries in `master.json` to the MongoDB collection,
and then starts the API. The endpoint is accessible at `localhost:3000/houses`.

**Locations must be entered in the following formats:**

`"Miami FL"` - separate location and state, state must be capital

`"San-Francisco CA"` - use a hyphen for locations with spaces

`"90210"` - you can also just use a zip code

> Currently will not work for full state searches (e.g. "Florida"). The url can't
take a radius parameter. This will be remedied in the future when it is a proper
CLI.

When the program is complete, you will see some stats in the console on how many
listings were scraped, how long it took, etc.

A CSV and JSON file will be generated and saved in the `scans` folder, located
at the root of the project directory. This folder will be created if it doesn't
already exist.

## DevLog

Implemented MongoDB. Uses the URL of each listing as the unique key. More work needs
to be done to handle updating existing entries, currently it will just error out
if there is a duplicate entry.

### Issues

#### Known Issues

~will sometimes randomly not return data, but appear to scan each page.
I am not sure what causes this, however you can just run the program again
and it will work. A quick check at the console print of the number of listings
scanned should tell you if it worked or not.~

I believe the above issue is fixed by applying a min radius of 1 mile. I made
radius a global const so it can be easily modified.

---

If you run into any problems you can open an issue, or submit a pull request.
