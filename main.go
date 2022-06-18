package main

import (
	"fmt"
	"log"
	"os"

	"github.com/airdb/go-browserstack/plugin"
	"github.com/airdb/go-browserstack/vendors"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	browserstackUsername := os.Getenv("BROWSERSTACK_USERNAME")
	browserstackAccessKey := os.Getenv("BROWSERSTACK_ACCESS_KEY")
	browserstackURL := os.Getenv("BROWSERSTACK_URL")

	remoteUrl := fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", browserstackUsername, browserstackAccessKey)

	capList := vendors.GetCapList()
	for _, cap := range capList {
		log.Println(cap["os"], "/", cap["os_version"], "/", cap["browserName"], "/", cap["browserVersion"])
		// log.Println(cap["device"], "/", cap["os_version"])
		wd, err := selenium.NewRemote(cap, remoteUrl)
		if err != nil {
			log.Println(err)
			continue
		}
		requestURL := fmt.Sprintf("%s?browser_name=%s&browser_version=%s&os=%s&os_version=%s&real_mobile=%s&device=%s",
			browserstackURL,
			cap["browserName"],
			cap["browserVersion"],
			cap["os"],
			cap["os_version"],
			cap["realMobile"],
			cap["device"],
		)

		plugin.VisitLocal(wd, requestURL)
		wd.Quit()
	}

}
