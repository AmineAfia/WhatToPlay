package main

import (
	"fmt"	
	"github.com/AmineAfia/WhatToPlay/server/router"
	"github.com/AmineAfia/WhatToPlay/server/models"
)

func main() {
	fmt.Println("sanity check")
	models.InitDB()
	router.Run("127.0.0.1:8080")
}
