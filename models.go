package main

type Source struct {
	Host     string `json:"host"`
	Desired  int    `json:"desired"`
	Interval int    `json:"interval"`
	Method   string `json:"method"`
	Proxy    string `json:"proxy"`
	LastCode int    `json:"lastCode"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
