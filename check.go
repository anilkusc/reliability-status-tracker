package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func ControlRestart() bool {
	for {
		if restart == false {
			continue
		} else {
			if os.Getenv("DEBUG") == "true" {
				fmt.Println("1")
			}
			return true
		}
	}
}

func Control() {
	for {
		var sources []Source
		if os.Getenv("TEST") == "true" {
			sources = Select(dtbs)
		} else {
			sources = Select(NewDbConn())
		}

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
				Timeout:   30 * time.Second,
			}
		} else {
			transport := &http.Transport{}
			client = &http.Client{
				Transport: transport,
				Timeout:   10 * time.Second,
			}
		}
		request, err := http.NewRequest(source.Method, source.Host, nil)
		if err != nil {
			fmt.Println("Cannot do request: " + source.Host + " with proxy: " + source.Proxy)
			source.LastCode = 0
			if os.Getenv("TEST") == "true" {
				Update(dtbs, source)
			} else {
				Update(NewDbConn(), source)
			}

		}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("Cannot reach address: " + source.Host + " with proxy: " + source.Proxy)
			source.LastCode = 0
			if os.Getenv("TEST") == "true" {
				Update(dtbs, source)
			} else {
				Update(NewDbConn(), source)
			}
		} else {
			source.LastCode = response.StatusCode
			if os.Getenv("TEST") == "true" {
				Update(dtbs, source)
			} else {
				Update(NewDbConn(), source)
			}
		}

		time.Sleep(time.Duration(source.Interval) * time.Second)

	}
}
