package main

import (
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"github.com/ducnt114/testprj/utils"
	"log"
	"github.com/ducnt114/testprj/cmd/testprj/services"
)

func init() {
	utils.LoadConfig()
}

func main() {

	router := services.NewRouter()

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("server.port")), router))
}
