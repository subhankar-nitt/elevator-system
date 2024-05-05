package building

import "fmt"

type Building interface {
	AddElevator(int)
	RemoveElevator(int)
	GetFloor() int
	GetElevator() int
}

type BuildingImpl struct {
	floorNo      int
	noOfElevator int
}

func CreateBuilding(floorNo int, noOfElevator int) Building {
	building := new(BuildingImpl)
	building.floorNo = floorNo
	building.noOfElevator = noOfElevator

	return building
}

func (b *BuildingImpl) AddElevator(no int) {

	b.noOfElevator += no
	fmt.Println("Added elevator", b.noOfElevator)
}
func (b *BuildingImpl) RemoveElevator(no int) {
	b.noOfElevator -= no
	fmt.Println("Removed elevator", b.noOfElevator)
}
func (b *BuildingImpl) GetFloor() int {
	return b.floorNo
}

func (b *BuildingImpl) GetElevator() int {
	return b.noOfElevator
}
