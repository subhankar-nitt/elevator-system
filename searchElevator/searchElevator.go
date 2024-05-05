package searchElevator

import (
	"awesomeProject/directionEnum"
	"awesomeProject/elevatorService"
	"math"
)

type SearchElevator interface {
	FindElevator(int, string) int
}

type SearchElevatorImpl struct {
	elevatorService elevatorService.ElevatorService
}

func SearchElevatorService(elevatorService elevatorService.ElevatorService) SearchElevator {
	return &SearchElevatorImpl{elevatorService: elevatorService}
}

func (this *SearchElevatorImpl) FindElevator(srcFloor int, userDirection string) int {
	elevatorMap := this.elevatorService.GetElevators()
	minDistance := math.MaxInt
	minIndex := 1
	for k, v := range elevatorMap {
		currentFloor := v.GetCurrentFloor()
		elevatorDirection := v.GetDirection()
		distance := 0
		if userDirection == elevatorDirection || elevatorDirection == directionEnum.IDEAL {
			distance = abs(currentFloor, srcFloor)
		} else {
			requestQueue := v.GetRequestQueue()
			size := len(requestQueue)

			if elevatorDirection == directionEnum.UP {

				distance, _ = calculateUpPeek(requestQueue, currentFloor, size, srcFloor)
			} else {
				distance, _ = calculateDownPeek(requestQueue, currentFloor, size, srcFloor)
			}

		}
		if distance < minDistance {
			minDistance = distance
			minIndex = k
		}
	}
	return minIndex
}

func calculateUpPeek(queue []int, currElevatorFloor int, end int, srcFloor int) (int, int) {

	index := end - 1
	//for i := end - 2; i >= 0; i-- {
	//	queue[i] = queue[i+1] + queue[i]
	//}

	for i := end - 1; i >= 0; i-- {
		if queue[i] != 0 {
			break
		} else {
			index = i
		}
	}

	return (index - srcFloor) + (index - currElevatorFloor), index
}

func calculateDownPeek(queue []int, currElevatorFloor int, end int, srcFloor int) (int, int) {

	index := 0
	//for i := 1; i < end; i++ {
	//	queue[i] = queue[i-1] + queue[i]
	//}

	for i := 0; i < end; i++ {
		if queue[i] != 0 {
			break
		} else {
			index = i
		}
	}

	return (srcFloor - index) + (currElevatorFloor - index), index
}
func abs(a int, b int) int {
	sub := a - b
	if sub < 0 {
		sub = -sub
	}
	return sub
}
