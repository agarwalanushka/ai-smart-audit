package main

import (
	"ai-smart-audit/config"
	"context"
	"fmt"
)

func main() {
	fmt.Println("Server is running...")
	_ = context.Background()
	config := config.NewAppConfig()
	//Todo: Pass all the functions that need to validate the config, like db, cache etc
	config.Validate()
}
