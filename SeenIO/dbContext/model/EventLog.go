package model

import "github.com/jinzhu/gorm"

type EventLog struct {
	gorm.Model

	LandingPageHits int

	VideoPlays int

	UserID int
}
