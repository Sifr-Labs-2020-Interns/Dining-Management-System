package main

import (
	"encoding/csv"
	"fmt"
	"github.com/go-macaron/oauth2"
	"github.com/go-macaron/session"
	goauth2 "golang.org/x/oauth2"
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
	m.Use(session.Sessioner())
	m.Use(oauth2.Google(
		&goauth2.Config{
			ClientID:     "client_id",
			ClientSecret: "client_secret",
			Scopes:       []string{"https://www.googleapis.com/auth/drive"},
			RedirectURL:  "redirect_url",
		},
	))

	// Tokens are injected to the handlers
	m.Get("/", func(tokens oauth2.Tokens) string {
		if tokens.Expired() {
			return "not logged in, or the access token is expired"
		}
		return "logged in"
	})

	// Routes that require a logged in user
	// can be protected with oauth2.LoginRequired handler.
	// If the user is not authenticated, they will be
	// redirected to the login path.
	m.Get("/restrict", oauth2.LoginRequired, func(tokens oauth2.Tokens) string {
		return tokens.Access()
	})
	
	m.Get("/", func() string {
		return "Hello, Go!"

	})

	m.Run()

}
