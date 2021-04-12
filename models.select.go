package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Select struct {
	Id      int     `json:"id"`
	Year    int     `json:"year"`
	Name    string  `json:"Name"`
	Average float32 `json:"average"`
	Rate    int     `json:"rate"`
	Weight  float32 `json:"weight"`
	Minp    int     `json:"minp"`
	Maxp    int     `json:"maxp"`
	Time    int     `json:"time"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var SelectList = []Select{}

// Return a list of all the articles
func getSelect(condition string) ([]Select, error) {
	var SelectList_empty = []Select{}
	SelectList = SelectList_empty
	var select_temp = Select{}
	fmt.Println(condition)
	sql_cond := "Average DESC"
	i := 1
	db, err := sql.Open("mysql",
		"root:password@tcp(127.0.0.1:3306)/testDB")
	if err != nil {
		log.Fatal(err)
	}
	if condition == "popular" {
		sql_cond = "rate DESC"
	} else if condition == "hard" {
		sql_cond = "weight DESC"
	} else if condition == "easy" {
		sql_cond = "weight "
	} else if condition == "rate" {
		sql_cond = "Average DESC"
	} else {
		sql_cond = "Average DESC"
	}
	rows, err := db.Query("SELECT * FROM board ORDER BY " + sql_cond)
	if err != nil {
		// do something here
		log.Fatal(err)

	}
	if !rows.Next() {
		return nil, errors.New("Select not found")
	} else {
		err = rows.Scan(&select_temp.Id, &select_temp.Name, &select_temp.Average, &select_temp.Rate, &select_temp.Weight, &select_temp.Minp, &select_temp.Maxp, &select_temp.Time, &select_temp.Year)
		if err != nil {
			// do something here
			log.Fatal(err)
		}
		SelectList = append(SelectList, select_temp)
	}
	for rows.Next() {
		if i == 50 {
			break
		}
		err = rows.Scan(&select_temp.Id, &select_temp.Name, &select_temp.Average, &select_temp.Rate, &select_temp.Weight, &select_temp.Minp, &select_temp.Maxp, &select_temp.Time, &select_temp.Year)
		if err != nil {
			// do something here
			log.Fatal(err)

		}
		SelectList = append(SelectList, select_temp)
		i++

	}
	return SelectList, nil
}
