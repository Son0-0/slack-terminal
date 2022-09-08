package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Son0-0/slack-terminal/handlers"
)

var serverLogger *log.Logger

func main() {
	outLog, err := os.OpenFile("out.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(outLog)
	defer outLog.Close()

	if handlers.HandleError(err) {
		panic(err)
	}

	http.HandleFunc("/slack/cmd", handlers.ExecCommand)

	fmt.Println("Server running on port 9999")
	http.ListenAndServe(":9999", nil)
}
