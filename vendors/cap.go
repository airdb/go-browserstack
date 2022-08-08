package vendors

import (
	"os"

	"github.com/tebeka/selenium"
)

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

	ret := GetWindowsCapList()
	// return GetMacCapList()
	// return GetAndroidCapList()
	// return GetIosCapList()
	return ret
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

func GetDesktopCapListCore(allOSVersion, allBrowserVersion []string, system, browserName string) []selenium.Capabilities {
	var capList []selenium.Capabilities

	for _, osVersion := range allOSVersion {
		for _, browserVersion := range allBrowserVersion {
			cap := selenium.Capabilities{
				"build":   system,
				"name":    "testing",
				"project": "auto-test",
			}

			browserstackHosts := os.Getenv("BROWSERSTACK_HOSTS")
			browserstackRunLocal := os.Getenv("BROWSERSTACK_RUN_LOCAL")
			if browserstackRunLocal == "true" {
				cap["browserstack.tunnel"] = "ture"
				cap["browserstack.local"] = browserstackRunLocal
				cap["browserstack.host"] = browserstackHosts
			}

			cap["os"] = system
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
				"build":   browserName,
				"name":    "testing",
				"project": "auto-test",
			}

			browserstackHosts := os.Getenv("BROWSERSTACK_HOSTS")
			browserstackRunLocal := os.Getenv("BROWSERSTACK_RUN_LOCAL")
			if browserstackRunLocal == "true" {
				cap["browserstack.tunnel"] = "ture"
				cap["browserstack.local"] = browserstackRunLocal
				cap["browserstack.host"] = browserstackHosts
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
