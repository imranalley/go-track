package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
	if (kind == "car" || kind == "truck") {
		license := true
		return license
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")
	decision := ""
	if option1 > option2{
		decision += option2 + " is clearly the better choice."
	} else{
		decision += option1 + " is clearly the better choice."
	}
	return (decision)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")
	if (age<3){
		return(originalPrice*.8)
	} else if (age>3 && age <10){
		return(originalPrice*.7)	
	} else {
		return(originalPrice*.5)
	}
}
