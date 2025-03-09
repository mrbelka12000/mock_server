package internal

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
