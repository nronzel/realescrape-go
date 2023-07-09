package models

import "time"

type House struct {
	Price     int       `json:"price" bson:"price"`
	Beds      int       `json:"beds" bson:"beds"`
	Baths     float64   `json:"baths" bson:"baths"`
	Sqft      int       `json:"sqft" bson:"sqft"`
	LotSize   float64   `json:"lotSize" bson:"lotSize"`
	LotUnit   string    `json:"lotUnit" bson:"lotUnit"`
	LotSqft   int       `json:"lotSqft" bson:"lotSqft"`
	Hty       float64   `json:"hty" bson:"hty"`
	HtyPcnt   float64   `json:"htyPcnt" bson:"htyPcnt"`
	Street    string    `json:"street" bson:"street"`
	City      string    `json:"city" bson:"city"`
	State     string    `json:"state" bson:"state"`
	Zip       string    `json:"zip" bson:"zip"`
	Link      string    `json:"link" bson:"link"`
	CrawlTime time.Time `json:"crawlTime" bson:"crawlTime"`
}
