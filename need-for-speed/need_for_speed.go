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
	newCar := Car{
		batteryDrain: batteryDrain,
		speed:        speed,
		battery:      100,
		distance:     0,
	}
	return newCar
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{
		distance: distance,
	}
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	var updatedCar Car
	if car.battery <= 0 || car.battery < car.batteryDrain {
		updatedCar = car
	} else {
		updatedCar = Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			distance:     car.speed + car.distance,
			battery:      car.battery - car.batteryDrain,
		}
	}
	return updatedCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	ticksToFinish := (track.distance - car.distance) / car.speed
	remainingBattery := car.battery - (ticksToFinish * car.batteryDrain)
	canFinish := remainingBattery >= 0
	return canFinish
}
