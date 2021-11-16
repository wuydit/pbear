package main

import (
	"fmt"
	"log"
	"net/http"
	"pbear/routers"
	"pbear/pkg/setting"
	"github.com/gin-gonic/gin"
)

var config setting.Config

func init(){
	config = setting.InitConfig()
	log.Println("准备开始进行Main init!")
}


func main(){
	log.Println("message")
	gin.SetMode("debug")
	endPoint := fmt.Sprintf(":%d", config.Server.Port)
	// readTimeout: =
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
 }