package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type product struct {
	Id            int
	Name          string
	PathImage     string
	Price         float64
	ReduceProcent float64
	HasReduc      bool
}

var listTemplats = []product{
	{Id: 1, Name: "Produit 1", PathImage: "/static/icons/Logo/1.png", Price: 19.99, ReduceProcent: 10, HasReduc: true},
	{Id: 2, Name: "Produit 2", PathImage: "/static/icons/Logo/2.png", Price: 29.99, ReduceProcent: 0, HasReduc: false},
	{Id: 3, Name: "Produit 3", PathImage: "/static/icons/Logo/3.png", Price: 39.99, ReduceProcent: 15, HasReduc: true},
}

func main() {
	tmpl, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println("Erreur template:", err)
		os.Exit(1)
	}

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "../asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", listTemplats)
	})

	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println("Erreur serveur:", err)
		os.Exit(1)
	}
}
