package main

import (
	"io"
	"net/http"
	"html/template"
)

type BoardPageContext struct {
	Title        string
}

type HomePageContext struct {
	Title        string
}

type ControlPageContext struct {
	Title        string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/_base.tmpl.html",
		"templates/layouts/main.tmpl.html",
		"templates/pages/home.tmpl.html", 
	)
	if err != nil {
		panic(err.Error())
	}

	tmplCtx := HomePageContext{
		Title: "Workout Board",
	}
	
	err = tmpl.Execute(w, tmplCtx)
	if err != nil {
		panic(err.Error())
	}

}

func BoardPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/_base.tmpl.html",
		"templates/layouts/main.tmpl.html",
		"templates/pages/board.tmpl.html",
	)
	if err != nil {
		panic(err.Error())
	}
	
	tmplCtx := BoardPageContext{
		Title: "Workout Board",
	}

	err = tmpl.Execute(w, tmplCtx)
	if err != nil {
		panic(err.Error())
	}

}

func ControlPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(
		"templates/_base.tmpl.html",
		"templates/layouts/main.tmpl.html",
		"templates/pages/control.tmpl.html",
	)
	if err != nil {
		panic(err.Error())
	}
	
	tmplCtx := ControlPageContext{
		Title: "Control",
	}

	err = tmpl.Execute(w, tmplCtx)
	if err != nil {
		panic(err.Error())
	}

}

func APIRoomsHandler(w http.ResponseWriter, r *http.Request) {
	
	req, err := http.NewRequest(r.Method, "http://node:3000/api/rooms", nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header = r.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {

	http.HandleFunc("/home", HomePageHandler)
	http.HandleFunc("/board", BoardPageHandler)
	http.HandleFunc("/control", ControlPageHandler)

	http.HandleFunc("/api/rooms", APIRoomsHandler)

	_ = http.ListenAndServe(":3333", nil)
}
