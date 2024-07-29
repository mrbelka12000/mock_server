package bizon

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	failure struct {
		FailureMessage string      `json:"failure_message"`
		OrderID        interface{} `json:"order_id"`
		FailureType    string      `json:"failure_type"`
		Errors         []struct {
			Message string `json:"message"`
			URI     string `json:"uri"`
		} `json:"errors"`
	}

	apiStatusResponse struct {
		Orders []struct {
			Descriptor     string `json:"descriptor"`
			Status         string `json:"status"`
			FailureMessage string `json:"failure_message"`
		}
	}
)

func PurchaseCreateError(w http.ResponseWriter, r *http.Request) {

	obj := &failure{
		FailureMessage: "test",
		OrderID:        "1234567",
		FailureType:    "reject",
		Errors: []struct {
			Message string `json:"message"`
			URI     string `json:"uri"`
		}{
			{
				Message: "no balance",
				URI:     "",
			},
			{
				Message: "empty order id",
				URI:     "",
			},
		},
	}

	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(fmt.Errorf("bizon get response error: %w", err))
		return
	}
}

func StatusResponseHanlder(w http.ResponseWriter, r *http.Request) {
	obj := &apiStatusResponse{
		Orders: []struct {
			Descriptor     string `json:"descriptor"`
			Status         string `json:"status"`
			FailureMessage string `json:"failure_message"`
		}{
			{
				Status:         "rejected",
				FailureMessage: "invalid order id",
			},
		},
	}

	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(fmt.Errorf("bizon get response error: %w", err))
		return
	}
}
