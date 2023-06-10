package models

import "time"

type House struct {
	Price     int       `json:"Price" bson:"price"`
	Beds      int       `json:"Beds" bson:"beds"`
	Baths     float64   `json:"Baths" bson:"baths"`
	Sqft      int       `json:"Sqft" bson:"sqft"`
	LotSize   float64   `json:"LotSize" bson:"lotSize"`
	LotUnit   string    `json:"LotUnit" bson:"lotUnit"`
	LotSqft   int       `json:"LotSqft" bson:"lotSqft"`
	Hty       float64   `json:"Hty" bson:"hty"`
	HtyPcnt   float64   `json:"HtyPcnt" bson:"htyPcnt"`
	Street    string    `json:"Street" bson:"street"`
	City      string    `json:"City" bson:"city"`
	State     string    `json:"State" bson:"state"`
	Zip       string    `json:"Zip" bson:"zip"`
	Link      string    `json:"Link" bson:"link"`
	CrawlTime time.Time `json:"CrawlTime" bson:"crawlTime"`
}
