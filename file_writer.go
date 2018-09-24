package fw

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const WriteDir = "/tmp/test/"

var (
	incr      = 0
	startTime = time.Now()
	ticker    = time.NewTicker(1 * time.Second)
)

func init() {
	if err := os.MkdirAll(WriteDir, 0755); err != nil {
		panic(err)
	}
}

func TickAndWrite() {
	for range ticker.C {
		incr++
		str := fmt.Sprintf("tick tock, function started %s ago", time.Since(startTime))
		if err := ioutil.WriteFile(fmt.Sprintf("%s%d.txt", WriteDir, incr), []byte(str), 0755); err != nil {
			fmt.Println(err)
		}
	}
}
