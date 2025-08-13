package basicsec

import (
	"bufio"
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"
)

func RequestInit() *Request {
	return &Request{Wordlist: nil, thread: 5, lines: make(chan string), ParamsType: "FORM"}
}

func (s *Request) TakeInputs(Url, Type, Params, Cookie string, wordlist *multipart.FileHeader) int {
	if Type == "GET" {
		Url = addQueryToTheGetUrl(Url, Params)
		if Url == "6" {
			// Hatalı params girildi diye hata döndür
			return 6
		}
	}

	if Url == "" {
		return 1
	}

	// Eğer Get ise ve url de keyword yok ise keyword ekle uyarısı at.
	if Type == "GET" {
		if !strings.Contains(Url, "*") {
			return 5
		}
	}
	s.Url = Url

	// Type doğru mu geliyor bakalım.
	if Type != "POST" && Type != "GET" {
		return 2
	}
	s.Type = Type

	// Paramları ayırıyorum.
	if Params == "" && Type == "POST" {
		return 3
	}

	Params = strings.TrimSpace(Params)
	paramsArr := strings.Split(Params, "\r\n")
	s.Params = paramsArr

	s.Cookie = Cookie

	s.Wordlist = wordlist

	return -1
}

func (s *Request) EmptyRequest(Responses *[]Response) {
	newUrl := GetUrl(s.Url, "", "*", s.Type)
	s.Scan(newUrl, " ", Responses, nil)
}

// Thread'leri başlatıyorum.
func (s *Request) Start() []Response {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	var Responses []Response
	s.EmptyRequest(&Responses)
	for i := 0; i < s.thread; i++ {
		wg.Add(1)
		go s.ProcessLine(&wg, &Responses, &mutex)
	}

	s.readWordlist()
	close(s.lines)

	wg.Wait()

	return Responses
}

// Okuduğum satırları burada işliyorum
func (s *Request) ProcessLine(wg *sync.WaitGroup, Responses *[]Response, mutex *sync.Mutex) {
	defer wg.Done()

	for e := range s.lines {
		url := GetUrl(s.Url, e, "*", s.Type)
		s.Scan(url, e, Responses, mutex)
	}
}

// Wordlist'i okuyorum.
func (s *Request) readWordlist() {
	wordlist, err := s.Wordlist.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wordlist.Close()

	scanner := bufio.NewScanner(wordlist)
	for scanner.Scan() {
		// url encode atarak yolluyorum ki sıkıntı çıkmasın
		// Eğer get ise urlencode atıyorum post iste gerek yok.
		var line string
		if s.Type == "GET" {
			line = url.QueryEscape(scanner.Text())

		} else if s.Type == "POST" {
			line = scanner.Text()
		}

		s.lines <- line
	}
}

func (s *Request) Scan(Url, line string, Responses *[]Response, mutex *sync.Mutex) {
	switch s.Type {
	case "GET":
		s.Get(Url, line, Responses, mutex)
		break
	case "POST":
		s.Post(Url, line, Responses, mutex)
		break
	}
}

func (s *Request) Get(Url, line string, Responses *[]Response, mutex *sync.Mutex) {
	req, err := http.NewRequest(http.MethodGet, Url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := s.getClient(req.URL)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Client Timeout!, Make sure your URL is correct")
		fmt.Println(err)
		return
	} else {
		// Response oluşturuyorum ve atıyorum
		response := Response{}

		response.Line = line
		response.ContentLength = res.ContentLength
		response.Status = res.StatusCode

		// Array'e atıyorum.
		// Race condition atmadan önce bir thread girdiyse diye lock var
		if mutex != nil {
			mutex.Lock()
			*Responses = append(*Responses, response)
			mutex.Unlock()
		} else {
			*Responses = append(*Responses, response)
		}
	}
	defer res.Body.Close()
}

func (s *Request) Post(Url, line string, Responses *[]Response, mutex *sync.Mutex) {
	// Aldığım parametreleri ekliyorum request'e form olarak
	// Data 0 gelmezse bir hata var direkt resutrn atıyorum
	data, dErr := postData(s.Params, line, "*", s.ParamsType)
	if dErr != 0 {
		return
	}
	req, err := http.NewRequest(http.MethodPost, Url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	if s.ParamsType == "JSON" {
		req.Header.Set("Content-Type", "application/json")
	} else if (s.ParamsType) == "FORM" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := s.getClient(req.URL)
	if client == nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Client Timeout!, Make sure your URL is correct")
		return
	} else {
		response := Response{}

		response.Line = line
		response.ContentLength = res.ContentLength
		response.Status = res.StatusCode

		// Array'e atıyorum.
		// Race condition atmadan önce bir thread girdiyse diye lock var
		if mutex != nil {
			mutex.Lock()
			*Responses = append(*Responses, response)
			mutex.Unlock()
		} else {
			*Responses = append(*Responses, response)
		}
	}
	defer res.Body.Close()
}

func (s *Request) getClient(requestURL *url.URL) *http.Client {
	// Cookie'leri tutmak için cookie jar oluşturduk.
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// cookie jar'ı client'a ekledik.
	client := &http.Client{
		Jar: cookieJar,
	}

	client.Timeout = time.Duration(10) * time.Second

	if s.Cookie != "" && strings.Count(s.Cookie, ":") == 1 {
		piece := strings.Split(s.Cookie, ":")
		name := piece[0]
		value := piece[1]

		// Cookie setliyorum
		var cookies []*http.Cookie
		cookie := &http.Cookie{
			Name:  name,
			Value: value,
		}

		cookies = append(cookies, cookie)

		client.Jar.SetCookies(requestURL, cookies)
	}

	return client
}
