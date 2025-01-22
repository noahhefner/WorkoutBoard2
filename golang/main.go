package main

import (
	"os"
	"io"
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
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
	Workouts	[]Workout
}

type Workout struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Movements []Movement
}

type Movement struct {
	Name string `json:"name"`
	Duration int `json:"duration"`
}

var workouts []Workout

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
		"templates/layouts/board.tmpl.html",
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
		Workouts: workouts,
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

func GetWorkoutByID(w http.ResponseWriter, r *http.Request) {

	formVal := r.FormValue("id")
	if formVal == "" {
		http.Error(w, "No workout ID provided.", http.StatusBadRequest)
		return
	}

	workoutID, err := strconv.Atoi(formVal)
    if err != nil {
        http.Error(w, "Invalid workout ID.", http.StatusBadRequest)
		return
    }

	var workout Workout
	for _, w := range workouts {
		if w.ID == workoutID {
			workout = w
			break
		}
	}

	tmpl, err := template.ParseFiles(
		"templates/partials/workout.tmpl.html",
	)
	if err != nil {
		panic(err.Error())
	}

	err = tmpl.Execute(w, workout)
	if err != nil {
		panic(err.Error())
	}

}

func main() {

	data, err := os.ReadFile("./workouts.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to open workouts.json file: %s", err.Error()))
	}

	err = json.Unmarshal(data, &workouts)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal workouts: %s", err.Error()))
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/home", HomePageHandler)
	http.HandleFunc("/board", BoardPageHandler)
	http.HandleFunc("/control", ControlPageHandler)

	http.HandleFunc("/api/rooms", APIRoomsHandler)
	http.HandleFunc("/api/workout", GetWorkoutByID)

	_ = http.ListenAndServe(":3333", nil)
}
