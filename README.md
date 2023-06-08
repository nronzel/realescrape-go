# realescrape-go ![go](https://img.shields.io/github/languages/top/nronzel/realescrape-go?style=flat-square) ![go-version](https://img.shields.io/badge/Go-v1.20-blue)

Rewrite of my [Python based scraper](https://github.com/nronzel/realescrape)
in Go using the [Colly](https://github.com/gocolly/colly) web scraping framework.

> NOTE This project is for educational purposes only. Please use with care.



## Upcoming

TODO:

- [x] ~lotsize conversion~
- [x] ~ratios~
- [x] ~split utility functions into separate modules~
- [x] ~export to JSON~
- [x] ~combine json files into a master file with all data~
- [x] ~MongoDB~
- [x] ~API endpoint~
- [ ] Unit tests for the DB and API
- [x] ~Make API a little more robust~
- [ ] Websockets for updates
- [ ] Front End
- [ ] Dockerize the app
- [x] ~Split code into separate packages for easier maintanability~

## Description

Simple web scraper for Realtor.com.

This program is intended to be run locally only.
I do intend on creating Docker files so you can just run this with Docker.

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

#### API

Uses the [Echo](https://echo.labstack.com) framework for the API.

Currently has a GET endpoint that retrieves all items in the MongoDB
collection.

I will be implementing more endpoints in the coming weeks and may also put
some basic authentication in just as a basic layer of security and as a way
to learn how to even do so as this is my first API. Currently the auth is not
a priority.

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

You should also be able to just run the following command and it should install
the dependencies automatically.

```bash
go run main.go
```

#### 4. Run the program

Run the program with the following command

```bash
go run main.go
```
When you run `main.go` it will connect to the MongoDB and start the API
on `localhost:3000`. You can then go into the `frontend/realescrape/` folder
and run the command `pnpm run start` or `npm run start` and it will start the
frontend on `localhost:3100`. Visit `localhost:3100` in your browser to
interact with the database.

**Locations must be entered in the following formats:**

`"Miami FL"` - separate location and state, state must be capital

`"San-Francisco CA"` - use a hyphen for locations with spaces

`"90210"` - you can also just use a zip code

> Currently will not work for full state searches (e.g. "Florida"). The url can't
> take a radius parameter. This will be remedied in the future.

When the program is complete, you will see some stats in the console on how many
listings were scraped, how long it took, etc.

### Issues

#### Known Issues

Need to set up websockets so when the database updates it can update the
frontend properly. For now, a manual refresh button will do.

---

If you run into any problems you can open an issue, or submit a pull request.
