package main

import (
	"api/src/router"
	"api/src/settings"
	"fmt"
	"log"
	"net/http"
)



func main(){
	settings.LoadSettings()
	r := router.Generate()

	fmt.Printf("API Starting on port %d", settings.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", settings.Port), r))
}