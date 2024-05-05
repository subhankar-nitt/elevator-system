package requestController

import (
	"awesomeProject/elevatorService"
	"awesomeProject/processRequest"
	"awesomeProject/request"
	"awesomeProject/requestService"
)

type RequestController interface {
	ProcessRequest()
	CreateNewReq(int, int)
}

type RequestControllerImpl struct {
	processReq processRequest.ProcessRequest
}

func NewRequestController(floorNo int, elevatorCount int) RequestController {

	reqService := requestService.NewRequestService(floorNo)
	eleService := elevatorService.NewElevatorService(floorNo, elevatorCount)
	processReq := processRequest.ProcessNewRequest(reqService, eleService)

	return &RequestControllerImpl{processReq: processReq}

}

func (this *RequestControllerImpl) CreateNewReq(srcFloor int, destFloor int) {
	req := request.NewRequest(srcFloor, destFloor)
	this.processReq.AddNewReq(req)
}

func (this *RequestControllerImpl) ProcessRequest() {
	this.processReq.ProcessRequest()
}
