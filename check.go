package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func ControlRestart() bool {
	for {
		if restart == false {
			continue
		} else {
			return true
		}
	}
}

func Control() {
	for {
		sources := Select()
		for _, source := range sources {
			go Check(source)
		}
		if ControlRestart() == true {
			restart = false
			continue
		}
	}

}
func Check(source Source) {
	for {
		var client *http.Client
		if source.Proxy != "" {
			proxyStr := source.Proxy
			proxyURL, _ := url.Parse(proxyStr)

			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			client = &http.Client{
				Transport: transport,
			}
		} else {
			transport := &http.Transport{}
			client = &http.Client{
				Transport: transport,
			}
		}
		request, _ := http.NewRequest(source.Method, source.Host, nil)

		response, _ := client.Do(request)

		source.LastCode = response.StatusCode
		Update(source)
		fmt.Println(source)
		time.Sleep(time.Duration(source.Interval) * time.Second)

	}
}
