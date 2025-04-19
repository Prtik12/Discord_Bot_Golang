package main

import (
	"fmt"

	"github.com/Prtik12/Discord_Bot_Golang/bot"
	"github.com/Prtik12/Discord_Bot_Golang/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
