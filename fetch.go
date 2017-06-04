// prints the content found at a URL

package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

)

func main(){
	for _, url := range os.Args[1:]{
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://"){
			url = "http://" + url
		}	
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		b, err :=  ioutil.ReadAll(resp.Body)
		s:= resp.Status
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", b)
		fmt.Printf("%s\n", s)
	}


}
