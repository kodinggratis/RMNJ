package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kodinggratis/RMNJ.git/internal/handlers"
	"github.com/kodinggratis/RMNJ.git/internal/routers"
	"github.com/kodinggratis/RMNJ.git/internal/service"
)

func main() {
	// 1. Inisialisasi Service
	uService := service.NewUserService()

	// 2. Inisialisasi Handler & Inject Service
	uHandler := &handlers.UserHandler{Service: uService}

	// 3. Setup Router
	router := routers.NewRouter(routers.RouterDeps{UserH: uHandler})

	// 4. Jalankan Server dengan Config Produksi
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second, // Mencegah Slowloris attack
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server starts on :8080")
	log.Fatal(server.ListenAndServe())
}
