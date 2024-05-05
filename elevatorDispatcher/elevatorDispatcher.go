package elevatorDispatcher

type Elevator interface {
	GetCurrentFloor() int
	GetDirection() string
	SetDirection(string)
	SetCurrentFloor(int)
	GetRequestQueue() []int
	SetRequestQueue(int)
	CreateRequestQueue(int)
	ClearRequestQueue(requestQueue int)
}
type ElevatorImpl struct {
	currentFloor int
	direction    string
	requestQueue []int
}

func NewElevator(currentFloor int, direction string) Elevator {
	elevator := new(ElevatorImpl)
	elevator.direction = direction
	elevator.currentFloor = currentFloor
	return elevator
}

func (elevatorImpl *ElevatorImpl) GetCurrentFloor() int {
	return elevatorImpl.currentFloor
}
func (elevatorImpl *ElevatorImpl) GetDirection() string {
	return elevatorImpl.direction
}
func (elevationImpl *ElevatorImpl) SetDirection(direction string) {
	elevationImpl.direction = direction
}
func (elevationImpl *ElevatorImpl) SetCurrentFloor(floor int) {
	elevationImpl.currentFloor = floor
}
func (elevationImpl *ElevatorImpl) GetRequestQueue() []int {
	return elevationImpl.requestQueue
}
func (elevationImpl *ElevatorImpl) SetRequestQueue(requestQueue int) {

	elevationImpl.requestQueue[requestQueue] += 1
}
func (elevationImpl *ElevatorImpl) CreateRequestQueue(requestQueue int) {
	elevationImpl.requestQueue = make([]int, requestQueue)
}
func (elevatorImpl *ElevatorImpl) ClearRequestQueue(requestQueue int) {
	elevatorImpl.requestQueue[requestQueue-1] = 0
}
