package basicsec

import "mime/multipart"

type Request struct {
	Url        string
	Type       string
	Params     []string
	Cookie     string
	Wordlist   *multipart.FileHeader
	ParamsType string
	thread     int
	lines      chan string
}

type Response struct {
	Line          string
	ContentLength int64
	Status        int
}
