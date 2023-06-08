package models

import "time"

type House struct {
	Price     string    `json:"Price" bson:"price"`
	Beds      string    `json:"Beds" bson:"beds"`
	Baths     string    `json:"Baths" bson:"baths"`
	Sqft      string    `json:"Sqft" bson:"sqft"`
	LotSize   string    `json:"LotSize" bson:"lotSize"`
	LotUnit   string    `json:"LotUnit" bson:"lotUnit"`
	LotSqft   string    `json:"LotSqft" bson:"lotSqft"`
	Hty       string    `json:"Hty" bson:"hty"`
	HtyPcnt   string    `json:"HtyPcnt" bson:"htyPcnt"`
	Street    string    `json:"Street" bson:"street"`
	City      string    `json:"City" bson:"city"`
	State     string    `json:"State" bson:"state"`
	Zip       string    `json:"Zip" bson:"zip"`
	Link      string    `json:"Link" bson:"link"`
	CrawlTime time.Time `json:"CrawlTime" bson:"crawlTime"`
}
