package main

import "awesomeProject/requestController"

func main() {
	requestCtrl := requestController.NewRequestController(6, 2)
	requestCtrl.CreateNewReq(1, 2)
	requestCtrl.CreateNewReq(2, 3)

	requestCtrl.CreateNewReq(4, 2)

	requestCtrl.CreateNewReq(6, 2)
	requestCtrl.CreateNewReq(4, 5)

	requestCtrl.ProcessRequest()

}
