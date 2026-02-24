package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WeatherResponse struct {
	Data map[string]interface{} `json:"data"`
}

func fetchWeather() (map[string]interface{}, error) {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=55.67&longitude=12.56&current_weather=true")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

// @title WhoKnows
// @version 0.1.0
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Serve static files
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	// @Summary Serve Root Page
	// @Router / [get]
	// @Success 200 {string} string "text/html"
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>WhoKnows</h1>"))
	}).Methods("GET")

	// @Summary Serve Weather Page
	// @Router /weather [get]
	// @Success 200 {string} string "text/html"
	r.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		data, err := fetchWeather()
		if err != nil {
			http.Error(w, "Could not fetch weather", http.StatusInternalServerError)
			return
		}
		current := data["current_weather"].(map[string]interface{})
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Weather in Copenhagen</h1><p>Temperature: " +
			fmt.Sprintf("%.1f", current["temperature"]) + "°C</p><p>Windspeed: " +
			fmt.Sprintf("%.1f", current["windspeed"]) + " km/h</p>"))
	}).Methods("GET")

	// @Summary Serve Register Page
	// @Router /register [get]
	// @Success 200 {string} string "text/html"
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Register</h1>"))
	}).Methods("GET")

	// @Summary Serve Login Page
	// @Router /login [get]
	// @Success 200 {string} string "text/html"
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<h1>Login</h1>"))
	}).Methods("GET")

	// API routes

	// @Summary Search
	// @Router /api/search [get]
	// @Param q query string true "Search query"
	// @Param language query string false "Language code (e.g. 'en')"
	// @Success 200 {object} map[string]interface{}
	r.HandleFunc("/api/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []interface{}{},
		})
	}).Methods("GET")

	// @Summary Weather
	// @Router /api/weather [get]
	// @Success 200 {object} map[string]interface{}
	r.HandleFunc("/api/weather", func(w http.ResponseWriter, r *http.Request) {
		data, err := fetchWeather()
		if err != nil {
			http.Error(w, "Could not fetch weather", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(WeatherResponse{Data: data})
	}).Methods("GET")

	// @Summary Register
	// @Router /api/register [post]
	// @Param username formData string true "Username"
	// @Param email formData string true "Email"
	// @Param password formData string true "Password"
	// @Param password2 formData string true "Confirm Password"
	// @Success 200 {object} map[string]interface{}
	r.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 200,
			"message":    "OK",
		})
	}).Methods("POST")

	// @Summary Login
	// @Router /api/login [post]
	// @Param username formData string true "Username"
	// @Param password formData string true "Password"
	// @Success 200 {object} map[string]interface{}
	r.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 200,
			"message":    "OK",
		})
	}).Methods("POST")

	// @Summary Logout
	// @Router /api/logout [get]
	// @Success 200 {object} map[string]interface{}
	r.HandleFunc("/api/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"statusCode": 200,
			"message":    "OK",
		})
	}).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
