package main

import (
	"fmt"
	"strings"

	// "math/rand"
	// "time"

	"github.com/gocolly/colly"
)

type Cat struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Breed    string `json:"breed"`
	Age      string `json:"age"`
	ImageURL string `json:"imageUrl"`
	//Feature string `json:"feature"`
}
var features = []string{"friendly", "playful", "adorable", "loving", "affectionate", "curious", "energetic", "loves attention", "loves to cuddle", "loves to play", "loves to be pet"}


// make a default value since many websites do not have all the information for the cat object
func SetDefaultUnknown(catDefault *Cat) {
	if catDefault.Name == "" {
		catDefault.Name = "Unknown"
	}
	if catDefault.Age == "" {
		catDefault.Age = "Unknown"
	}
	if catDefault.Gender == "" {
		catDefault.Gender = "Unknown"
	}
	if catDefault.Breed == "" {
		catDefault.Breed = "Unknown"
	}
	if catDefault.ImageURL == "" {
		catDefault.ImageURL = "Unknown"
	}
}
func SetDefaultInWebsite(catWebDefault *Cat) {
	if catWebDefault.Name == "" {
		catWebDefault.Name = "On Website"
	}
	if catWebDefault.Age == "" {
		catWebDefault.Age = "On Website"
	}
	if catWebDefault.Gender == "" {
		catWebDefault.Gender = "On Website"
	}
	if catWebDefault.Breed == "" {
		catWebDefault.Breed = "On Website"
	}
	if catWebDefault.ImageURL == "" {
		catWebDefault.ImageURL = "On Website"
	}
}
func MiamiDade(cats []Cat) []Cat {
	//var catz []Cat
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"))

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement) { //gets the elements from the gridResult
		// rand.Seed(time.Now().UnixNano())
        // ranFeature := rand.Intn(len(features))
		cat := Cat{
			Name:     h.ChildText("span.text_Name.results"),
			Gender:   h.ChildText("span.text_Gender.results"),
			Breed:    h.ChildText("span.text_Breed.results"),
			Age:      h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),
			//Feature: features[ranFeature],

		}
		cats = append(cats, cat) // adds to the vector of cat objects

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

func LakeCounty(cats []Cat) []Cat {

	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"))

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement) { //makes the cat object from the tags on the HTML page
		// rand.Seed(time.Now().UnixNano())
        // ranFeature := rand.Intn(len(features))
		cat := Cat{
			Name:     h.ChildText("span.text_Name.results"),
			Gender:   h.ChildText("span.text_Gender.results"),
			Breed:    "",
			Age:      h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),
			//Feature: features[ranFeature],

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat) //Adds them into the vector of cat objects
	})

	c.Visit("https://24petconnect.com/LakeCountyAdoptablePets?at=CAT&sb=id_asc")
	fmt.Println(cats)

	return cats
}

func Peggy(cats []Cat) []Cat {
	c := colly.NewCollector()

	c.OnHTML("div[class=animal-select]", func(h *colly.HTMLElement) {
		// rand.Seed(time.Now().UnixNano())
        // ranFeature := rand.Intn(len(features))
		cat := Cat{
			Name: h.ChildText("div.animal-name.text-center"),
			// Gender : h.ChildText("div.list-animal-sexSN"),
			// Breed : h.ChildText("div.list-animal-breed"),
			// Age : h.ChildText("div.list-animal-age"),
			ImageURL: h.ChildAttr("img", "src"),
			//Feature: features[ranFeature],

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)

	})
	c.Visit("https://www.peggyadams.org/adopt-a-cat")
	fmt.Println(cats)

	return cats
}
func KeyWest(cats []Cat) []Cat {
	c := colly.NewCollector()

	c.OnHTML("td[class=list-item]", func(h *colly.HTMLElement) {
		// rand.Seed(time.Now().UnixNano())
        // ranFeature := rand.Intn(len(features))
		cat := Cat{
			Name:     h.ChildText("div.list-animal-name"),
			Gender:   h.ChildText("div.list-animal-sexSN"),
			Breed:    h.ChildText("div.list-animal-breed"),
			Age:      "",
			ImageURL: h.ChildAttr("img", "src"),
			//Feature: features[ranFeature],

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)
	})
	c.Visit("https://ws.petango.com/webservices/adoptablesearch/wsAdoptableAnimals.aspx?species=Cat&sex=A&agegroup=All&location=&site=00002831&onhold=A&orderby=ID&colnum=3&css=https://fkspca.org/wp-content/themes/FKSPCA/iframe.css&authkey=roapr67swoaniu7rlutho7wlesoukier1udriaslezoa0rleyl&recAmount=&detailsInPopup=No&featuredPet=Include&stageID=")
	fmt.Println(cats)

	return cats
}
func Marathon(cats []Cat) []Cat {
	c := colly.NewCollector()
	c.OnHTML("td[class=list-item]", func(h *colly.HTMLElement) {
		// rand.Seed(time.Now().UnixNano())
        // ranFeature := rand.Intn(len(features))
		cat := Cat{
			Name:     h.ChildText("div.list-animal-name"),
			Gender:   h.ChildText("div.list-animal-sexSN"),
			Breed:    h.ChildText("div.list-animal-breed"),
			Age:      "",
			ImageURL: h.ChildAttr("img", "src"),
			//Feature: features[ranFeature],

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)
	})
	c.Visit("https://ws.petango.com/webservices/adoptablesearch/wsAdoptableAnimals.aspx?species=Cat&sex=A&agegroup=All&location=&site=00002832&onhold=A&orderby=ID&colnum=3&css=https://fkspca.org/wp-content/themes/FKSPCA/iframe.css&authkey=roapr67swoaniu7rlutho7wlesoukier1udriaslezoa0rleyl&recAmount=&detailsInPopup=No&featuredPet=Include&stageID=")
	fmt.Println(cats)

	return cats
}

// ------------------------Lower functions are same as top but converting results into strings------------------------//
func miamiDadeString() string {
	var cats []Cat
	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"))

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement) { //gets the elements from the gridResult
		
		cat := Cat{
			Name:     h.ChildText("span.text_Name.results"),
			Gender:   h.ChildText("span.text_Gender.results"),
			Breed:    h.ChildText("span.text_Breed.results"),
			Age:      h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),

		}
		cats = append(cats, cat) // adds to the vector of cat objects

	})
	c.Visit("https://24petconnect.com/miad?at=CAT")
	//fmt.Println(cats)
	var builder strings.Builder
	for _, cat := range cats {
		builder.WriteString(fmt.Sprintf("{%v %v %v %v %v} ", cat.Name, cat.Gender, cat.Breed, cat.Age, cat.ImageURL))
	}
	result := builder.String()
	return result
}
func lakeCountyString() string {
	var cats []Cat

	c := colly.NewCollector(colly.AllowedDomains("24petconnect.com"))

	c.OnHTML("div[class=gridResult]", func(h *colly.HTMLElement) { //makes the cat object from the tags on the HTML page
		
		cat := Cat{
			Name:     h.ChildText("span.text_Name.results"),
			Gender:   h.ChildText("span.text_Gender.results"),
			Breed:    "",
			Age:      h.ChildText("span.text_Age.results"),
			ImageURL: h.ChildAttr("img", "src"),

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat) //Adds them into the vector of cat objects
	})

	c.Visit("https://24petconnect.com/LakeCountyAdoptablePets?at=CAT&sb=id_asc")
	fmt.Println(cats)
	var builder strings.Builder
	for _, cat := range cats {
		builder.WriteString(fmt.Sprintf("{%v %v %v %v %v} ", cat.Name, cat.Gender, cat.Breed, cat.Age, cat.ImageURL))
	}
	result := builder.String()
	return result
}
func peggyString() string {
	var cats []Cat

	c := colly.NewCollector()

	c.OnHTML("div[class=animal-select]", func(h *colly.HTMLElement) {
		
		cat := Cat{
			Name: h.ChildText("div.animal-name.text-center"),
			// Gender : h.ChildText("div.list-animal-sexSN"),
			// Breed : h.ChildText("div.list-animal-breed"),
			// Age : h.ChildText("div.list-animal-age"),
			ImageURL: h.ChildAttr("img", "src"),
			
		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)

	})
	c.Visit("https://www.peggyadams.org/adopt-a-cat")
	fmt.Println(cats)
	var builder strings.Builder
	for _, cat := range cats {
		builder.WriteString(fmt.Sprintf("{%v %v %v %v %v} ", cat.Name, cat.Gender, cat.Breed, cat.Age, cat.ImageURL))
	}
	result := builder.String()
	return result
}
func keyWestString() string {
	var cats []Cat

	c := colly.NewCollector()

	c.OnHTML("td[class=list-item]", func(h *colly.HTMLElement) {
		
		cat := Cat{
			Name:     h.ChildText("div.list-animal-name"),
			Gender:   h.ChildText("div.list-animal-sexSN"),
			Breed:    h.ChildText("div.list-animal-breed"),
			Age:      "",
			ImageURL: h.ChildAttr("img", "src"),
			

		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)
	})
	c.Visit("https://ws.petango.com/webservices/adoptablesearch/wsAdoptableAnimals.aspx?species=Cat&sex=A&agegroup=All&location=&site=00002831&onhold=A&orderby=ID&colnum=3&css=https://fkspca.org/wp-content/themes/FKSPCA/iframe.css&authkey=roapr67swoaniu7rlutho7wlesoukier1udriaslezoa0rleyl&recAmount=&detailsInPopup=No&featuredPet=Include&stageID=")
	fmt.Println(cats)
	var builder strings.Builder
	for _, cat := range cats {
		builder.WriteString(fmt.Sprintf("{%v %v %v %v %v} ", cat.Name, cat.Gender, cat.Breed, cat.Age, cat.ImageURL))
	}
	result := builder.String()
	return result
}
func marathonString() string {
	var cats []Cat

	c := colly.NewCollector()
	c.OnHTML("td[class=list-item]", func(h *colly.HTMLElement) {
		
		cat := Cat{
			Name:     h.ChildText("div.list-animal-name"),
			Gender:   h.ChildText("div.list-animal-sexSN"),
			Breed:    h.ChildText("div.list-animal-breed"),
			Age:      "",
			ImageURL: h.ChildAttr("img", "src"),
			
		}
		SetDefaultInWebsite(&cat)
		cats = append(cats, cat)
	})
	c.Visit("https://ws.petango.com/webservices/adoptablesearch/wsAdoptableAnimals.aspx?species=Cat&sex=A&agegroup=All&location=&site=00002832&onhold=A&orderby=ID&colnum=3&css=https://fkspca.org/wp-content/themes/FKSPCA/iframe.css&authkey=roapr67swoaniu7rlutho7wlesoukier1udriaslezoa0rleyl&recAmount=&detailsInPopup=No&featuredPet=Include&stageID=")
	fmt.Println(cats)
	var builder strings.Builder
	for _, cat := range cats {
		builder.WriteString(fmt.Sprintf("{%v %v %v %v %v} ", cat.Name, cat.Gender, cat.Breed, cat.Age, cat.ImageURL))
	}
	result := builder.String()
	return result
}

func main() {
	var catsItem []Cat

	// catsItem = MiamiDade(catsItem)
	// fmt.Println(" ")
	// catsItem = LakeCounty(catsItem)
	// fmt.Println(" ")
	// catsItem = Peggy(catsItem)
	// fmt.Println(" ")
	// catsItem = Marathon(catsItem)
	// fmt.Println(" ")
	catsItem=KeyWest(catsItem)
	//fmt.Println(len(catsItem))

}
