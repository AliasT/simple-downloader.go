package main

// https://golangcode.com/handle-ctrl-c-exit-in-terminal/

import (
	"fmt"
	. "io"
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
		os.Remove(tempName)
		log.Fatalln("网络请求失败:", err)
	}

	defer res.Body.Close()
	copy(tmp, res.Body)
	os.Rename(tempName, name)
}

// CopyBuffer copy from std copyBuffer
func copy(dst Writer, src Reader) (written int64, err error) {
	buf := make([]byte, 32*1024)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
				fmt.Print("\r", written)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
