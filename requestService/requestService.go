package requestService

import (
	"awesomeProject/directionEnum"
	"awesomeProject/request"
	"fmt"
)

type RequestService interface {
	Add(request.Request)
	Remove(string, int, int)
	GetUpRequests() map[int][]request.Request
	GetDownRequests() map[int][]request.Request
}

type RequestServiceImpl struct {
	upRequestsSrc   map[int][]request.Request
	downRequestsSrc map[int][]request.Request
	upRequests      map[int][]request.Request
	downRequests    map[int][]request.Request
}

func NewRequestService(floorNo int) RequestService {
	requestService := new(RequestServiceImpl)
	requestService.upRequests = make(map[int][]request.Request)
	for i := 1; i <= floorNo; i++ {
		requestService.upRequests[i] = []request.Request{}
	}

	requestService.downRequests = make(map[int][]request.Request)
	for i := 1; i <= floorNo; i++ {
		requestService.downRequests[i] = []request.Request{}
	}
	requestService.upRequestsSrc = make(map[int][]request.Request)
	for i := 1; i <= floorNo; i++ {
		requestService.upRequestsSrc[i] = []request.Request{}
	}

	requestService.downRequestsSrc = make(map[int][]request.Request)
	for i := 1; i <= floorNo; i++ {
		requestService.downRequestsSrc[i] = []request.Request{}
	}
	return requestService
}

func (this *RequestServiceImpl) Add(request request.Request) {
	if request.GetStatus() == directionEnum.WAITING && request.GetDirection() == directionEnum.USER_DOWN {
		this.downRequestsSrc[request.GetCurrentFloor()] = append(this.downRequestsSrc[request.GetCurrentFloor()], request)
		this.downRequests[request.GetDestFloor()] = append(this.downRequests[request.GetDestFloor()], request)
	} else if request.GetStatus() == directionEnum.WAITING && request.GetDirection() == directionEnum.USER_UP {
		this.upRequests[request.GetDestFloor()] = append(this.upRequests[request.GetDestFloor()], request)
		this.upRequestsSrc[request.GetCurrentFloor()] = append(this.upRequestsSrc[request.GetCurrentFloor()], request)
	}
}

func (this *RequestServiceImpl) Remove(direction string, elevatorId int, destFloor int) {
	var hmap map[int][]request.Request

	if direction == directionEnum.UP {
		hmap = this.upRequests
	} else {
		hmap = this.downRequests
	}
	requests := hmap[destFloor]

	dupRequests := []request.Request{}
	for _, request := range requests {
		if request.GetElevatorId() != elevatorId && request.GetStatus() == directionEnum.WAITING {
			dupRequests = append(dupRequests, request)
		}
	}
	diff := len(requests) - len(dupRequests)
	if diff != 0 {

		fmt.Println(diff, " people have stepped out ")
	}
	hmap[elevatorId] = dupRequests

}

func (this *RequestServiceImpl) GetUpRequests() map[int][]request.Request {
	return this.upRequestsSrc
}
func (this *RequestServiceImpl) GetDownRequests() map[int][]request.Request {
	return this.downRequestsSrc
}
