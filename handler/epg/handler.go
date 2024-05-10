package epg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ( // StatusRespnonse
	statusResponse struct {
		OrderID           string    `json:"order_id"`       // Our ID
		TransactionID     string    `json:"transaction_id"` // Provider ID
		Amount            int       `json:"amount"`
		TransactionType   string    `json:"transaction_type"`
		Currency          string    `json:"currency"`
		Status            string    `json:"status"`
		StatusCode        any       `json:"status_code"`
		StatusDescription string    `json:"status_description"`
		Created           time.Time `json:"created"`
		Descriptor        string    `json:"descriptor"`
		RRN               string    `json:"rrn,omitempty"`
	}

	statusRequest struct {
	}

	// PurchaseResponse
	PurchaseResponse struct {
		OrderID           string    `json:"order_id"`       // Our ID
		TransactionID     string    `json:"transaction_id"` // Provider's ID
		Amount            int       `json:"amount"`
		Currency          string    `json:"currency"`
		Status            string    `json:"status"`
		StatusCode        any       `json:"status_code"`
		StatusDescription string    `json:"status_description"`
		Created           time.Time `json:"created"`
		CardToken         string    `json:"card_token"`
		CustomerToken     string    `json:"customer_token"`
		Descriptor        string    `json:"descriptor"`
		Redirect          *Redirect `json:"redirect,omitempty"`
		RRN               string    `json:"rrn,omitempty"`
	}

	// Redirect is 3ds redirect parameters.
	Redirect struct {
		URL        string `json:"url"`
		Method     string `json:"method"`
		Parameters []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"parameters"`
	}
)

const (
	url = "https://api.gateway-services.com/api/status?order_id=d1efa43f-8d30-46b0-97f7-5e0d82c3ed13&transaction_id=GWS202405091331029952889"
)

var (
	store = make(map[string]int)
)

func StatusResponseHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("order_id")
	trxID := r.URL.Query().Get("transaction_id")

	store[orderID]++

	if store[orderID] >= 3 {
		w.Write(getStatusResponse(orderID, trxID, "DECLINED"))
		return
	}

	w.Write(getStatusResponse(orderID, trxID, "PENDING"))
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	type tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	body, _ := json.Marshal(tokenResponse{AccessToken: "dasfsjdfadsjfasdj"})

	w.Write(body)
}

func PurchaseHandler(w http.ResponseWriter, r *http.Request) {
	resp := PurchaseResponse{
		OrderID:           uuid.New().String(),
		TransactionID:     uuid.New().String(),
		Amount:            1234,
		Currency:          "USD",
		Status:            "PENDING",
		StatusCode:        293,
		StatusDescription: "Initiated",
		Created:           time.Now(),
		CardToken:         "e1d70c60-0df3-11ef-b376-17908d235ba3",
		CustomerToken:     "e1d19cd0-0df3-11ef-ba06-75653f07d7f4",
		Descriptor:        "31242t12",
		Redirect:          nil,
		RRN:               "2314123412",
	}

	body, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func getStatusResponse(orderID, trxID, status string) []byte {
	resp := statusResponse{
		OrderID:           orderID,
		TransactionID:     trxID,
		TransactionType:   "",
		Currency:          "USD",
		Status:            status,
		StatusCode:        nil,
		StatusDescription: "",
		Created:           time.Now(),
		Descriptor:        "123431",
		RRN:               "dfjsfajsdfaj",
	}

	body, _ := json.Marshal(resp)

	return body
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
