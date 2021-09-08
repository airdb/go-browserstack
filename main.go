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

    remoteUrl := fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", browserstackUsername, browserstackAccessKey)

    // device:https://www.browserstack.com/list-of-browsers-and-platforms/automate
    // https://www.browserstack.com/automate/capabilities
    // https://www.browserstack.com/docs/automate/selenium/select-browsers-and-devices#selenium-capabilities
    caps := selenium.Capabilities{
        "build":   "auto-test",
        "name":    "Windows testing",
        "project": "auto-test",
        // "browserstack.debug":  "true",
        // "browserstack.local":  "true",
        // "browserstack.tunnel": "true",
        "browserName":        "firefox",
        "browserVersion":     "88.0",
        "os":                 "Windows",
        "os_version":         "8.1",
        "browserstack.hosts": browserstackHosts,
    }

    wd, err := selenium.NewRemote(caps, remoteUrl)
    if err != nil {
        panic(err)
    }
    defer wd.Quit()

    plugin.VisitLocal(wd, browserstackURL)
}
