package database

import (
	"time"
)


type VisitedLink struct {
	Website string `bson:"website"`
	Link string		`bson:"link"`
	VisitedDate time.Time	`bson:"visite_date"`
}