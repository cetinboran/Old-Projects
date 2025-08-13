package basicsec

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

func addQueryToTheGetUrl(Url string, params string) string {
	if params == "" {
		// Eğer parametre yok ise bir şey yapmaya gerek yok
		return Url
	}
	// Query i url ye ekliyorum.
	baseURL, _ := url.Parse(Url)

	params = strings.TrimSpace(params)
	paramsArr := strings.Split(params, "\r\n")

	invalid := false
	query := url.Values{}
	for i := 0; i < len(paramsArr); i++ {
		// Eğer içeriyorsa : doğru formatta girdi içemiyorsa hata döndür
		if strings.Contains(paramsArr[i], ":") {
			holder := strings.Split(paramsArr[i], ":")[0]
			value := strings.Split(paramsArr[i], ":")[1]
			query.Add(holder, value)
		} else {
			invalid = true
			break
		}
	}

	if invalid {
		return "6"
	}

	baseURL.RawQuery = query.Encode()

	return baseURL.String()
}

func GetUrl(myUrl, fuzz, keyword, requestType string) string {
	newUrl := strings.Replace(myUrl, keyword, fuzz, 1)

	return newUrl
}

func getFormData(params []string, line, keyword string) []byte {
	formData := url.Values{}

	for _, v := range params {
		if !strings.Contains(v, ":") {
			// Eğer : yoksa format hatası var boş yolladım
			return nil
		}
		dataArr := strings.Split(v, ":")

		dataArr[1] = strings.ReplaceAll(dataArr[1], "\\r", "")
		dataArr[1] = strings.ReplaceAll(dataArr[1], "\\r", "")
		if dataArr[1] == keyword {
			formData.Add(dataArr[0], line)
		} else {
			formData.Add(dataArr[0], dataArr[1])
		}
	}

	bytes := formData.Encode()

	return []byte(bytes)
}

func getJsonData(params []string, line, keyword string) []byte {
	jsonData := make(map[string]interface{})

	for _, v := range params {
		dataArr := strings.Split(v, ":")

		if dataArr[1] == keyword {
			jsonData[dataArr[0]] = line
		} else {
			jsonData[dataArr[0]] = dataArr[1]
		}
	}

	// Json'a çeviriyorum map'i
	bytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func postData(params []string, line, keyword, dataType string) ([]byte, int) {
	var jsonData []byte
	var formData []byte

	if dataType == "FORM" {
		formData = getFormData(params, line, keyword)
	} else if dataType == "JSON" {
		jsonData = getJsonData(params, line, keyword)
	}

	if jsonData != nil && dataType == "JSON" {
		return jsonData, 1
	}

	return formData, 0
}
