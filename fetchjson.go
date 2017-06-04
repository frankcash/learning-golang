// designed to make a web request to https://api.etherscan.io/api?module=account&action=balance&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&tag=latest&apikey=YourApiKeyToken

package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type AddressRespSingle struct{
	Status string
	Message string
	Result string

}

func main(){
	ch := make(chan string)

	for _, url:= range os.Args[1:]{
		go fetch(url, ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
}


func fetch(url string, ch chan<-string){
	resp, err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err) // send to ch
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	var data AddressRespSingle
	json.Unmarshal(body, &data)

	ch <- fmt.Sprintf("Results: %v\n", data.Result)



}
