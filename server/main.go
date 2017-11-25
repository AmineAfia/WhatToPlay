package main

import (
	"fmt"	
	"github.com/AmineAfia/WhatToPlay/server/router"
)

func main() {
	fmt.Println("sanity check")
	router.Run("127.0.0.1:8080")
}
