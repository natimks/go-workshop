package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-workshop/src"
)

func main(){
	engine := gin.New();
	health := engine.Group("/");

	userRoute := engine.Group("/user")
	debtRoute := engine.Group("/debt")

	//user
	userRoute.GET("/", src.GetUsers)
	userRoute.GET("/:id", src.GetUser)
	userRoute.GET("/:id/debts", src.GetDebtsByUser)
	userRoute.POST("/", src.PostUser)
	userRoute.PUT("/:id", src.PutUser)
	userRoute.DELETE("/:id", src.DeleteUser)

	//user
	debtRoute.GET("/", src.GetDebts)
	debtRoute.GET("/:id", src.GetDebt)
	debtRoute.POST("/", src.PostDebt)
	debtRoute.PUT("/:id", src.PutDebt)
	debtRoute.DELETE("/:id", src.DeleteDebt)

	src.AutoMigration()


	health.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"status": "Go healthy!",
		})
	})

	engine.Run(fmt.Sprintf(":%v", 8088))
	//engine.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}