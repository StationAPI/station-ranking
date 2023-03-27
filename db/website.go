package db

import "gorm.io/gorm"

type Website struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	IconURL     string   `json:"icon_url"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
  LastBumped  int `json:"last_bumped"`
  Created     int64 `json:"created"`
  Bumps       int `json:"bumps"`
  Upvotes int `json:"upvotes"`
  UpvotesToday int `json:"upvotes_today"`
}

func ListWebsite(db gorm.DB) []Website {
	var websites []Website

	db.Find(websites)

	return websites
}
