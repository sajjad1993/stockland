package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// 📌 داده‌های تستی
var products = []gin.H{
	{"name": "گوشی هوشمند", "category": "الکترونیک"},
	{"name": "لپ‌تاپ", "category": "کامپیوتر"},
	{"name": "هدفون", "category": "لوازم جانبی"},
	{"name": "ساعت هوشمند", "category": "گجت"},
	{"name": "کوله‌پشتی", "category": "مد و پوشاک"},
}

func main() {
	r := gin.Default()

	// 📌 سرو کردن فایل‌های استاتیک (index.html، CSS، JS)
	r.Static("/static", "./static")

	// 📌 API برای دریافت کل لیست یا جستجو
	r.GET("/api/search", func(c *gin.Context) {
		query := c.Query("q")

		// اگر جستجو خالی بود، کل لیست رو برگردون
		if query == "" {
			c.JSON(http.StatusOK, products)
			return
		}

		// فیلتر کردن نتایج مطابق با عبارت جستجو
		var results []gin.H
		for _, product := range products {
			if strings.Contains(strings.ToLower(product["name"].(string)), strings.ToLower(query)) {
				results = append(results, product)
			}
		}

		c.JSON(http.StatusOK, results)
	})

	// 📌 سرو کردن `index.html` به عنوان صفحه اصلی
	r.GET("/", func(c *gin.Context) {
		c.File("static/index.html")
	})

	port := os.Getenv("PORT") // Railway متغیر `PORT` رو تنظیم می‌کنه
	if port == "" {
		port = "3000" // برای تست لوکال
	}
	// 📌 اجرای سرور روی پورت 3000
	fmt.Printf("🚀 سرور روی  railway %s اجرا شد\n", port)
	r.Run(fmt.Sprintf(":%s", port))
}
