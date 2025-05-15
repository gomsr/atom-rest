package core

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer(port int, Router *gin.Engine) {

	address := fmt.Sprintf(":%d", port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Printf("server run success on address: %v\n", address)
	fmt.Printf(s.ListenAndServe().Error())
}
