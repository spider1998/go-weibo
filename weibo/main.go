package main

import (
	"WeiPro/weibo/app"
	"WeiPro/weibo/handler"
	"fmt"
)

func main() {
	//-----------------------------初始化配置---------------------------------------------------------------------
	err := app.Init()
	if err != nil {
		fmt.Println(err)
	}
	//-----------------------------执行服务---------------------------------------------------------------------
	for true {
		handler.Run()
	}
}
