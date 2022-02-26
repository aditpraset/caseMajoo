package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type Pagination struct {
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
	Sort         string `json:"sort"`
	TotalRows    int    `json:"total_rows"`
	FirstPage    string `json:"first_page"`
	PreviousPage string `json:"previous_page"`
	NextPage     string `json:"next_page"`
	LastPage     int    `json:"last_page"`
	FromRow      int    `json:"from_row"`
	ToRow        int    `json:"to_row"`
	Dates        []Date `json:"date"`
}

type ResponseJson struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Paging  *Pagination
	Data    interface{} `json:"data"`
}

type Search struct {
	Column string `json:"column"`
	Action string `json:"action"`
	Query  string `json:"query"`
}

type Date struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

func Paging(r *http.Request) *Pagination {
	var data Pagination
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&data); err != nil {
				log.Fatal(err)
			}
		}
	}
	var limit, page int
	page = data.Page
	limit = data.Limit
	if data.Page == 0 {
		page = 1
	}
	if data.Limit == 0 {
		limit = 25
	}
	if page < 1 {
		page = 1
	}
	begin := (limit * page) - limit
	return &Pagination{Limit: limit, Page: page, FromRow: begin, Dates: data.Dates}
}
