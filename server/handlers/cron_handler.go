package handlers

import (
	"log"

	"github.com/Rishabhcodes65536/StockinGo/services"
)

// MarketStartNotification handler
func MarketStartNotification(alertService services.AlertService) func() {
	return func() {
		err := alertService.SendMarketStartNotification()
		if err != nil {
			log.Printf("Error sending market start notification: %v", err)
		}
	}
}

// MarketEndSummary handler
func MarketEndSummary(alertService services.AlertService) func() {
	return func() {
		err := alertService.SendMarketEndSummary()
		if err != nil {
			log.Printf("Error sending market end summary: %v", err)
		}
	}
}

// SignificantChanges handler
func SignificantChanges(alertService services.AlertService) func() {
	return func() {
		err := alertService.CheckSignificantChanges()
		if err != nil {
			log.Printf("Error checking significant changes: %v", err)
		}
	}
}

// WeeklySummary handler
func WeeklySummary(alertService services.AlertService) func() {
	return func() {
		err := alertService.SendWeeklySummary()
		if err != nil {
			log.Printf("Error sending weekly summary: %v", err)
		}
	}
}
