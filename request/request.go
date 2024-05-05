package request

import (
	"awesomeProject/directionEnum"
)

type Request interface {
	GetCurrentFloor() int
	GetDestFloor() int
	GetDirection() string
	GetStatus() string
	GetElevatorId() int
	SetElevatorId(int)
	SetStatus(string)
}
type RequestImpl struct {
	currentFloor int
	destFloor    int
	direction    string
	status       string
	elevatorId   int
}

func NewRequest(currentFloor, destFloor int) *RequestImpl {
	request := new(RequestImpl)
	request.currentFloor = currentFloor
	request.destFloor = destFloor
	request.status = directionEnum.WAITING
	request.elevatorId = -1

	if request.currentFloor < destFloor {
		request.direction = directionEnum.USER_UP
	} else if request.currentFloor > destFloor {
		request.direction = directionEnum.USER_DOWN
	} else {
		panic("You have choosen the same floor you are currently in")
	}
	return request
}

func (this *RequestImpl) GetCurrentFloor() int {
	return this.currentFloor
}
func (this *RequestImpl) GetDestFloor() int {
	return this.destFloor
}
func (this *RequestImpl) GetDirection() string {
	return this.direction
}
func (this *RequestImpl) GetStatus() string {
	return this.status
}
func (this *RequestImpl) GetElevatorId() int {
	return this.elevatorId
}
func (this *RequestImpl) SetElevatorId(id int) {
	this.elevatorId = id
}
func (this *RequestImpl) SetStatus(status string) {
	this.status = status
}
