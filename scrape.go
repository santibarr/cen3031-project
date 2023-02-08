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
//make a default value since many websites do not have all the information for the cat object
func (catDefault *cat) SetDefault(){
	if catDefault.Name == ""{
		catDefault.Name = "Unknown"
	}
	if catDefault.Age == ""{
		catDefault.Age = "Unknown"
	}
	if catDefault.Gender == ""{
		catDefault.Gender = "Unknown"
	}
	if catDefault.Breed == ""{
		catDefault.Breed = "Unknown"
	}
	if catDefault.ImageURL == ""{
		catDefault.ImageURL = "Unknown"
	}
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
		// c.OnHTML("[class=page-item]", func(h *colly.HTMLElement){
		// 	nextPage := h.Request.AbsoluteURL(h.Attr("onclick"))
		// 	c.Visit(nextPage)
		// })

		// c.OnRequest(func(r *colly.Request){ //this prints out the url of the pages it visits
		// 	fmt.Println(r.URL.String())
		// })

		c.Visit("https://24petconnect.com/miad?at=CAT")
		fmt.Println(cats)
}
func humaneSocietyBroward(cats []cat){
	c := colly.NewCollector(colly.AllowedDomains("humanebroward.com"),)
	
	c.OnHTML("div[id=pets-container]", func(h *colly.HTMLElement){
		// cat := cat{
		// 	Name : h.ChildText("h3.pet-name"),
		// 	Gender : h.ChildText("div.pet-detail"),
		// 	Breed : h.ChildText("div.pet-detail"),
		// 	Age : h.ChildText("div.pet-detail"),
		// 	ImageURL : h.ChildAttr("img", "src"),
		// }
		// cats = append(cats, cat)
		fmt.Println(h.Text)
	})
	c.Visit("https://humanebroward.com/all-pets/?type=CAT&pg=1")
	//fmt.Println(cats)

}
func lakeCounty(cats []cat){
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"),)

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement){
		cat := cat{
			Name : h.ChildText("span.text_Name.results"),
			Gender : h.ChildText("span.text_Gender.results"),
			Age : h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),

		}

		cats = append(cats, cat)
	})

	c.Visit("https://24petconnect.com/LakeCountyAdoptablePets?at=CAT&sb=id_asc")
	fmt.Println(cats)
}
func seminoleCounty(cats []cat){
	c := colly.NewCollector(colly.AllowedDomains("petharbor.com"),)

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement){
		cat :=cat{
			Name : h.ChildText("div.gridText"),
			// Gender : h.ChildText("div.gridText"),
			// Breed : h.ChildText("div.gridText"),
			// Age : h.ChildText("div.gridText"),
			// ImageURL: h.ChildAttr("img", "src"),

		}
		cats = append(cats, cat)	

	})
	c.Visit("https://petharbor.com/results.asp?searchtype=ADOPT&start=4&stylesheet=https://ominosity.github.io/smnl.css&grid=1&friends=1&samaritans=1&nosuccess=0&rows=10&imght=120&imgres=detail&tWidth=200&view=sysadm.v_smnl&nomax=1&nocustom=1&fontface=arial&fontsize=10&col_bg=ac5a5a&col_bg2=37c94e&miles=20&shelterlist=%27smnl%27&atype=&where=type_CAT&PAGE=1")
	fmt.Println(cats)

}
func main(){
	var catsItem []cat
	//humaneSocietyBroward(catsItem)
	// lakeCounty(catsItem)
	// miamiDade(catsItem)
	seminoleCounty(catsItem)
}