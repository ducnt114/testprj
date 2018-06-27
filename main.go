package main

import (
	"fmt"
	"github.com/ducnt114/testprj/cmd"
	"github.com/ducnt114/testprj/drivers/mongo"
	"github.com/ducnt114/testprj/utils"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	utils.LoadConfig()

	initMongo()
}

func initMongo() {
	// Init database connection
	mgConn, err := mongo.NewConnection()
	if err != nil {

	}
	cmd.MongoConnection = mgConn
}

func main() {

	router := cmd.NewRouter()

	// Use middleware to catches panics and responds with a 500 response code
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(router)

	// Handle signal terminate
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		onClose()
		os.Exit(1)
	}()

	log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d",
		viper.GetString("server.host"), viper.GetInt("server.port")), n))
}

func onClose() {
	cmd.MongoConnection.Session.Close()
}
