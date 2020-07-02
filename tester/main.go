package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	params := []WaitParam{
		{
			Interval: 20 * time.Second,
			Wait:     30 * time.Second,
		},
	}

	for i := 0; i < 10; i++ {
		params = append(params, []WaitParam{
			{
				Interval: 1 * time.Second,
				Wait:     1 * time.Second,
			},
			{
				Interval: 500 * time.Millisecond,
				Wait:     2 * time.Second,
			},
			{
				Interval: 500 * time.Millisecond,
				Wait:     3 * time.Second,
			},
			{
				Interval: 500 * time.Millisecond,
				Wait:     5 * time.Second,
			},
		}...)
	}

	var wg sync.WaitGroup

	for _, param := range params {
		param := param
		wg.Add(1)
		go func() {
			defer wg.Done()
			reqWait(param)
		}()
	}

	wg.Wait()
}

type WaitParam struct {
	Interval time.Duration
	Wait     time.Duration
}

func reqWait(param WaitParam) {
	PORT := getEnv("PORT", "8080")
	client := makeHttpClient()
	waitSec := int(param.Wait.Seconds())
	url := fmt.Sprintf("http://localhost:%s/wait/%d", PORT, waitSec)

	for {
		jitteredSleep(3 * time.Second)

		err := func() error {
			resp, err := client.Get(url)
			if err != nil {
				return err
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			fmt.Println(string(body))
			return nil
		}()

		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		time.Sleep(param.Interval)
	}
}

func jitteredSleep(max time.Duration) {
	sec := rand.Float64() * max.Seconds()
	time.Sleep(time.Duration(sec) * time.Second)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func makeHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				KeepAlive: 10 * time.Second,
			}).DialContext,
		},
	}
}
