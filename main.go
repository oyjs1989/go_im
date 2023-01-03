package main

import (
	_ "im_go/docs"
	"im_go/router"
	"im_go/utils"
)

// https://profilinator.rishav.dev/
func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
