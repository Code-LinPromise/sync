package main

import (
	"example.com/m/config"
	"example.com/m/server"
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"syscall"
)

////go:embed frontend/dist/*
//var FS embed.FS

func main() {
	go server.Run()
	var ui lorca.UI
	ui, _ = lorca.New("http://127.0.0.1:"+config.GetPort()+"/static/index.html", "", 1280, 900, "--disable-sync", "--disable-translate", "--remote-allow-origins=*")
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ui.Done():
	case <-chSignal:
	}
	ui.Close()
}
