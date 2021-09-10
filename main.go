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
	browserstackHosts := os.Getenv("BROWSERSTACK_HOSTS")
	browserstackRunLocal := os.Getenv("BROWSERSTACK_RUN_LOCAL")

	remoteUrl := fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", browserstackUsername, browserstackAccessKey)

	// device:https://www.browserstack.com/list-of-browsers-and-platforms/automate
	// https://www.browserstack.com/automate/capabilities
	// https://www.browserstack.com/docs/automate/selenium/select-browsers-and-devices#selenium-capabilities

	browserName := "firefox"
	browserVersion := "88.0"
	os := "Windows"
	osVersion := "8.1"
	realMobile := "false"
	device := ""

	caps := selenium.Capabilities{
		"build":              "auto-test",
		"name":               "send testing",
		"project":            "auto-test",
		"browserstack.local": browserstackRunLocal,
		// "browserstack.debug":  "true",
		// "browserstack.tunnel": "true",
		"browserstack.hosts": browserstackHosts,
		"browserName":        browserName,    // for pc
		"browserVersion":     browserVersion, // for pc
		"os":                 os,             // for pc
		"os_version":         osVersion,      // for pc and mobile
		"RealMobile":         realMobile,     // for mobile
		"Device":             device,         // for mobile
	}

	wd, err := selenium.NewRemote(caps, remoteUrl)
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	requestURL := fmt.Sprintf("%s?browser_name=%s&browser_version=%s&os=%s&os_version=%s&real_mobile=%s&device=%s", browserstackURL,
		browserName,
		browserVersion,
		os,
		osVersion,
		realMobile,
		device)

	plugin.VisitLocal(wd, requestURL)
}
