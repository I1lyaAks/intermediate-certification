package main

import (
	"hi/pkg/file"
	s "hi/pkg/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
)

func main() {
	file.OpenFileCsv()
	go func() {
		log.Fatal(http.ListenAndServe(":8080", run()))
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	file.CloseFileCsv()
}

func run() http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", s.CityID)
	r.Post("/Create", s.CreateCity)
	r.Delete("/{id}", s.DeleteCity)
	r.Put("/{id}", s.UpdatePopulation)
	r.Get("/Cityregion/{region}", s.CityRegion)
	r.Get("/Citydistrict/{district}", s.CityDistrict)
	r.Get("/Population/{min}:{max}", s.Population)
	r.Get("/Poundation/{min}:{max}", s.Poundation)
	return r
}
