package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/mrbelka12000/mock_server/internal"
)

func (h *DynamicRouter) HandleService(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.addService(w, r)
	case http.MethodGet:
		h.getService(w, r)
	}
}

func (h *DynamicRouter) addService(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.With("error", err).Error("error reading body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var s internal.Service
	err = json.Unmarshal(body, &s)
	if err != nil {
		h.log.With("error", err).Error("error unmarshalling body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.srv.AddService(r.Context(), s)
	if err != nil {
		h.log.With("error", err).Error("error adding service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (h *DynamicRouter) getService(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		services, err := h.srv.ListServices(r.Context())
		if err != nil {
			h.log.With("error", err).Error("error listing services")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(services)
		if err != nil {
			h.log.With("error", err).Error("error encoding response")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.With("error", err).Error("error parsing id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	service, err := h.srv.GetServiceByID(r.Context(), id)
	if err != nil {
		h.log.With("error", err).Error("error getting service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(service)
	if err != nil {
		h.log.With("error", err).Error("error encoding response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
