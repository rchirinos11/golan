package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/rchirinos11/golan/wol"
)

var (
	wolService wol.WolUtil
	templater  *template.Template
)

func Run(port string) {
	initVars()
	handleAll()
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func initVars() {
	var err error
	wolService = wol.WolUtil{}
	templater, err = template.ParseGlob("views/*")
	if err != nil {
		log.Println("Error parsing templates", err)
	}
}

func handleAll() {
	http.HandleFunc("/golan/", index)
	http.HandleFunc("/golan/mac", getMac)
	http.HandleFunc("/golan/wake", wake)
	http.HandleFunc("/golan/click", click)
	http.HandleFunc("/golan/hide", hide)
}

func getMac(w http.ResponseWriter, _ *http.Request) {
	mac, err := wolService.GetMacAddr()
	if err != nil {
		log.Println("Error getting mac addresses", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r := make(map[string]string)
	r["Mac"] = mac
	response, err := json.Marshal(r)
	if err != nil {
		log.Println("Error marshalling:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Obtaining server's mac address")
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

// Wakes from preconfigured mac address
func wake(w http.ResponseWriter, _ *http.Request) {
	if err := triggerWol(w); err != nil {
		return
	}

	resp := []byte("Ok")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func triggerWol(w http.ResponseWriter) error {
	mac := os.Getenv("WOLMAC")
	if mac == "" {
		log.Println("Mac address to wake up not configured, set WOLMAC variable and restart the server")
		w.WriteHeader(http.StatusConflict)
		return fmt.Errorf("Mac address not configured")
	}

	magic, err := wolService.MakeMagic(mac)
	if err != nil {
		log.Println("Error creating magic packet", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	log.Println("Sending magic packet for mac:", mac)
	if err = wolService.SendMagic(magic); err != nil {
		log.Println("Error sending magic packet", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return err
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("GET: Index")
	if err := templater.ExecuteTemplate(w, "main", nil); err != nil {
		log.Println("Error executing template", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func click(w http.ResponseWriter, r *http.Request) {
	log.Println("POST: Htmx method")
	triggerWol(w)
	if err := templater.ExecuteTemplate(w, "wake", nil); err != nil {
		log.Println("Error executing template", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func hide(w http.ResponseWriter, _ *http.Request) {
	log.Println("PUT: Hide method")
	w.Write([]byte{})
}
