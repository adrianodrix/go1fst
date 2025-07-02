package main

import (
	"fmt"
	"log"
	"net/http"

	"meu-projeto/utils"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := utils.FormatMessage("Olá! Este é meu primeiro servidor web em Go!")
		fmt.Fprintf(w, message)
	})

	r.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		timestamp := utils.GetCurrentTime()
		fmt.Fprintf(w, `{"status": "online", "timestamp": "%s", "message": "Servidor funcionando!"}`, timestamp)
	})

	r.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		info := utils.GetInternalInfo()
		fmt.Fprintf(w, `{"info": "%s"}`, info)
	})

	fmt.Println(utils.FormatMessage("Servidor rodando em http://localhost:8080"))
	log.Fatal(http.ListenAndServe(":8080", r))
}
