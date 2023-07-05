package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var chosen string
	if option1 < option2 {
		chosen = option1
	} else {
		chosen = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", chosen)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age >= 10 {
		return originalPrice * 0.5
	} else if age > 5 {
		return originalPrice * 0.7
	} else {
		return originalPrice * 0.8
	}
}
