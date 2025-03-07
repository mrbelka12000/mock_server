package internal

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type (
	//Service
	Service struct {
		ID   int64  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`

		Handlers []Handler `json:"handlers,omitempty"`
	}

	// Handler
	Handler struct {
		ID        int64          `json:"id,omitempty"`
		ServiceID int64          `json:"service_id,omitempty"`
		Route     string         `json:"route,omitempty"`
		Cases     []HandlerCases `json:"cases,omitempty"`

		Service *Service `json:"service,omitempty"`
	}

	HandlerCU struct {
		ServiceID *int64           `json:"service_id,omitempty"`
		Route     *string          `json:"route,omitempty"`
		Cases     []HandlerCasesCU `json:"cases,omitempty"`
	}

	HandlerCases struct {
		ID              int64   `json:"id,omitempty"`
		HandlerID       int64   `json:"handler_id,omitempty"`
		Tag             Tag     `json:"tag,omitempty"`
		RequestBody     *string `json:"request_body,omitempty"`
		RequestHeaders  *Header `json:"request_headers,omitempty"`
		ResponseBody    *string `json:"response_body,omitempty"`
		ResponseHeaders *Header `json:"response_headers,omitempty"`
	}

	HandlerCasesCU struct {
		Tag             *Tag    `json:"tag,omitempty"`
		RequestBody     *string `json:"request_body,omitempty"`
		RequestHeaders  *Header `json:"request_headers,omitempty"`
		ResponseBody    *string `json:"response_body,omitempty"`
		ResponseHeaders *Header `json:"response_headers,omitempty"`
	}

	Header map[string][]string
)

// Value ..
func (j Header) Value() (driver.Value, error) {
	if j == nil {
		return []byte(`{}`), nil
	}
	return json.Marshal(j)
}

// Scan ..
func (j *Header) Scan(value interface{}) error {
	if value == nil {
		*j = Header{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid data type for JSONB")
	}

	return json.Unmarshal(bytes, j)
}
