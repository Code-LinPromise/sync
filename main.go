package main

import (
	"embed"
	"example.com/m/config"
	"example.com/m/server"
	"os"
	"os/exec"
	"os/signal"
)

//go:embed frontend/dist/*
var FS embed.FS

func main() {
	go server.Run()
	cmd := startBrowser()
	chSignal := listenToInterrupt()
	select {
	case <-chSignal:
		cmd.Process.Kill()
	}
}
func startBrowser() *exec.Cmd {
	// 先写死路径，后面再照着 lorca 改
	chromePath := "C:\\Users\\林鹏涛\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+config.GetPort()+"/static/index.html")
	cmd.Start()
	return cmd
}

func listenToInterrupt() chan os.Signal {
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}
