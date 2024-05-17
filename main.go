package main

import (
	"fmt"

	"github.com/felipespinassi/goportunities/config"
	"github.com/felipespinassi/goportunities/router"
)

func main() {
	//initialize configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	//initialize router
	router.Initialize()
}
