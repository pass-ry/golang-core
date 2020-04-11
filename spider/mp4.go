package main

import (
	"fmt"
	//	"time"
	"io"
	"net/http"
	"os"
)

func main() {
	rsp, err := http.Get("http://pgcvideo.cdn.xiaodutv.com/1052227682_2956640659_2020040206580420200402112920.mp4?Cache-Control=max-age%3D8640000&responseExpires=Sat%2C+11+Jul+2020+11%3A50%3A55+GMT&xcode=1c438afc53c6a74ba294641f1d59eae217e452892dc309b5&time=1586596971&_=1586511658037")
	if err != nil {
		//fmt.Println(resp)
		fmt.Println(err)
	}

	defer rsp.Body.Close()

	f, err := os.Create("./video.mp4")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fileSize, writeErr := io.Copy(f, rsp.Body)
	if writeErr != nil {
		fmt.Println(writeErr)
	}

	fmt.Println("finish", fileSize)
}
