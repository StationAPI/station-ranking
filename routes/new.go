package routes

import (
	"encoding/json"
	"net/http"
	"sort"

	neon "github.com/stationapi/station-ranking/db"
	"gorm.io/gorm"
)

type ByCreated []neon.Website

func (a ByCreated) Len() int { return len(a) }
func (a ByCreated) Less(i, j int) bool { return a[i].Created > a[i].Created }
func (a ByCreated) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func New(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
  websites := neon.ListWebsite(db)

  sort.Sort(ByCreated(websites))

  stringified, err := json.Marshal(websites[:199])

  if err != nil {
    http.Error(w, "there was an error fetching the new websites", http.StatusInternalServerError)

    return err
  }

  w.WriteHeader(200)
  w.Write(stringified)

  return nil
}
