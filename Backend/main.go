package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/melzareix/MrMeeseeksBot/Database"
	"net/http"
	"os"
	"github.com/melzareix/MrMeeseeksBot/Api"
	"github.com/melzareix/MrMeeseeksBot/Models"
	"github.com/rs/cors"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/welcome", Api.WelcomeHandler)
	mux.HandleFunc("/calendar/callback", Api.CalendarAuthorizationCallbackHandler)
	mux.HandleFunc("/auth/calendar", Api.CalendarAuthorizationHandler)
	mux.HandleFunc("/chat", Api.ChatHandler)
	mux.HandleFunc("/", ErrorHandler)
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	handler := cors.AllowAll().Handler(mux)
	fmt.Println("Server started at http://localhost:" + port)
	http.ListenAndServe(":"+port, handler)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := Models.Error{
		Status:  false,
		Code:    http.StatusNotFound,
		Message: "404 Not Found."}
	err.ErrorAsJSON(w)
}

func main() {
	godotenv.Load()
	Database.Init()
	StartServer()
}