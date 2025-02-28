package main

import (
	"fmt"
	"net/http"
	"os"
	"stockland/config"
	usecase "stockland/service/product/app"
	"stockland/service/product/infra/mem"
	"stockland/service/product/infra/router"
	"stockland/service/product/infra/third_pary/openai"
)

func main() {
	// 📌 دریافت پورت از `Railway`
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Railway معمولاً از پورت 8080 استفاده می‌کند
	}
	config := config.NewConfigFromViper()
	repo := mem.NewMemRepository()
	mem.Seed(repo)
	apiKey := config.AIKey
	aiThirdParty := openai.NewOpenAIProcessor(apiKey)
	app := usecase.NewProductRequestUsecase(repo, aiThirdParty)
	handler := router.NewRouter(app)

	// 📌 اجرای سرور روی `0.0.0.0` برای دسترسی عمومی
	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handler,
	}

	fmt.Printf("🚀 سرور روی Railway با پورت %s اجرا شد\n", port)
	server.ListenAndServe()
}
