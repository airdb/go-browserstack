package plugin

import (
	"log"

	"github.com/tebeka/selenium"
)

func VisitLocal(wd selenium.WebDriver, url string) {
	log.Println("start visit website...", url)
	wd.Get(url)
	log.Println("visit website finished!")
}
