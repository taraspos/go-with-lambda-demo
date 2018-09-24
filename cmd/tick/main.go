package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fw "github.com/trane9991/go-with-lambda"
)

var fwMode = os.Getenv("FW_MODE")

func handler() (*events.APIGatewayProxyResponse, error) {
	var data string

	files, err := ioutil.ReadDir(fw.WriteDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			path := fw.WriteDir + f.Name()
			d, err := ioutil.ReadFile(path)
			fmt.Println(err)
			data += fmt.Sprintf("%s: %s\n", path, string(d))
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}, nil
}

func main() {
	if fwMode == "goroutine" {
		go fw.TickAndWrite()
	} else {
		cmd := exec.Command("bin/fw")
		if err := cmd.Start(); err != nil {
			panic(err)
		}

		if err := cmd.Process.Release(); err != nil {
			panic(err)
		}
	}

	if len(os.Args) > 1 && os.Args[1] == "start" {
		resp, err := handler()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Response: %+v\n", resp)
	} else {
		lambda.Start(handler)
	}
}
