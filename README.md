# doRequest

Simple function to create custom request. Used in my projects to reduce the amount of code.

# Example usage

```go
package main

import (
    "fmt"
    "github.com/hugmouse/doRequest"
)

type Response struct {
	Args struct {
	} `json:"args"`
	Headers struct {
		Accept         string `json:"Accept"`
		AcceptEncoding string `json:"Accept-Encoding"`
		AcceptLanguage string `json:"Accept-Language"`
		Dnt            string `json:"Dnt"`
		Host           string `json:"Host"`
		Referer        string `json:"Referer"`
		UserAgent      string `json:"User-Agent"`
	} `json:"headers"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
    var DecodeToMe Response
    err := doRequest.RequestDecode("GET", "https://httpbin.org/get", "", nil, &DecodeToMe)
    if err != nil {
        fmt.Errorf("RequestDecode() error = %v", err)
    }
    fmt.Println(DecodeToMe.URL) // "https://httpbin.org/get"
}
```