package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"quotation/service"
	"quotation/storage"
)

func StartServer() {
	db, err := storage.Connect("Administrator", "sasha123")
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if cbData, err := service.FetchDataFromCentralBank(); err != nil {
			htmlResponse(w, "templates/error.html", struct {
				Message string
			}{"Error fetching data: " + err.Error()})
		} else {
			db.SaveToDatabase(cbData)
			htmlResponse(w, "templates/app.html", cbData)
		}
	})

	http.HandleFunc("/currencies", func(w http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("dt")
		if cbData, err := db.GetFromDatabase(date); err != nil {
			jsonResponse(w, "Incorrect date", nil)
		} else {
			jsonResponse(w, "OK", cbData)
		}
	})

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func jsonResponse(rw http.ResponseWriter, message string, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(struct {
		Message string
		Data    interface{}
	}{
		message,
		data,
	})
}

func htmlResponse(rw http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.ParseFiles(templateName)
	if err != nil {
		log.Fatal("Error render HTML template: ", err.Error())
	}
	tmpl.Execute(rw, data)
}
