package utility

import (
	"crypto/md5"
	"encoding/hex"
	"math"
	"strconv"

	"github.com/cetinboran/basicsec/basicsec"
	"github.com/cetinboran/basicsec/models"
)

func ConvertToMd5(data string) string {
	hash := md5.New()

	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

func FilterResponse(url string, Responses []basicsec.Response) []basicsec.Response {
	// İlk eleman her zaman empty request.
	// Empty reuqestin değerlerine göre filtreleme yapıyorum.
	filterBy := Responses[0]

	var filteredResponses []basicsec.Response
	for _, r := range Responses {
		// Empty request'ten farklı ise eğerler girsin değil ise girmesin.
		// || r.Status != filterBy.Status bunu şimdilik kapadım belki aynı status aynı content length gelir.
		if r.ContentLength != filterBy.ContentLength {
			response := basicsec.Response{}

			response.ContentLength = r.ContentLength
			response.Line = r.Line
			response.Status = r.Status

			filteredResponses = append(filteredResponses, response)
		}
	}

	return filteredResponses
}

func Pages(arr []models.Scanes, currentPage string) models.Page {
	page := models.Page{}

	currentPageFloat, err := strconv.ParseFloat(currentPage, 64)
	if err != nil {
		return models.Page{}
	}

	var maxList float64 = 13
	length := len(arr)

	pageCount := math.Ceil(float64(length) / maxList)

	start := (currentPageFloat - 1) * maxList
	end := ((currentPageFloat - 1) * maxList) + maxList

	var pageCountSlice []int
	for i := 1; i < int(pageCount)+1; i++ {
		pageCountSlice = append(pageCountSlice, i)
	}

	page.PageCount = pageCountSlice
	page.Start = int(start)
	page.End = int(end)

	return page
}
