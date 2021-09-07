package main

import (
	"fmt"
	"log"
	"os"

	"github.com/airdb/go-browserstack/plugin"

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

	// https://www.browserstack.com/automate/capabilities
	//	"browserstack.local": "true",
	caps := selenium.Capabilities{
		"project":        "auto-test",
		"browserName":    "firefox",
		"browserVersion": "88.0",
		"os":             "Windows",
		"os_version":     "8.1",
	}
	wd, err := selenium.NewRemote(caps, remoteUrl)
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	plugin.Visit(wd, browserstackURL)
}
