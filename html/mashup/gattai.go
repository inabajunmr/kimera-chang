package mashup

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/inabajunmr/gattai/html/path"
)

func Gattai(url1 string, url2 string) string {

	resp1, err := http.Get(url1)
	if err != nil {
		print("Can not access to " + url1)
		print(err)
		os.Exit(1)
	}

	resp2, err := http.Get(url2)
	if err != nil {
		print("Can not access to " + url2)
		print(err)
		os.Exit(1)
	}

	absHTML1 := path.ModifyToAbsoluteURLInHTML(resp1.Body, url1)
	absHTML2 := path.ModifyToAbsoluteURLInHTML(resp2.Body, url2)

	doc1, err := goquery.NewDocumentFromReader(strings.NewReader(absHTML1))
	if err != nil {
		print("Something wrong")
		print(err)
		os.Exit(1)
	}

	doc2, err := goquery.NewDocumentFromReader(strings.NewReader(absHTML2))
	if err != nil {
		print("Something wrong")
		print(err)
		os.Exit(1)
	}

	result, _ := goquery.NewDocumentFromReader(strings.NewReader("<!DOCTYPE html><html><head></head><body></body></html>"))
	// head
	rand.Seed(time.Now().UnixNano())

	head1 := doc1.Find("head").Children()
	head2 := doc2.Find("head").Children()
	result.Find("head").AppendSelection(head1)
	result.Find("head").AppendSelection(head2)

	// body
	body1 := doc1.Find("body")
	body2 := doc2.Find("body")

	result.Find("body").AppendSelection(body1).AppendSelection(body2)
	val, err := result.Html()

	return val

}

func extract(s *goquery.Selection) {
	// get depth and node index
	// get random node and remove it
	a := goquery.NewDocumentFromNode(s.Nodes[0]).Selection
	fmt.Println(a)
}
