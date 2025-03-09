package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/mrbelka12000/mock_server/internal"
	"github.com/mrbelka12000/mock_server/pkg/pointer"
)

func (h *DynamicRouter) HandleHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.assignHandler(w, r)
	case http.MethodGet:
	}
}

func (h *DynamicRouter) getHandler(w http.ResponseWriter, r *http.Request) {}

func (h *DynamicRouter) assignHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.With("error", err).Error("error reading body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var obj internal.HandlerCU
	err = json.Unmarshal(body, &obj)
	if err != nil {
		h.log.With("error", err).Error("error unmarshalling body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.srv.AssignHandlerToService(r.Context(), obj)
	if err != nil {
		h.log.With("error", err).Error("error assigning handler")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	srv, err := h.srv.GetServiceByID(r.Context(), pointer.Value(obj.ServiceID))
	if err != nil {
		h.log.With("error", err).Error("error getting service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add new paths in router
	h.processer(srv)
	w.Write(okResponse)
}
