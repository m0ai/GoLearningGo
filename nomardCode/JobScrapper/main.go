package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/m0ai/JobScrapper/scrapper"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScarpe(c echo.Context) error {

	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScarpe)
	e.Logger.Fatal(e.Start(":1323"))

	scrapper.Scrape("term")
}
