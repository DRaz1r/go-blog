/*
*
@author: Azir
@desc:
@date: 6/19/24
*
*/
package main

import (
	"backend/common"
	"backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := common.InitDB()
	defer DB.Close()
	r := gin.Default()
	r.StaticFS("/images", http.Dir("./static/images/"))
	routes.CollectRoutes(r)
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run(":8080")
}
