package main 

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"	
	"marksheets-pipeline/config"
)


func main() {
	cfg, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config %v", err)
	}
	// logger := logging.NewLogger()
	// python processor client 
	// api key validator 
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60*time.Second))
	
}