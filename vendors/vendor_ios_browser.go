package vendors

const (
	IosBrowserName = "iPhone"
	IosRealMobile  = "true"

	IosOSVersion10     = "10"
	IosOSVersion11     = "11"
	IosOSVersion12     = "12"
	IosOSVersion13     = "13"
	IosOSVersion14     = "14"
	IosOSVersion15Beta = "15 Beta"
)

// Iphone devices
const (
	IosDeviceIphone11       = "iPhone 11"
	IosDeviceIphone11Pro    = "iPhone 11 Pro"
	IosDeviceIphone11ProMax = "iPhone 11 Pro Max"
	IosDeviceIphone12       = "iPhone 12"
	IosDeviceIphone12Mini   = "iPhone 12 Mini"
	IosDeviceIphone12Pro    = "iPhone 12 Pro"
	IosDeviceIphone12ProMax = "iPhone 12 Pro Max"
	IosDeviceIphone6        = "iPhone 6"
	IosDeviceIphone6S       = "iPhone 6S"
	IosDeviceIphone6SPlus   = "iPhone 6S Plus"
	IosDeviceIphone7        = "iPhone 7"
	IosDeviceIphone7Plus    = "iPhone 7 Plus"
	IosDeviceIphone8        = "iPhone 8"
	IosDeviceIphone8Plus    = "iPhone 8 Plus"
	IosDeviceIphoneSE       = "iPhone SE"
	IosDeviceIphoneSE2020   = "iPhone SE 2020"
	IosDeviceIphoneX        = "iPhone X"
	IosDeviceIphoneXR       = "iPhone XR"
	IosDeviceIphoneXS       = "iPhone XS"
	IosDeviceIphoneXSMax    = "iPhone XS Max"
)

// Ipad devices
const (
	IosDevicesIpad5th         = "iPad 5th"
	IosDevicesIpad6th         = "iPad 6th"
	IosDevicesIpad7th         = "iPad 7th"
	IosDevicesIpad8th         = "iPad 8th"
	IosDevicesIpadAir2019     = "iPad Air 2019"
	IosDevicesIpadAir4        = "iPad Air 4"
	IosDevicesIpadMini2019    = "iPad Mini 2019"
	IosDevicesIpadMini4       = "iPad Mini 4"
	IosDevicesIpadPro112018   = "iPad Pro 11 2018"
	IosDevicesIpadPro112020   = "iPad Pro 11 2020"
	IosDevicesIpadPro112021   = "iPad Pro 11 2021"
	IosDevicesIpadPro12_92017 = "iPad Pro 12.9 2017"
	IosDevicesIpadPro12_92018 = "iPad Pro 12.9 2018"
	IosDevicesIpadPro12_92020 = "iPad Pro 12.9 2020"
	IosDevicesIpadPro12_92021 = "iPad Pro 12.9 2021"
	IosDevicesIpadPro9_72016  = "iPad Pro 9.7 2016"
)

func GetAllIosOSVersion() []string {
	return []string{
		IosOSVersion10,
		IosOSVersion11,
		IosOSVersion12,
		IosOSVersion13,
		IosOSVersion14,
		IosOSVersion15Beta,
	}
}

func GetAllIosDevice() []string {
	var iosDeviceList []string
	iosDeviceList = append(iosDeviceList, GetAllIphoneDevice()...)
	iosDeviceList = append(iosDeviceList, GetAllIpadDevice()...)
	return iosDeviceList
}

func GetAllIphoneDevice() []string {
	return []string{
		IosDeviceIphone11,
		IosDeviceIphone11Pro,
		IosDeviceIphone11ProMax,
		IosDeviceIphone12,
		IosDeviceIphone12Mini,
		IosDeviceIphone12Pro,
		IosDeviceIphone12ProMax,
		IosDeviceIphone6,
		IosDeviceIphone6S,
		IosDeviceIphone6SPlus,
		IosDeviceIphone7,
		IosDeviceIphone7Plus,
		IosDeviceIphone8,
		IosDeviceIphone8Plus,
		IosDeviceIphoneSE,
		IosDeviceIphoneSE2020,
		IosDeviceIphoneX,
		IosDeviceIphoneXR,
		IosDeviceIphoneXS,
		IosDeviceIphoneXSMax,
	}
}

func GetAllIpadDevice() []string {
	return []string{
		IosDevicesIpad5th,
		IosDevicesIpad6th,
		IosDevicesIpad7th,
		IosDevicesIpad8th,
		IosDevicesIpadAir2019,
		IosDevicesIpadAir4,
		IosDevicesIpadMini2019,
		IosDevicesIpadMini4,
		IosDevicesIpadPro112018,
		IosDevicesIpadPro112020,
		IosDevicesIpadPro112021,
		IosDevicesIpadPro12_92017,
		IosDevicesIpadPro12_92018,
		IosDevicesIpadPro12_92020,
		IosDevicesIpadPro12_92021,
		IosDevicesIpadPro9_72016,
	}
}
