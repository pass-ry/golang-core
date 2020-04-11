package main

import (
	//	"fmt"
	"github.com/beiping96/grace"
	"test/block/golang-core/httpd"
	"time"
)

func main() {
	grace.Go(httpd.StartHttpsServe)
	grace.Run(time.Duration(1) * time.Second)
}
