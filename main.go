package main

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/macaron.v1"
	"io"
	"log"
	"os"
)



type dish struct {
	ProductID string `json:"product_id"`
	ItemName  string `json:"item_name"`
	Category  string `json:"category"`
	Price     string `json:"price"`
	Veg       string `json:"veg"`
	Available string `json:"available"`

}

func getMenuFromCSV(filepath string) []dish {
	csvFile, err := os.Open(filepath)
	if err != nil{
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	var menuItems []dish

	for {
		line, error := reader.Read()
		if error == io.EOF{
			break
		} else if error != nil {
			log.Fatal(error)
		}

		menuItems = append(menuItems, dish{
			ProductID: line[0],
			ItemName:  line[1],
			Category:  line[2],
			Price:     line[3],
			Veg:       line[4],
			Available: line[5],
		})
	}

	for _, element:= range menuItems{
		fmt.Printf("ID: %s	|	Price: AED%s 	|	Item Name: %s	\n", element.ProductID, element.Price, element.ItemName)
	}

	return menuItems

}


func main() {

	menuPath:= "Menu.csv"
	getMenuFromCSV(menuPath)

	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello, Go!"

	})
	m.Run()
}
