package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sunil-dev608/mars-rover/config"
	"github.com/sunil-dev608/mars-rover/internal/world"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	err = godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file:", zap.Error(err))
	}

	cfg := config.GetConfig()
	w := world.NewWorld()
	file, err := os.Open(cfg.Filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	if err = w.ReadData(file); err != nil {
		logger.Fatal("Error reading file:", zap.Error(err))
	}
	w.MoveRobots()
	w.PrintWorld()
}
