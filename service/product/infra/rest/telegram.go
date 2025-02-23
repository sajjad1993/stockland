package rest

func registerTelegramRoutes(manager *ManagerHandler) {
	manager.group.telegram.POST("/request", manager.handler.ProductRequest)

}
