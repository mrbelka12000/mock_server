package main

import (
	"fmt"
	"net/http"

	"mock_server/handler/epg"
)

func main() {

	// EPG handlers
	{
		http.HandleFunc("/api/purchase", epg.PurchaseHandler)
		http.HandleFunc("/api/oauth/token", epg.GetTokenHandler)
		http.HandleFunc("/api/status", epg.StatusResponseHandler)
	}

	fmt.Println("Listening on port 5555")
	http.ListenAndServe(":5555", nil)
}
