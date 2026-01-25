package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	// panic("Please implement the NewCar function")
	return Car {
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	// panic("Please implement the NewTrack function")
	return Track {
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// panic("Please implement the Drive function")
	var remainingBattery int
	var remainingDistance int
	if car.battery - car.batteryDrain < 0 {
		remainingBattery = car.battery
		remainingDistance = car.distance
	} else {
		remainingBattery = car.battery - car.batteryDrain
		remainingDistance = car.distance + car.speed
	}

	return Car {
		battery: remainingBattery,
		batteryDrain: car.batteryDrain,
		speed: car.speed,
		distance: remainingDistance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// panic("Please implement the CanFinish function")
	remainingDistance := track.distance - car.distance
	
	// If already at or past the finish line, can finish
	if remainingDistance <= 0 {
		return true
	}
	
	// Calculate how many drives are needed to cover the remaining distance
	// We need to round up: (remainingDistance + car.speed - 1) / car.speed
	drivesNeeded := (remainingDistance + car.speed - 1) / car.speed
	
	// Calculate battery needed
	batteryNeeded := drivesNeeded * car.batteryDrain
	
	// Check if car has enough battery
	return car.battery >= batteryNeeded
}
