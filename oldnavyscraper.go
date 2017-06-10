// http://oldnavy.gap.com/resources/productSearch/v1/search?cid=85729&isFacetsEnabled=true&globalShippingCountryCode=&globalShippingCurrencyCode=&locale=en_US&

package main

import(
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}
	
	//fmt.Println(<-ch)
	//for{
	//	fmt.Println(<-ch)
	//	fmt.Println(<-end)
	//}

	//<-done
	for msg := range ch{
		fmt.Println(msg)
	}

}


func fetch(url string, ch chan<-string){
	resp, err := http.Get(url)

	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	var data AddressRespSingle // initialize struct
	json.Unmarshal(body, &data)
	for _, childCat := range data.ProductCategoryFacetedSearch.ProductCategory.ChildCategories{
		for _, childProd := range childCat.ChildProducts{
			ch <- fmt.Sprintf("Prod name: %v, Prod price: %v\n",childProd.Name, childProd.Price.CurrentMaxPrice)
		}
	}
	//end <- true 
	//ch <- fmt.Sprintf("Results: %v\n", data)
	close(ch)
}
