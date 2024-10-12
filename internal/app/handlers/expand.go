package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handlers) ExpandHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody ExpandHandlerRequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}
	expanded, err := h.store.LoadURL(requestBody.Code)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusInternalServerError)
		return
	}

	responseBody := ExpandHandlerResponseBody{
		Result: expanded,
	}
	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write(responseJSON); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

// ExpandHandlerRequestBody - определение структуры тела запроса.
type ExpandHandlerRequestBody struct {
	Code string `json:"code"`
}

// ExpandHandlerResponseBody - определение структуры тела ответа.
type ExpandHandlerResponseBody struct {
	Result string `json:"result"`
}
