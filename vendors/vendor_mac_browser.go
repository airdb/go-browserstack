package vendors

const (
	MacOs = "OS X"

	MacOsVersionBigSur       = "Big Sur"
	MacOsVersionCatalina     = "Catalina"
	MacOsVersionMojave       = "Mojave"
	MacOsVersionHighSierra   = "High Sierra"
	MacOsVersionSierra       = "Sierra"
	MacOsVersionElCapitan    = "El Capitan"
	MacOsVersionYosemite     = "Yosemite"
	MacOsVersionMavericks    = "Mavericks"
	MacOsVersionMountainLion = "Mountain Lion"
	MacOsVersionLion         = "Lion"
	MacOsVersionSnowLeopard  = "Snow Leopard"
)

func GetAllMacOSVersion() []string {
	return []string{
		MacOsVersionBigSur,
		MacOsVersionCatalina,
		MacOsVersionMojave,
		MacOsVersionHighSierra,
		MacOsVersionSierra,
		MacOsVersionElCapitan,
		MacOsVersionYosemite,
		MacOsVersionMavericks,
		MacOsVersionMountainLion,
		MacOsVersionLion,
		MacOsVersionSnowLeopard,
	}
}
