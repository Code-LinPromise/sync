package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		gin.SetMode(gin.DebugMode)
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.Writer.Write([]byte("sdsd"))
		})
		router.Run(":8080")
	}()
	// 先写死路径，后面再照着 lorca 改
	chromePath := "C:\\Users\\林鹏涛\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chromePath, "--app=http://127.0.0.1:8080/")
	cmd.Start()
	select {}
}
