package main

import (
	"fmt"

	"github.com/gocolly/colly"
)
type cat struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	Breed string `json:"breed"`
	Age string `json:"age"`
	ImageURL string `json:"imageUrl"`
}
func miamiDade(cats []cat){
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"),)

		
		c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement){
			cat := cat{
				Name : h.ChildText("span.text_Name.results"),
				Gender : h.ChildText("span.text_Gender.results"),
				Breed : h.ChildText("span.text_Breed.results"),
				Age : h.ChildText("span.text_Age.results"),
				ImageURL : h.ChildAttr("img", "src"),

			}
			cats = append(cats, cat)	
		})

		//this is to access the next page
		c.OnHTML("[class=page-item]", func(h *colly.HTMLElement){
			nextPage := h.Request.AbsoluteURL(h.Attr("onclick"))
			c.Visit(nextPage)
		})

		c.OnRequest(func(r *colly.Request){ //this prints out the url of the pages it visits
			fmt.Println(r.URL.String())
		})

		c.Visit("https://24petconnect.com/miad?at=CAT")
		fmt.Println(cats)
}
func humaneSocietyBroward(cats []cat){
	c := colly.NewCollector(colly.AllowedDomains("humanebroward.com"),)
	
	c.OnHTML("div[pet-item]", func(h *colly.HTMLElement){
		cat := cat{
			Name : 
		}
	})

}

func main(){
	var catsItem []cat
	humaneSocietyBroward(catsItem)
}