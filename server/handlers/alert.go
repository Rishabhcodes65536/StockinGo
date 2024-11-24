package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Rishabhcodes65536/StockinGo/services"
	"github.com/Rishabhcodes65536/StockinGo/models"
)

// CreateAlert handler
func CreateAlert(alertService services.AlertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var alert models.Alert
		if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := alertService.CreateAlert(&alert); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Alert created successfully")
	}
}

// UpdateAlert handler
func UpdateAlert(alertService services.AlertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alertID := r.URL.Query().Get("id")
		var alert models.Alert
		if err := json.NewDecoder(r.Body).Decode(&alert); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := alertService.UpdateAlert(alertID, &alert); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Alert updated successfully")
	}
}

// DeleteAlert handler
func DeleteAlert(alertService services.AlertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alertID := r.URL.Query().Get("id")
		if err := alertService.DeleteAlert(alertID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Alert deleted successfully")
	}
}

// GetAlerts handler
func GetAlerts(alertService services.AlertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		alerts, err := alertService.GetAlerts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(alerts)
	}
}
