package doRequest

import (
    "io"
    "log"
    "testing"
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

func TestRequestDecode(t *testing.T) {
	var DecodeToMe Response
	type args struct {
		method   string
		url      string
		token    string
		body     io.Reader
		decodeTo interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"HTTPGet", args{
			method:   "GET",
			url:      "https://httpbin.org/get",
			token:    "",
			body:     nil,
			decodeTo: &DecodeToMe,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RequestDecode(tt.args.method, tt.args.url, tt.args.token, tt.args.body, tt.args.decodeTo); (err != nil) != tt.wantErr {
				t.Errorf("RequestDecode() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				log.Printf("%+v", DecodeToMe)
			}
		})
	}
}
