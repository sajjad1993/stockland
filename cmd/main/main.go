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
	{"name": "گوشی هوشمند", "description": "الکترونیک"},
	{"name": "لپ‌تاپ", "description": "کامپیوتر"},
	{"name": "هدفون", "description": "لوازم جانبی"},
	{"name": "ساعت هوشمند", "description": "گجت"},
	{"name": "کوله‌پشتی", "description": "مد و پوشاک"},
}

func main() {
	// 📌 فعال کردن `Logger` برای نمایش لاگ درخواست‌ها
	r := gin.Default()
	r.Use(gin.Logger())

	// 📌 اطمینان از دسترسی به درخواست‌های `Proxy`
	r.SetTrustedProxies(nil)

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
		//40 mini
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
		c.File("./static/index.html")
	})

	// 📌 دریافت پورت از `Railway`
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Railway معمولاً از پورت 8080 استفاده می‌کند
	}

	// 📌 اجرای سرور روی `0.0.0.0` برای دسترسی عمومی
	fmt.Printf("🚀 سرور روی Railway با پورت %s اجرا شد\n", port)
	r.Run("0.0.0.0:" + port)
}
