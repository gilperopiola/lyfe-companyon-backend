package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterActions interface {
	Setup()
}

type MyRouter struct {
	*gin.Engine
}

func (router *MyRouter) Setup() {
	gin.SetMode(gin.DebugMode)
	router.Engine = gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	public := router.Group("/")
	{
		public.POST("/Signup", SignUp)
		public.POST("/Login", Login)
	}

	users := router.Group("/Users", validateToken())
	{
		users.POST("", CreateUser)
		users.GET("", SearchUsers)
		users.GET("/:id_user", GetUser)
		users.PUT("/:id_user", UpdateUser)
	}

	tags := router.Group("/Tags", validateToken())
	{
		tags.POST("", CreateTag)
		tags.GET("", SearchTags)
		tags.GET("/:id_tag", GetTag)
		tags.PUT("/:id_tag", UpdateTag)

		tags.GET("/:id_tag/Tasks", GetTagTasks)
	}

	tasks := router.Group("/Tasks", validateToken())
	{
		tasks.POST("", CreateTask)
		tasks.GET("", SearchTasks)
		tasks.GET("/:id_task", GetTask)
		tasks.PUT("/:id_task", UpdateTask)
		tasks.PUT("/:id_task/Doing", SetTaskDoing)
		tasks.PUT("/:id_task/Complete", CompleteTask)
		tasks.PUT("/:id_task/Archive", ArchiveTask)
	}

	history := router.Group("/History", validateToken())
	{
		history.GET("/Week", GetWeeklyDoneAndArchivedTasks)
		history.GET("/Month", GetMonthlyDoneAndArchivedTasks)
	}
}
