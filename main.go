package main

import "sign/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run(":6789")

}
