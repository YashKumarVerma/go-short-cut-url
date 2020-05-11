package shortener

import (
	"YashKumarVerma/go-short-cut-url/src/env"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// struct to read server response
type responseStructure struct {
	URL responseHelper
}

type responseHelper struct {
	Status    int
	FullLink  string
	Date      string
	ShortLink string
	Title     string
}

// Shorten returns the short url for passed url
func Shorten(url string) string {
	requestURL := generateURL(url)

	resp, err := http.Get(requestURL)
	if err != nil {
		panic("Unable to connect to server")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Invalid body received")
	}

	var response responseStructure
	json.Unmarshal([]byte(body), &response)

	switch response.URL.Status {
	case 1:
		return "The shortened link comes from the domain that shortens the link, i.e. the link has already been shortened."
	case 2:
		return "The entered link is not a link"
	case 3:
		return "The preferred link name is already taken"
	case 4:
		return "Invalid API key"
	case 5:
		return "The link has not passed the validation. Includes invalid characters"
	case 6:
		return "The link provided is from a blocked domain"
	case 7:
		return response.URL.ShortLink
	default:
		return "invalid response from server"
	}

}

func generateURL(longURL string) string {
	// first create a url
	URL, _ := url.Parse("https://cutt.ly/api/api.php")
	parameters := url.Values{}
	parameters.Add("key", env.APIKey())
	parameters.Add("short", longURL)
	URL.RawQuery = parameters.Encode()
	return URL.String()
}
