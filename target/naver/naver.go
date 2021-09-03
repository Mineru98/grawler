package naver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

func getHtml() (string, error) {
	url := "https://www.naver.com"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

func Crawler(index int) (string, error) {
	html, err := getHtml()

	if err != nil {
		return html, err
	}

	doc, err := htmlquery.Parse(strings.NewReader((html)))
	
	list := htmlquery.Find(doc, "//ul/li")

	fmt.Println("======== naver ========", index, len(list))

	return html, err
}