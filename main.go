package main

import (
	"embed"
	"example.com/m/server/controllers"
	"example.com/m/server/initializers"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

//go:embed frontend/dist/*
var FS embed.FS

func GetPort() string {
	return "27149"
}

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		initializers.InitCors(router)
		staticFile, _ := fs.Sub(FS, "frontend/dist")
		router.POST("/api/v1/texts", controllers.TextController)
		router.StaticFS("/static", http.FS(staticFile))
		router.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			if strings.HasPrefix(path, "/static/") {
				reader, err := staticFile.Open("index.html")
				if err != nil {
					log.Fatal(err)
				}
				defer reader.Close()
				stat, err := reader.Stat()
				if err != nil {
					log.Fatal(err)
				}
				c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)

			} else {
				c.Status(http.StatusNotFound)
			}
		})
		router.Run(":" + GetPort())
	}()
	// 先写死路径，后面再照着 lorca 改
	chromePath := "C:\\Users\\林鹏涛\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:"+GetPort()+"/static/index.html")
	cmd.Start()
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	select {
	case <-chSignal:
		err := cmd.Process.Kill()
		if err != nil {
			return
		}
	}
}
