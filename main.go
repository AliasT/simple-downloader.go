package main

// https://golangcode.com/handle-ctrl-c-exit-in-terminal/

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
)

func main() {

	if (len(os.Args)) < 2 {
		log.Fatalln("请输入下载地址")
	}

	target := os.Args[1]

	name := path.Base(target)
	tempName := name + ".download"

	tmp, err := os.Create(tempName)

	if err != nil {
		log.Fatalln("创建文件失败:", err)
	}

	// ctrl + c handler
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(tempName)
		os.Exit(0)
	}()

	res, err := http.Get(target)

	if err != nil {
		log.Fatalln("网络请求失败:", err)
	}

	defer res.Body.Close()

	io.Copy(tmp, res.Body)

	os.Rename(tempName, name)
}
