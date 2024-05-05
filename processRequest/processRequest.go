package processRequest

import (
	"awesomeProject/directionEnum"
	"awesomeProject/elevatorService"
	"awesomeProject/request"
	"awesomeProject/requestService"
	"awesomeProject/searchElevator"
	"fmt"
	"math"
	"sync"
)

type ProcessRequest interface {
	ProcessRequest()
	AddNewReq(request.Request)
}

type ProcessRequestImpl struct {
	requestService  requestService.RequestService
	elevatorService elevatorService.ElevatorService
}

func ProcessNewRequest(requestService requestService.RequestService, elevatorService elevatorService.ElevatorService) *ProcessRequestImpl {

	return &ProcessRequestImpl{requestService, elevatorService}
}

func (this *ProcessRequestImpl) ProcessRequest() {
	upReqs := this.requestService.GetUpRequests()
	downReqs := this.requestService.GetDownRequests()

	var wg sync.WaitGroup
	wg.Add(2)
	go this.processUpRequest(upReqs, &wg)
	go this.processDownRequest(downReqs, &wg)
	wg.Wait()

}

func (this *ProcessRequestImpl) processUpRequest(requests map[int][]request.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	lowestReqFloor := math.MaxInt
	elevatorId := -1
	for k := 1; k <= len(requests); k++ {
		v := requests[k]
		if len(v) != 0 && k < lowestReqFloor {
			search_elevator := searchElevator.SearchElevatorService(this.elevatorService)
			elevatorId = search_elevator.FindElevator(k, directionEnum.USER_UP)
			lowestReqFloor = k
		}
		if elevatorId != -1 {

			fmt.Println("elevator ", elevatorId, "comming to ", k, " floor.")
			elevator := this.elevatorService.GetElevators()[elevatorId]
			elevator.SetDirection(directionEnum.UP)
			elevator.SetCurrentFloor(k)

			this.requestService.Remove(directionEnum.UP, elevatorId, k)
			elevator.ClearRequestQueue(k)
			if len(v) != 0 {

				for _, req := range v {
					req.SetStatus(directionEnum.PROCESSING)
					elevator.SetRequestQueue(req.GetDestFloor() - 1)

				}
				fmt.Println("Passengers On boarded")
			} else {
				sum := 0
				lastReq := true
				reqQueue := elevator.GetRequestQueue()
				for i := k; i < len(reqQueue); i++ {
					if len(requests[i]) != 0 {
						lastReq = false
					}
					sum += reqQueue[i]
				}
				if sum == 0 && lastReq {
					elevator.SetDirection(directionEnum.IDEAL)
					fmt.Println("No more passengers , going to ideal at ", k, " floor.")
					break
				}

			}
		}

	}

}

func (this *ProcessRequestImpl) processDownRequest(requests map[int][]request.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	heighestReqFloor := math.MinInt
	elevatorId := -1
	for k := len(requests); k > 0; k-- {
		v := requests[k]
		if len(v) != 0 && k > heighestReqFloor {
			search_elevator := searchElevator.SearchElevatorService(this.elevatorService)
			elevatorId = search_elevator.FindElevator(k, directionEnum.USER_DOWN)
			heighestReqFloor = k
		}
		if elevatorId != -1 {

			fmt.Println("elevator ", elevatorId, "comming to ", k, " floor.")
			elevator := this.elevatorService.GetElevators()[elevatorId]
			elevator.SetDirection(directionEnum.DOWN)
			elevator.SetCurrentFloor(k)

			this.requestService.Remove(directionEnum.DOWN, elevatorId, k)
			elevator.ClearRequestQueue(k)
			if len(v) != 0 {

				for _, req := range v {
					req.SetStatus(directionEnum.PROCESSING)
					elevator.SetRequestQueue(req.GetDestFloor() - 1)

				}
				fmt.Println("Passengers On boarded")
			} else {
				sum := 0
				lastReq := true
				reqQueue := elevator.GetRequestQueue()
				for i := k; i > 0; i-- {
					if len(requests[i]) != 0 {
						lastReq = false
					}
					sum += reqQueue[i-1]
				}
				if sum == 0 && lastReq {
					elevator.SetDirection(directionEnum.IDEAL)
					fmt.Println("No more passengers , going to ideal at ", k, " floor.")
					break
				}

			}
		}

	}
}

func (this *ProcessRequestImpl) AddNewReq(req request.Request) {
	this.requestService.Add(req)
}
