package main

import (
	"fmt"
	"songs_library/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("failed to load config: %v", err)
		return
	}

	fmt.Println("port: %v", cfg.DBHost)
}
