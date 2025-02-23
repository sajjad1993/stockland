package rest

import "stockland/service/product/infra/rest/middleware"

func registerAdminRoutes(manager *ManagerHandler) {
	manager.group.users.Use(middleware.AdminRequired())
	{
		// manager.group.users.GET("", manager.handler.GetAllUsers)
		// manager.group.users.POST("", manager.handler.CreateUser)
		// manager.group.users.PUT("/:id", manager.handler.UpdateUser)
		// manager.group.users.DELETE("/:id", manager.handler.DeleteUser)
	}

	manager.group.servers.Use(middleware.AdminRequired())
	{
		// manager.group.servers.GET("", manager.handler.GetAllServers)
		// manager.group.servers.POST("", manager.handler.CreateServer)
		// manager.group.servers.PATCH("/:id", manager.handler.UpdateServer)
		// manager.group.servers.DELETE("/:id", manager.handler.DeleteServer)
	}

	manager.group.plans.Use(middleware.AdminRequired())
	{
		// manager.group.plans.GET("", manager.handler.GetAllPlans)
		// manager.group.plans.POST("", manager.handler.CreatePlan)
		// manager.group.plans.DELETE("/:id", manager.handler.DeletePlan)
		// manager.group.plans.PATCH("/:id", manager.handler.UpdatePlans)
	}

	manager.group.adminResellers.Use(middleware.AdminRequired())
	{
		// manager.group.adminResellers.POST("/associate/:id", manager.handler.AssociatePlans)
		// manager.group.adminResellers.GET("/:id", manager.handler.GetReseller)
		// manager.group.adminResellers.GET("", manager.handler.GetAllResellers)
		// manager.group.adminResellers.POST("", manager.handler.CreateReseller)
		// manager.group.adminResellers.PUT("/:id", manager.handler.UpdateReseller)
		// manager.group.adminResellers.DELETE("/:id", manager.handler.DeleteReseller)
	}

	manager.group.transactions.Use(middleware.AdminRequired())
	{
		// manager.group.transactions.GET("/:id", manager.handler.GetTransaction)
		// manager.group.transactions.POST("", manager.handler.CreateTransaction)
		// manager.group.transactions.DELETE("/:id", manager.handler.DeleteTransaction)
	}

	manager.group.adminAccounts.Use(middleware.AdminRequired())
	{
		// manager.group.adminAccounts.GET("/deleted", manager.handler.GetAllDeletedAccounts(manager.config))
		// manager.group.adminAccounts.GET("/warn", manager.handler.AdminExpireWarn)
		// manager.group.adminAccounts.GET("", manager.handler.GetAllAccounts(manager.config))
		// manager.group.adminAccounts.POST("/clear", manager.handler.ClearAccount)
	}

	manager.group.logs.Use(middleware.AdminRequired())
	{
		// manager.group.logs.GET("", manager.handler.GetLogs)
	}
}
