package vendors

const (
	WindowsOs = "Windows"

	WindowsOsVersion7   = "7"
	WindowsOsVersion8   = "8"
	WindowsOsVersion8_1 = "8.1"
	WindowsOsVersion10  = "10"
	WindowsOsVersionXP  = "XP"
)

func GetAllWindowsOSVersion() []string {
	return []string{
		WindowsOsVersion7,
		WindowsOsVersion8,
		WindowsOsVersion8_1,
		WindowsOsVersion10,
		WindowsOsVersionXP,
	}
}
