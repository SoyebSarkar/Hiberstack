package main

import (
	"log"
	"net/http"

	"github.com/SoyebSarkar/Hiberstack/internal/config"
	"github.com/SoyebSarkar/Hiberstack/internal/proxy"
)

func main() {
	cfg := config.Load()

	p, err := proxy.New(cfg.Typesense.URL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("hiberstack listening on", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, p)
}
