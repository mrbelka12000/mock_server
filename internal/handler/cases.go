package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func (h *DynamicRouter) HandleCases(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.deleteCase(w, r)
	}
}

func (h *DynamicRouter) deleteCase(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.With("error", err).Error("error parsing id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.srv.DeleteCase(r.Context(), id)
	if err != nil {
		h.log.With("error", err).Error("error deleting case")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(okResponse)
	fmt.Println(err)
}
