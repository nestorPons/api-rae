package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

type Def struct {
	Id  int    `json:"id"`
	Def string `json:"data"`
}

type Deft struct {
	Title string
	Data  []Def
}

type Defs []Deft

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/buscar/:w", func(c *gin.Context) {
		word := c.Param("w")
		url := "https://dle.rae.es/" + word
		res, err := http.Get(url)
		// If url could not be opened, we inform the channel chFailedUrls:
		chkerror200(res, err)
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		chkerror(err)

		defs := Defs{}
		deft := Deft{}

		doc.Find("#resultados p").Each(func(i int, s *goquery.Selection) {

			d := s.Text()
			r := "^[0-9]{1}"
			b, e := regexp.MatchString(r, d)
			chkerror(e)

			if !b {
				if i > 0 {
					print("nuevo")
					defs = append(defs, deft)
					deft = Deft{}
				}
				deft.Title = d
			} else {
				d := Def{i, d}
				deft.Data = append(deft.Data, d)
			}

		})
		c.JSON(200, defs)
	})

	router.Run(":" + port)
}

func chkerror(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func chkerror200(res *http.Response, err error) {
	if err != nil || res.StatusCode != 200 {
		fmt.Println(err)
	}
}
