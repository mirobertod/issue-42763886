package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	url := os.Args[1]
	count := 0

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}

	dt := time.Now()
	fmt.Println("Start time: ", dt.String())

	for {
		count++
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			dt := time.Now()
			fmt.Println("End time: ", dt.String())
			log.Fatalln(err)
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("count: ", count)
			fmt.Println("url: ", url)

			dt := time.Now()
			fmt.Println("End time: ", dt.String())

			respDump, err := httputil.DumpResponse(resp, true)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("RESPONSE:\n%s", string(respDump))
			break
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Printf("%d requests to %s\n", count, url)
	}
}
