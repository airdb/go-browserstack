package vendors

import (
	"log"
	"os"

	"github.com/tebeka/selenium"
)

func GetCapList() []selenium.Capabilities {
	var capList []selenium.Capabilities

	browserstackHosts := os.Getenv("BROWSERSTACK_HOSTS")
	browserstackRunLocal := os.Getenv("BROWSERSTACK_RUN_LOCAL")

	// device:https://www.browserstack.com/list-of-browsers-and-platforms/automate
	// https://www.browserstack.com/automate/capabilities
	// https://www.browserstack.com/docs/automate/selenium/select-browsers-and-devices#selenium-capabilities

	os := OSWindows
	osVersion := OSWindowsVersion8_1

	browserName := OSWindowsBrowerFirefox
	browserVersion := OSWindowsBrowerFirefoxVersion88_0

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

	capList = append(capList, caps)

	return capList
}

func GetWindowsCapList() {
	GetWindowsFirefoxCapList()
	GetWindowsChromeCapList()
}

func GetWindowsFirefoxCapList() {
}

func GetWindowsChromeCapList() {
}

func GetAndoirdCapList() {
	var capList []selenium.Capabilities
	cap := selenium.Capabilities{
		"build":   "auto-test",
		"name":    "send testing",
		"project": "auto-test",
	}

	for _, osVersion := range AllAndroidOSVersion {
		for _, browserName := range AllAndroidBrowserName {
			cap["os_version"] = osVersion
			cap["browserName"] = browserName
			// log.Println(cap)
			capList = append(capList, cap)
		}
	}

	log.Println(capList)
}

func GetIPadCapList() {
}

func GetIPhoneCaplist() {
}
