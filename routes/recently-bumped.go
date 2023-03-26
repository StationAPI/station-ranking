package routes

import (
	"encoding/json"
	"net/http"
	"sort"

	neon "github.com/stationapi/station-ranking/db"
	"gorm.io/gorm"
)

type ByLastBumped []neon.Website

func (a ByLastBumped) Len() int { return len(a) }
func (a ByLastBumped) Less(i, j int) bool { return int64(a[i].LastBumped) > int64(a[i].LastBumped) }
func (a ByLastBumped) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func LastBumped(w http.ResponseWriter, r *http.Request, db gorm.DB) error {
  websites := neon.ListWebsite(db)

  sort.Sort(ByLastBumped(websites))

  stringified, err := json.Marshal(websites[:199])

  if err != nil {
    http.Error(w, "there was an error fetching the recently bumped websites", http.StatusInternalServerError)

    return err
  }

  w.WriteHeader(200)
  w.Write(stringified)

  return nil
}
