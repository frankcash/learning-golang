// http://oldnavy.gap.com/resources/productSearch/v1/search?cid=85729&isFacetsEnabled=true&globalShippingCountryCode=&globalShippingCurrencyCode=&locale=en_US&
// http://oldnavy.gap.com/resources/productSearch/v1/search?cid=5199&isFacetsEnabled=true&globalShippingCountryCode=&globalShippingCurrencyCode=&locale=en_US&
// http://www.gap.com/resources/productSearch/v1/search?cid=6998&isFacetsEnabled=true&globalShippingCountryCode=&globalShippingCurrencyCode=&locale=en_US&pageId=0&department=75

package main

import(
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/csv"
	"regexp"
)

type PriceInfo struct{
	CurrentMaxPrice string
}

type ChildProducts struct{
	BusinessCatalogItemId string
	Name string
	IsInStock string
	Price PriceInfo
}

type ChildCategories struct{
	ChildProducts []ChildProducts
}

type ProductCategory struct{
	Name string
	BusinessCatalogItemId string
	ChildCategories []ChildCategories
}


type ProductCategoryFacetedSearch struct{
	ProductCategory ProductCategory
}

type AddressRespSingle struct{
	ResourceVersion string
	ResourceUrl string
	ProductCategoryFacetedSearch ProductCategoryFacetedSearch
}

func main(){
	file, err := os.Create("oldnavy.csv")
	if err != nil{
		fmt.Println("Err: ", err)
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	column := []string{"Name","Max Price", "Is In Stock","Gender" , "URL"}

	writer.Write(column)
	chFinished := make(chan bool)
	ch := make(chan []string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch, chFinished)
	}

	//for msg := range ch{
		//fmt.Println(msg[1])
		//err := writer.Write(msg)
		//if err != nil{
		//	fmt.Println("Err: ", err)
		//}
	//}

	for c := 0; c < len(os.Args[1:]);{
		select{
		case msg := <- ch:
			writer.Write(msg)
		case <- chFinished:
			c++
		}
	}

}


func fetch(url string, ch chan<-[]string, chFinished chan <- bool){
	resp, err := http.Get(url)

	if err != nil{
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return
	}

	l, err := regexp.Compile(`(?i)women|ladies|lady|girl`)
	if err != nil{
		fmt.Printf("Compile err: ", err)
		return
	}
	m, err := regexp.Compile(`(?i)men|guys|guy|boy`)
	if err != nil{
		fmt.Printf("Compile err: ", err)
		return
	}

	var data AddressRespSingle // initialize struct
	json.Unmarshal(body, &data)
	for _, childCat := range data.ProductCategoryFacetedSearch.ProductCategory.ChildCategories{
		for _, childProd := range childCat.ChildProducts{
			f := []string{}
			if l.MatchString(childProd.Name) == true{
				f= []string{childProd.Name,childProd.Price.CurrentMaxPrice, childProd.IsInStock,"Female" , url}
			}else if m.MatchString(childProd.Name) == true{
				f= []string{childProd.Name,childProd.Price.CurrentMaxPrice, childProd.IsInStock,"Male" , url}
			}else{
				f= []string{childProd.Name,childProd.Price.CurrentMaxPrice, childProd.IsInStock, "None Identified" , url}
			}

			ch <- f
		}
	}
	defer func(){
		chFinished <- true
	}()
}
