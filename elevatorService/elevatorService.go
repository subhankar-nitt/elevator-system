package elevatorService

import (
	"awesomeProject/directionEnum"
	"awesomeProject/elevatorDispatcher"
)

type ElevatorService interface {
	AddElevator(int)
	RemoveElevator(int)
	GetElevators() map[int]elevatorDispatcher.Elevator
}

type ElevatorServiceImpl struct {
	floor int
	hmap  map[int]elevatorDispatcher.Elevator
}

func NewElevatorService(floor int, elevatorCount int) ElevatorService {

	eleService := new(ElevatorServiceImpl)
	eleService.floor = floor
	hmap := make(map[int]elevatorDispatcher.Elevator)

	for i := 1; i <= elevatorCount; i++ {
		elevator := elevatorDispatcher.NewElevator(0, directionEnum.IDEAL)
		elevator.CreateRequestQueue(floor)
		hmap[i] = elevator
	}
	eleService.hmap = hmap

	return eleService
}
func (elevatorService *ElevatorServiceImpl) AddElevator(elevatorCount int) {
	for i := len(elevatorService.hmap); i <= len(elevatorService.hmap)+elevatorCount; i++ {
		elevator := elevatorDispatcher.NewElevator(0, directionEnum.IDEAL)
		elevator.CreateRequestQueue(elevatorService.floor)
		elevatorService.hmap[i] = elevator
	}
}

func (elevatorService *ElevatorServiceImpl) RemoveElevator(elevatorNo int) {
	delete(elevatorService.hmap, elevatorNo)
}
func (elevatorService *ElevatorServiceImpl) GetElevators() map[int]elevatorDispatcher.Elevator {
	return elevatorService.hmap
}
