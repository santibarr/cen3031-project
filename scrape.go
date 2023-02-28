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
func  SetDefaultUnknown(catDefault *cat){
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
func SetDefaultInWebsite(catWebDefault *cat){
	if catWebDefault.Name == ""{
		catWebDefault.Name = "On Website"
	}
	if catWebDefault.Age == ""{
		catWebDefault.Age = "On Website"
	}
	if catWebDefault.Gender == ""{
		catWebDefault.Gender = "On Website"
	}
	if catWebDefault.Breed == ""{
		catWebDefault.Breed = "On Website"
	}
	if catWebDefault.ImageURL == ""{
		catWebDefault.ImageURL = "On Website"
	}
}


func miamiDade(cats []cat) []cat{
	//var catz []cat
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"),)

		
		c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement){ //gets the elements from the gridResult
			cat := cat{
				Name : h.ChildText("span.text_Name.results"),
				Gender : h.ChildText("span.text_Gender.results"),
				Breed : h.ChildText("span.text_Breed.results"),
				Age : h.ChildText("span.text_Age.results"),
				ImageURL : h.ChildAttr("img", "src"),
				

			}
			cats = append(cats, cat)// adds to the vector of cat objects
			
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
		return cats
}

func lakeCounty(cats []cat) []cat{
	
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"),)

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement){ //makes the cat object from the tags on the HTML page
		cat := cat{
			Name : h.ChildText("span.text_Name.results"),
			Gender : h.ChildText("span.text_Gender.results"),
			Breed : "",
			Age : h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat) //Adds them into the vector of cat objects
	})

	c.Visit("https://24petconnect.com/LakeCountyAdoptablePets?at=CAT&sb=id_asc")
	fmt.Println(cats)
	return cats
}

func peggy(cats []cat) []cat{
	c := colly.NewCollector()

	c.OnHTML("div[class=animal-select]", func(h *colly.HTMLElement){
		cat := cat{
			Name : h.ChildText("div.animal-name.text-center"),
			// Gender : h.ChildText("div.list-animal-sexSN"),
			// Breed : h.ChildText("div.list-animal-breed"),
			// Age : h.ChildText("div.list-animal-age"),
			ImageURL : h.ChildAttr("img", "src"),
		}
		SetDefaultInWebsite(&cat)
		cats = append(cats,cat)

	})
	c.Visit("https://www.peggyadams.org/adopt-a-cat")
	fmt.Println(cats)
	return cats
}
func osceola(cats []cat) []cat{
	c := colly.NewCollector()

	c.OnHTML("div[class=pet.col-md-4.col-sm-6]", func(h *colly.HTMLElement){
		cat := cat{
			Name : h.ChildText("div.pet-id.text-center"),
		}
		cats = append(cats,cat)
	})
	c.Visit("https://www.osceolacountypets.com/adoptable-pet-search/")
	fmt.Println(cats)
	return cats
}
func broward(cats []cat) []cat{
	c := colly.NewCollector(colly.AllowedDomains("humanebroward.com"),)

	c.OnHTML("div[class=elementor-widget-wrap elementor-element-populated]", func(h *colly.HTMLElement){
		cat := cat{
			Name : h.ChildText("h3.pet-name"),
			Gender : h.ChildText("div.pet-detail"),
			Age : "",
			Breed : "",
			ImageURL : "",
		}
	SetDefaultInWebsite(&cat)
	cats = append(cats,cat)

	})
	c.Visit("https://humanebroward.com/all-pets/?type=CAT&pg=1")
	fmt.Println(cats)
	return cats
}
func main(){
	var catsItem []cat
	
	
	// catsItem = miamiDade(catsItem)
	// fmt.Println(" ")
	// catsItem = lakeCounty(catsItem)
	// fmt.Println(" ")
	//catsItem = peggy(catsItem)
	// fmt.Println(" ")
	//catsItem = tampaBay(catsItem)
	catsItem = osceola(catsItem)
	fmt.Println(len(catsItem))

}