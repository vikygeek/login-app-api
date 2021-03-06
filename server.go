package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"github.com/vikygeek/login-app-api/handlers"
	"net/http"
	"os"
	"time"
)
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/all", handlers.GetAllUserHandler).Methods("GET")
	r.HandleFunc("/user", handlers.GetUserDetailHandler).Methods("GET")
	r.HandleFunc("/user/set_password", handlers.SetUserPasswordHandler).Methods("PUT")
	r.HandleFunc("/user/set_mobile_number", handlers.SetMobileNumberHandler).Methods("PUT")
	r.HandleFunc("/user/search", handlers.SearchUserHandler).Methods("GET")
	r.HandleFunc("/login", handlers.Authentication).Methods("POST")
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type","Authorization"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE"},
	}).Handler(r)
	port := os.Getenv("PORT")
	srv := &http.Server{
		Handler:      handler,
		Addr:         ":"+port,            //5500 - Production; 5501 -Development
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("--Server Started @", port,"--")
	log.Fatal(srv.ListenAndServe())
}
