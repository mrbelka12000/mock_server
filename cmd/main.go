package main

import (
	"fmt"
	"net/http"

	"mock_server/handler/bizon"
	"mock_server/handler/epg"
	"mock_server/handler/magua"
)

func main() {

	// EPG handlers
	{
		http.HandleFunc("/api/purchase", epg.PurchaseHandler)
		http.HandleFunc("/api/oauth/token", epg.GetTokenHandler)
		http.HandleFunc("/api/status", epg.StatusResponseHandler)
	}

	// Bizon
	{
		http.HandleFunc("/orders/authorize", bizon.PurchaseCreateError)
		http.HandleFunc("/orders", bizon.StatusResponseHanlder)
	}

	// Magua
	{
		http.HandleFunc("/withdrawal/init", magua.CreatePayoutHandler)
		http.HandleFunc("/check", magua.GetStatusHandler)
	}

	//health check
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	fmt.Println("Listening on port 5555")
	http.ListenAndServe(":5555", nil)
}
