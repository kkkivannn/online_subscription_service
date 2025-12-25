package main

import (
	"fmt"
	"online_subscription_service/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg.Host)
}
