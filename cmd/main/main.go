package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 📌 تنظیم فایل‌های استاتیک
	r.Static("/static", "./static")

	// 📌 مسیر اصلی برای سرو کردن `index.html`
	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	// 📌 API برای جستجو
	r.GET("/api/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API is working on Railway!"})
	})

	// 📌 دریافت پورت از `RAILWAY_PORT`
	port := os.Getenv("PORT") // Railway متغیر `PORT` رو تنظیم می‌کنه
	if port == "" {
		port = "3000" // برای تست لوکال
	}

	fmt.Println("🚀 Server is running on port railway", port)
	r.Run(":" + port)
}
