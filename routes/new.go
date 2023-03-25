package routes

import (
	"net/http"

	"gorm.io/gorm"
)

func New(w http.ResponseWriter, r *http.Request, db gorm.DB) {}
