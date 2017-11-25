package main

import (
	"fmt"

	"github.com/AmineAfia/WhatToPlay/server/config"
	"github.com/AmineAfia/WhatToPlay/server/models"
	"github.com/AmineAfia/WhatToPlay/server/router"
)

func main() {
	config.LoadConfiguration("config/cfg.secret")
	fmt.Println("sanity check")
	models.InitDB()
	router.Run("127.0.0.1:8080")
}
