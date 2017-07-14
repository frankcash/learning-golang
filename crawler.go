package example

// based on https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html

import(
	"fmt"
	"os"
	"strings"
	"net/http"
	"golang.org/x/net/html"
)
func getHref(t html.Token) (ok bool, href string){
	// iterate all of token's attr
	for _, a := range t.Attr{
		if a.Key == "href"{
			href = a.Val
			ok = true
		}
	}
	// "bare" return will return ok, href
	return
}


func crawl(url string, ch chan string, chFinished chan bool){
	resp, err := http.Get(url)
	defer func(){
		// Notify that we're done after this func

		chFinished <- true
	}()

	if err != nil{
		fmt.Println("Error: failed to crawl\"" + url + "\"")
		return
	}

	b:= resp.Body
	defer b.Close() // close Body when function returns

	z := html.NewTokenizer(b)

	for{
		tt := z.Next()
		switch{
			case tt == html.ErrorToken:
				// end of the document
				return
			case tt == html.StartTagToken:
				t := z.Token()

				isAnchor := t.Data == "a"
				if !isAnchor{
					continue
				}

				// check the href val
				ok, url:= getHref(t)
				if !ok{
					continue
				}

				// make sure the url begins in http**
				hasProto := strings.Index(url, "http") == 0
				if hasProto{
					ch <- url
				}
		}

	}
}

func main(){
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]
	// channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range seedUrls{
		go crawl(url, chUrls, chFinished)
	}

	// subscribe to both channels
	for c:= 0; c <len(seedUrls);{
		select{
		case url:= <- chUrls:
			foundUrls[url] = true
		case <- chFinished:
			c++
		}
	}

	// print results
	fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

	for url, _ := range foundUrls{
		fmt.Println(" - " + url)
	}

	close(chUrls)

}
