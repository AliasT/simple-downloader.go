package main

// https://stackoverflow.com/questions/11268943/is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in-a-defe

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
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

	res, err := http.Get(target)

	if err != nil {
		log.Fatalln("网络请求失败:", err)
	}

	defer res.Body.Close()

	io.Copy(tmp, res.Body)

	os.Rename(tempName, name)
}
