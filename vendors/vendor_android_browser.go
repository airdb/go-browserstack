package vendors

const (
	AndroidBrowserName = "Android"
	AndroidRealMobile  = "true"

	AndroidOSVersion4_4    = "4.4"
	AndroidOSVersion5_0    = "5.0"
	AndroidOSVersion5_1    = "5.1"
	AndroidOSVersion6_0    = "6.0"
	AndroidOSVersion7_0    = "7.0"
	AndroidOSVersion7_1    = "7.1"
	AndroidOSVersion8_0    = "8.0"
	AndroidOSVersion8_1    = "8.1"
	AndroidOSVersion9_0    = "9.0"
	AndroidOSVersion10_0   = "10.0"
	AndroidOSVersion11_0   = "11.0"
	AndroidOSVersion12Beta = "12 Beta"
)

// Galaxy devices
const (
	AndroidDeviceGalaxyA10      = "Samsung Galaxy A10"
	AndoirdDeviceGalaxyA11      = "Samsung Galaxy A11"
	AndoirdDeviceGalaxyA51      = "Samsung Galaxy A51"
	AndoirdDeviceGalaxyA8       = "Samsung Galaxy A8"
	AndoirdDeviceGalaxyJ7Prime  = "Samsung Galaxy J7 Prime"
	AndoirdDeviceGalaxyNote10   = "Samsung Galaxy Note 10"
	AndoirdDeviceGalaxy10Plus   = "Samsung Galaxy Note 10 Plus"
	AndoirdDeviceGalaxyNote20   = "Samsung Galaxy Note 20"
	AndoirdDeviceGalaxy20Ultra  = "Samsung Galaxy Note 20 Ultra"
	AndoirdDeviceGalaxyNote8    = "Samsung Galaxy Note 8"
	AndoirdDeviceGalaxyNote9    = "Samsung Galaxy Note 9"
	AndoirdDeviceGalaxyS10      = "Samsung Galaxy S10"
	AndoirdDeviceGalaxyS10Plus  = "Samsung Galaxy S10 Plus"
	AndoirdDeviceGalaxyS10e     = "Samsung Galaxy S10e"
	AndoirdDeviceGalaxyS20      = "Samsung Galaxy S20"
	AndoirdDeviceGalaxyS20Plus  = "Samsung Galaxy S20 Plus"
	AndoirdDeviceGalaxyS20Ultra = "Samsung Galaxy S20 Ultra"
	AndoirdDeviceGalaxyS21      = "Samsung Galaxy S21"
	AndoirdDeviceGalaxyS21Plus  = "Samsung Galaxy S21 Plus"
	AndoirdDeviceGalaxyS21Ultra = "Samsung Galaxy S21 Ultra"
	AndoirdDeviceGalaxyS6       = "Samsung Galaxy S6"
	AndoirdDeviceGalaxyS7       = "Samsung Galaxy S7"
	AndoirdDeviceGalaxyS8       = "Samsung Galaxy S8"
	AndoirdDeviceGalaxyS8Plus   = "Samsung Galaxy S8 Plus"
	AndoirdDeviceGalaxyS9       = "Samsung Galaxy S9"
	AndoirdDeviceGalaxyS9Plus   = "Samsung Galaxy S9 Plus"
	AndoirdDeviceGalaxyTab4     = "Samsung Galaxy Tab 4"
	AndoirdDeviceGalaxyTabS3    = "Samsung Galaxy Tab S3"
	AndoirdDeviceGalaxyTabS4    = "Samsung Galaxy Tab S4"
	AndoirdDeviceGalaxyTabS5e   = "Samsung Galaxy Tab S5e"
	AndoirdDeviceGalaxyTabS6    = "Samsung Galaxy Tab S6"
	AndoirdDeviceGalaxyTabS7    = "Samsung Galaxy Tab S7"
)

// Google devices
const (
	AndroidDeviceGoogleNexus5    = "Google Nexus 5"
	AndroidDeviceGoogleNexus6    = "Google Nexus 6"
	AndroidDeviceGoogleNexus9    = "Google Nexus 9"
	AndroidDeviceGooglePixel     = "Google Pixel"
	AndroidDeviceGooglePixel2    = "Google Pixel 2"
	AndroidDeviceGooglePixel3    = "Google Pixel 3"
	AndroidDeviceGooglePixel3XL  = "Google Pixel 3 XL"
	AndroidDeviceGooglePixel3a   = "Google Pixel3 a"
	AndroidDeviceGooglePixel3aXL = "Google Pixel3 a XL"
	AndroidDeviceGooglePixel4    = "Google Pixel 4"
	AndroidDeviceGooglePixel4XL  = "Google Pixel 4 XL"
	AndroidDeviceGooglePixel5    = "Google Pixel 5"
)

// OnePlus devices
const (
	AndroidDeviceOnePlus6T = "OnePlus 6T"
	AndroidDeviceOnePlus7  = "OnePlus 7"
	AndroidDeviceOnePlus7T = "OnePlus 7T"
	AndroidDeviceOnePlus8  = "OnePlus 8"
	AndroidDeviceOnePlus9  = "OnePlus 9"
)

// Xiaomi devices
const (
	AndroidDeviceXiaomiRedmiNote7 = "Xiaomi Redmi Note 7"
	AndroidDeviceXiaomiRedmiNote8 = "Xiaomi Redmi Note 8"
	AndroidDeviceXiaomiRedmiNote9 = "Xiaomi Redmi Note 9"
)

// Motorola devices
const (
	AndroidDeviceMotorolaMotoG7Play  = "Motorola Moto G7 Play"
	AndroidDeviceMotorolaMotoG9Play  = "Motorola Moto G9 Play"
	AndroidDeviceMotorolaMotoX2ndGen = "Motorola Moto X 2nd Gen"
)

//Vivo devices
const (
	AndroidDeviceVivoY50 = "Vivo Y50"
)

// Oppo devices
const (
	AndroidDeviceOppoReno3Pro = "Oppo Reno 3 Pro"
)

// Huawei devices
const (
	AndroidDeviceHuaweiP30 = "Huawei P30"
)

// Define android os version list.
func GetAllAndroidOSVersion() []string {
	return []string{
		AndroidOSVersion4_4,
		AndroidOSVersion5_0,
		AndroidOSVersion5_1,
		AndroidOSVersion6_0,
		AndroidOSVersion7_0,
		AndroidOSVersion7_1,
		AndroidOSVersion8_0,
		AndroidOSVersion8_1,
		AndroidOSVersion9_0,
		AndroidOSVersion10_0,
		AndroidOSVersion11_0,
		AndroidOSVersion12Beta,
	}
}

func GetAllAndroidDevice() []string {
	var androidDeviceList []string
	androidDeviceList = append(androidDeviceList, GetAllAndroidGalaxy()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidGoogle()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidOnePlus()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidXiaomi()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidMotorola()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidVivo()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidOppo()...)
	androidDeviceList = append(androidDeviceList, GetAllAndroidHuawei()...)
	return androidDeviceList
}

func GetAllAndroidGalaxy() []string {
	return []string{
		AndroidDeviceGalaxyA10,
		AndoirdDeviceGalaxyA11,
		AndoirdDeviceGalaxyA51,
		AndoirdDeviceGalaxyA8,
		AndoirdDeviceGalaxyJ7Prime,
		AndoirdDeviceGalaxyNote10,
		AndoirdDeviceGalaxy10Plus,
		AndoirdDeviceGalaxyNote20,
		AndoirdDeviceGalaxy20Ultra,
		AndoirdDeviceGalaxyNote8,
		AndoirdDeviceGalaxyNote9,
		AndoirdDeviceGalaxyS10,
		AndoirdDeviceGalaxyS10Plus,
		AndoirdDeviceGalaxyS10e,
		AndoirdDeviceGalaxyS20,
		AndoirdDeviceGalaxyS20Plus,
		AndoirdDeviceGalaxyS20Ultra,
		AndoirdDeviceGalaxyS21,
		AndoirdDeviceGalaxyS21Plus,
		AndoirdDeviceGalaxyS21Ultra,
		AndoirdDeviceGalaxyS6,
		AndoirdDeviceGalaxyS7,
		AndoirdDeviceGalaxyS8,
		AndoirdDeviceGalaxyS8Plus,
		AndoirdDeviceGalaxyS9,
		AndoirdDeviceGalaxyS9Plus,
		AndoirdDeviceGalaxyTab4,
		AndoirdDeviceGalaxyTabS3,
		AndoirdDeviceGalaxyTabS4,
		AndoirdDeviceGalaxyTabS5e,
		AndoirdDeviceGalaxyTabS6,
		AndoirdDeviceGalaxyTabS7,
	}
}

func GetAllAndroidGoogle() []string {
	return []string{
		AndroidDeviceGoogleNexus5,
		AndroidDeviceGoogleNexus6,
		AndroidDeviceGoogleNexus9,
		AndroidDeviceGooglePixel,
		AndroidDeviceGooglePixel2,
		AndroidDeviceGooglePixel3,
		AndroidDeviceGooglePixel3XL,
		AndroidDeviceGooglePixel3a,
		AndroidDeviceGooglePixel3aXL,
		AndroidDeviceGooglePixel4,
		AndroidDeviceGooglePixel4XL,
		AndroidDeviceGooglePixel5,
	}
}

func GetAllAndroidOnePlus() []string {
	return []string{
		AndroidDeviceOnePlus6T,
		AndroidDeviceOnePlus7,
		AndroidDeviceOnePlus7T,
		AndroidDeviceOnePlus8,
		AndroidDeviceOnePlus9,
	}
}

func GetAllAndroidXiaomi() []string {
	return []string{
		AndroidDeviceXiaomiRedmiNote7,
		AndroidDeviceXiaomiRedmiNote8,
		AndroidDeviceXiaomiRedmiNote9,
	}
}

func GetAllAndroidMotorola() []string {
	return []string{
		AndroidDeviceMotorolaMotoG7Play,
		AndroidDeviceMotorolaMotoG9Play,
		AndroidDeviceMotorolaMotoX2ndGen,
	}
}

func GetAllAndroidVivo() []string {
	return []string{
		AndroidDeviceVivoY50,
	}
}

func GetAllAndroidOppo() []string {
	return []string{
		AndroidDeviceOppoReno3Pro,
	}
}

func GetAllAndroidHuawei() []string {
	return []string{
		AndroidDeviceHuaweiP30,
	}
}
