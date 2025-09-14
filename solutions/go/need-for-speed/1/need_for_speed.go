package speed

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}

}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	speed := car.speed
	battery := car.battery
	batteryDrain := car.batteryDrain
	distance := car.distance

	if batteryDrain > battery {
		return car
	}

	distance += speed
	battery -= batteryDrain

	return Car{
		speed:        speed,
		battery:      battery,
		batteryDrain: batteryDrain,
		distance:     distance,
	}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {

	return car.battery/car.batteryDrain >= (track.distance-car.distance)/car.speed

}
