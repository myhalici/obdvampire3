package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

func hodo(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	url := "http://www.outilsobdfacile.com/location-plug-connector-obd/Acura"
	resp, _ := soup.Get(url)
	doc := soup.HTMLParse(resp)
	brandsslice := doc.Find("select", "name", "marque").FindAll("option")
	for _, i := range brandsslice {
		//fmt.Println("Brand: " + i.Text())
		url := "http://www.outilsobdfacile.com/location-plug-connector-obd/" + i.Text()
		resp, _ := soup.Get(url)
		doc := soup.HTMLParse(resp)
		modelsslice := doc.Find("select", "name", "modele").FindAll("option")
		for _, y := range modelsslice {
			//fmt.Println(i.Text() + " : " + y.Text())
			url := "http://www.outilsobdfacile.com/location-plug-connector-obd/" + i.Text() + "-" + y.Attrs()["value"]
			resp, _ := soup.Get(url)
			doc := soup.HTMLParse(resp)
			locationslice := doc.FindAll("p", "class", "legende_connecteur")
			fmt.Printf(i.Text() + " : " + y.Text() + " -> ")
			for _, z := range locationslice {
				fmt.Printf(z.Text() + ",")
			}
			fmt.Println("")
			photoslice := doc.Find("div", "class", "paragraphe_content").FindAll("img", "class", "connecteur")
			for _, t := range photoslice {
				fmt.Println(strings.Replace(t.Attrs()["src"], "../", "http://www.outilsobdfacile.com/", -1))
			}
		}
	}

}
