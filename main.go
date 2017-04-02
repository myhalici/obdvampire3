package main

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/tealeg/xlsx"
)

func hodo(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var carcounter = 0
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("obdVampire")
	hodo(err)
	row = sheet.AddRow()
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
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = i.Text()
			cell = row.AddCell()
			cell.Value = y.Text()
			cell = row.AddCell()
			cell.Value = ""
			cell = row.AddCell()
			cell.Value = ""
			cell = row.AddCell()
			cell.Value = ""
			var tmpz = ""
			for _, z := range locationslice {
				fmt.Printf(z.Text() + ",")
				tmpz = tmpz + "," + z.Text()
			}
			cell = row.AddCell()
			cell.Value = tmpz
			fmt.Println("")
			photoslice := doc.Find("div", "class", "paragraphe_content").FindAll("img", "class", "connecteur")
			var tmpt = ""
			for _, t := range photoslice {
				fmt.Println(strings.Replace(t.Attrs()["src"], "../", "http://www.outilsobdfacile.com/", -1))
				tmpt = tmpt + "," + strings.Replace(t.Attrs()["src"], "../", "http://www.outilsobdfacile.com/", -1)
			}
			cell = row.AddCell()
			cell.Value = tmpt
			cell = row.AddCell()
			cell.Value = ""
			carcounter++
			fmt.Println("Car count: ", carcounter)

			err = file.Save("obdVampire3.xlsx")
			hodo(err)
		}
	}

}
