package plugin

import (
	"log"

	"github.com/tebeka/selenium"
)

func Visit(wd selenium.WebDriver, url string) {
	log.Println("start visit website...")
	wd.Get("https://www.baidu.com")
	log.Println("visit website finished!")
}
