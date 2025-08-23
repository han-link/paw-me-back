package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"paw-me-back/internal/db"
	"paw-me-back/internal/env"
	"paw-me-back/internal/store"

	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/supertokens"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			PawMeBack API
//	@description	API for PawMeBack, an expense manager

// @description
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Couldn't load .env file")
	}

	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:5173"),
		env:         env.GetString("ENV", "development"),
	}

	cfg.apiURL = env.GetString("EXTERNAL_URL", fmt.Sprintf("localhost%s", cfg.addr))

	authBasePath := "/auth"

	cfg.superTokensConfig = supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://localhost:3567",
			APIKey:        env.GetString("SUPER_TOKEN_API_KEY", ""),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "SuperTokens Demo App",
			APIDomain:       fmt.Sprintf("http://%s", cfg.apiURL),
			WebsiteDomain:   cfg.frontendURL,
			APIBasePath:     &authBasePath,
			WebsiteBasePath: &authBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
			dashboard.Init(nil),
			userroles.Init(nil),
		},
	}

	logDir := "./logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}

	// Logger
	loggerConfig := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./logs/app.log"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	// Unmarshal the JSON configuration into a zap.Config
	var loggerCfg zap.Config
	if err := json.Unmarshal(loggerConfig, &loggerCfg); err != nil {
		log.Fatalf("Error unmarshaling zap config: %v", err)
	}

	// Build the logger from the custom configuration
	must := zap.Must(loggerCfg.Build())
	defer must.Sync()

	// Create a sugared logger from the built logger
	logger := must.Sugar()

	logger.Info("Sugared logger constructed successfully")

	database, err := db.New()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connection established")

	err = supertokens.Init(cfg.superTokensConfig)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Supertokens initialized")

	storage := store.NewStorage(database)

	app := &application{
		config: cfg,
		logger: logger,
		store:  storage,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
