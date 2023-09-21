package controller

import (
	"encoding/json"
	"fmt"
	"laptop/database"
	"laptop/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func All(w http.ResponseWriter, r *http.Request) {
	db := database.Getconnect()
	defer db.Close()
	ss := []model.Model{}
	s := model.Model{}
	rows, err := db.Query("select * from blog")
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		for rows.Next() {
			rows.Scan(&s.ID, &s.Title, &s.Post)
			ss = append(ss, s)
		}
		json.NewEncoder(w).Encode(ss)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := database.Getconnect()
	defer db.Close()
	s := model.Model{}
	json.NewDecoder(r.Body).Decode(&s)
	i := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(i)
	result, err := db.Exec("update blog set title=?,post=?,where id=?", id, s.Title, s.Post)
	if err == nil {
		_, err := result.RowsAffected()
		if err != nil {
			json.NewEncoder(w).Encode("error")
		} else {
			json.NewEncoder(w).Encode(s)
		}
	} else {
		fmt.Fprint(w, err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	db := database.Getconnect()
	defer db.Close()
	i := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(i)
	result, err := db.Exec("delete from blog where id=?", id)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		_, err := result.RowsAffected()
		if err != nil {
			fmt.Print(w, err)
		} else {
			fmt.Fprint(w, "Suscesssful")
		}
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := database.Getconnect()
	defer db.Close()
	s := model.Model{}
	json.NewDecoder(r.Body).Decode(&s)
	result, err := db.Exec("insert into blog(title,post)value(?,?)", s.Title, s.Post)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		_, err := result.LastInsertId()
		if err != nil {
			json.NewEncoder(w).Encode("{No record is inserted}")
		} else {
			json.NewEncoder(w).Encode(s)

		}
	}
}
