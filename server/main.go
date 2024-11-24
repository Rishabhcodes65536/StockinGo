package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Rishabhcodes65536/StockinGo/handlers"
	"github.com/Rishabhcodes65536/StockinGo/middleware"
	"github.com/Rishabhcodes65536/StockinGo/internal/repository"
	"github.com/Rishabhcodes65536/StockinGo/services"
	"github.com/Rishabhcodes65536/StockinGo/pkg/yahoo"
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// MongoDB initialization
	ctx := context.Background()
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatal("Error disconnecting MongoDB: ", err)
		}
	}()
	db := mongoClient.Database("stockingo")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	stockRepo := repository.NewStockRepository(db)
	alertRepo := repository.NewAlertRepository(db)

	// Initialize services
	yahooClient := yahoo.NewClient()
	authService := services.NewAuthService(userRepo)
	emailService := services.NewEmailService()
	stockService := services.NewStockService(stockRepo, yahooClient)
	alertService := services.NewAlertService(alertRepo, stockRepo, emailService)

	// Set up the router
	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/api/auth/register", handlers.Register(authService)).Methods("POST")
	router.HandleFunc("/api/auth/login", handlers.Login(authService)).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Stock routes
	protected.HandleFunc("/stocks/search/{symbol}", handlers.SearchStock(stockService)).Methods("GET")
	protected.HandleFunc("/stocks/favorites", handlers.GetFavorites(stockService)).Methods("GET")
	protected.HandleFunc("/stocks/favorites/{symbol}", handlers.AddFavorite(stockService)).Methods("POST")
	protected.HandleFunc("/stocks/favorites/{symbol}", handlers.RemoveFavorite(stockService)).Methods("DELETE")

	// Alert routes
	protected.HandleFunc("/alerts", handlers.CreateAlert(alertService)).Methods("POST")
	protected.HandleFunc("/alerts", handlers.GetAlerts(alertService)).Methods("GET")
	protected.HandleFunc("/alerts/{id}", handlers.UpdateAlert(alertService)).Methods("PUT")
	protected.HandleFunc("/alerts/{id}", handlers.DeleteAlert(alertService)).Methods("DELETE")

	// Initialize CRON jobs
	scheduler := gocron.NewScheduler(time.UTC)

	// Market start notification (5 minutes before)
	scheduler.Every(1).Day().At("08:55").Do(func() {
		if err := alertService.SendMarketStartNotification(); err != nil {
			log.Println("Error in SendMarketStartNotification:", err)
		}
	})

	// Market end summary
	scheduler.Every(1).Day().At("16:00").Do(func() {
		if err := alertService.SendMarketEndSummary(); err != nil {
			log.Println("Error in SendMarketEndSummary:", err)
		}
	})

	// Significant change check
	scheduler.Every(10).Minutes().Do(func() {
		if err := alertService.CheckSignificantChanges(); err != nil {
			log.Println("Error in CheckSignificantChanges:", err)
		}
	})

	// Weekly summary
	scheduler.Every(1).Friday().At("16:30").Do(func() {
		if err := alertService.SendWeeklySummary(); err != nil {
			log.Println("Error in SendWeeklySummary:", err)
		}
	})

	// Start the CRON scheduler
	scheduler.StartAsync()

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Server failed:", err)
	}
}
