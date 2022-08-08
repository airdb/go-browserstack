package plugin

import (
	"log"

	"github.com/tebeka/selenium"
)

func Visit(wd selenium.WebDriver, url string) {
	log.Println("start visit website...", url)
	wd.Get(url)
	log.Println("visit website finished!")
}
