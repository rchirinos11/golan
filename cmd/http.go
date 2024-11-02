package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/rchirinos11/golan/wol"
)

var wolService wol.WolUtil

func Run(port string) {
	wolService = wol.WolUtil{}
	http.HandleFunc("/golan/mac", getMac)
	http.HandleFunc("/golan/wake", wake)
	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
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
	mac := os.Getenv("WOLMAC")
	if mac == "" {
		log.Println("Mac address to wake up not configured, set WOLMAC variable and restart the server")
		w.WriteHeader(http.StatusConflict)
		return
	}

	magic, err := wolService.MakeMagic(mac)
	if err != nil {
		log.Println("Error creating magic packet", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = wolService.SendMagic(magic)
	if err != nil {
		log.Println("Error sending magic packet", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Sending magic packet for mac:", mac)
	w.WriteHeader(http.StatusNoContent)
}
