package routes

import (
	"net/http"

	"gorm.io/gorm"
)

func Hot(w http.ResponseWriter, r *http.Request, db gorm.DB) {}
