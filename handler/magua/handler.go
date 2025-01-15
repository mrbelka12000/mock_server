package magua

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type (
	// PayoutStatusResponse
	PayoutStatusResponse struct {
		Status       int    `json:"status"`
		RRN          string `json:"rrn"`
		Err          string `json:"err"`
		ErrorMessage string `json:"errorMessage"`
	}

	// CreatePayoutResponse
	CreatePayoutResponse struct {
		TrackingKey string `json:"order_id"`
	}
)

func GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(getStatusResponse())
}

func CreatePayoutHandler(w http.ResponseWriter, r *http.Request) {
	resp := CreatePayoutResponse{
		TrackingKey: uuid.New().String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getStatusResponse() []byte {
	resp := PayoutStatusResponse{
		Status:       -1,
		RRN:          "from_mock_server_with_error",
		Err:          "Undefined",
		ErrorMessage: "Decline from acquiring",
	}

	body, _ := json.Marshal(resp)
	return body
}
