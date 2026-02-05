package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	result := Car{
		batteryDrain: car.batteryDrain,
		speed:        car.speed,
	}
	if car.battery >= car.batteryDrain {
		result.battery = car.battery - car.batteryDrain
		result.distance += car.distance + car.speed
	} else {
		result.battery = car.battery
		result.distance = car.distance
	}
	return result
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if int(car.battery/car.batteryDrain)*car.speed >= track.distance {
		return true
	}
	return false
}
