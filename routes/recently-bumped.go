package routes

import (
	"net/http"

	"gorm.io/gorm"
)

func RecentlyBumped(w http.ResponseWriter, r *http.Request, db gorm.DB) {}
