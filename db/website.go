package db

import "gorm.io/gorm"

type Website struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	IconURL     string   `json:"icon_url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Created     int64
	Bumps       int
}

func ListWebsite(db gorm.DB) []Website {
	var websites []Website

	db.Find(websites)

	return websites
}
