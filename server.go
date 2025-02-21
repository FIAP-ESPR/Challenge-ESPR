package main

import (
	"ancora/controller"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	// Build Fase [ Frontend ]
	//exec.Command(shell(), "/c", "cd ./node && npm run build").Run()

	// Build Fase [ Backend ]
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Static("/static", "./website/static")
	r.LoadHTMLGlob("./website/html/*.html")
	r.GET("/", controller.Index)
	r.Run(":8080")
}

func shell() string {
	switch runtime.GOOS {
	case "windows":
		return "cmd"
	case "linux":
		return "sh"
	default:
		return "sh"
	}
}
