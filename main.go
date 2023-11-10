package main

import (
	"eckert/app/news_app"
	"eckert/infrastructure/router"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	//discord.GetNews()
	router.InitRouter()
	news_app.InitNews()
	_ = router.RT.Run(":8000")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
