package main

type Target struct {
	Host   string `json:"host"`
	Status int    `json:"status"`
}
type Source struct {
	Host     string `json:"host"`
	Interval int    `json:"interval"`
	Method   string `json:"method"`
	Proxy    string `json:"proxy"`
	LastCode int    `json:"lastCode"`
}
