package vendors

import (
	"os"

	"github.com/tebeka/selenium"
)

var browserstackHosts = os.Getenv("BROWSERSTACK_HOSTS")
var browserstackRunLocal = os.Getenv("BROWSERSTACK_RUN_LOCAL")

// https://www.browserstack.com/list-of-browsers-and-platforms/automate
// device:https://www.browserstack.com/list-of-browsers-and-platforms/automate
// https://www.browserstack.com/automate/capabilities
func GetCapList() []selenium.Capabilities {

	// capList := []selenium.Capabilities{}
	// capList = append(capList, GetWindowsCapList()...)
	// capList = append(capList, GetMacCapList()...)
	// capList = append(capList, GetAndroidCapList()...)
	// capList = append(capList, GetIosCapList()...)
	// return capList

	return GetWindowsCapList()
	// return GetMacCapList()
	// return GetAndroidCapList()
	// return GetIosCapList()
}

func GetWindowsCapList() []selenium.Capabilities {
	allWindowsOSVersion := GetAllWindowsOSVersion()
	return GetDesktopCapList(allWindowsOSVersion, WindowsOs)
}

func GetMacCapList() []selenium.Capabilities {
	allMacOSVersion := GetAllMacOSVersion()
	return GetDesktopCapList(allMacOSVersion, MacOs)
}

func GetIosCapList() []selenium.Capabilities {
	allIosDevice := GetAllIosDevice()
	allIosOSVersion := GetAllIosOSVersion()
	return GetMobileCaplist(allIosDevice, allIosOSVersion, IosBrowserName, IosRealMobile)
}

func GetAndroidCapList() []selenium.Capabilities {
	allAndroidDevice := GetAllAndroidDevice()
	allAndroidOSVersion := GetAllAndroidOSVersion()
	return GetMobileCaplist(allAndroidDevice, allAndroidOSVersion, AndroidBrowserName, AndroidRealMobile)
}

func GetDesktopCapList(allOSVerion []string, os string) []selenium.Capabilities {
	var capList []selenium.Capabilities

	allBrowserName := GetAllBrowserName()

	for _, browserName := range allBrowserName {
		switch browserName {
		case BrowserFirefox:
			allBrowserVersion := GetAllBrowserVersionFirefox()
			capList = append(capList, GetDesktopCapListCore(allOSVerion, allBrowserVersion, os, browserName)...)
		case BrowserChrome:
			allBrowserVersion := GetAllBrowserVersionChrome()
			capList = append(capList, GetDesktopCapListCore(allOSVerion, allBrowserVersion, os, browserName)...)
		case BrowserIE:
			allBrowserVersion := GetAllBrowserVersionIE()
			capList = append(capList, GetDesktopCapListCore(allOSVerion, allBrowserVersion, os, browserName)...)
		case BrowserOpera:
			allBrowserVersion := GetAllBrowserVersionOpera()
			capList = append(capList, GetDesktopCapListCore(allOSVerion, allBrowserVersion, os, browserName)...)
		case BrowserSafari:
			allBrowserVersion := GetAllBrowserVersionSafari()
			capList = append(capList, GetDesktopCapListCore(allOSVerion, allBrowserVersion, os, browserName)...)
		default:
		}
	}

	return capList
}

func GetDesktopCapListCore(allOSVersion, allBrowserVersion []string, os, browserName string) []selenium.Capabilities {
	var capList []selenium.Capabilities
	for _, osVersion := range allOSVersion {
		for _, browserVersion := range allBrowserVersion {
			cap := selenium.Capabilities{
				"build":               os,
				"name":                "testing",
				"project":             "auto-test",
				"browserstack.tunnel": "true",
				"browserstack.local":  browserstackRunLocal,
				"browserstack.hosts":  browserstackHosts,
			}
			cap["os"] = os
			cap["os_version"] = osVersion
			cap["browserName"] = browserName
			cap["browserVersion"] = browserVersion
			capList = append(capList, cap)
		}
	}
	return capList
}

func GetMobileCaplist(allDevice, allOSVersion []string, browserName, realMobile string) []selenium.Capabilities {
	var capList []selenium.Capabilities
	for _, device := range allDevice {
		for _, osVersion := range allOSVersion {
			cap := selenium.Capabilities{
				"build":               browserName,
				"name":                "testing",
				"project":             "auto-test",
				"browserstack.tunnel": "true",
				"browserstack.local":  browserstackRunLocal,
				"browserstack.hosts":  browserstackHosts,
			}
			cap["os_version"] = osVersion
			cap["browserName"] = browserName
			cap["realMobile"] = realMobile
			cap["device"] = device
			capList = append(capList, cap)
		}
	}
	return capList
}
