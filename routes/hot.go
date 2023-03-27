package routes

import (
	"encoding/json"
	"net/http"
	"sort"

	neon "github.com/stationapi/station-ranking/db"
	"gorm.io/gorm"
)

type ByUpvotes []neon.Website

func (a ByUpvotes) Len() int { return len(a) }
func (a ByUpvotes) Less(i, j int) bool { return a[i].UpvotesToday < a[j].UpvotesToday }
func (a ByUpvotes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Hot(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
  websites := neon.ListWebsite(db)

  sort.Sort(ByUpvotes(websites))

  stringified, err := json.Marshal(websites[:199])

  if err != nil {
    http.Error(w, "there was an error fetching the hot websites", http.StatusInternalServerError)

    return err
  }

  w.WriteHeader(200)
  w.Write(stringified)

  return nil
}
