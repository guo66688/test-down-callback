package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// 页面：展示按钮
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"InvoiceNo": "TEST123", // 可用作占位符
		})
	})

	// 提供一个固定 PDF 文件供下载
	r.GET("/download-pdf/:invoiceNo", func(c *gin.Context) {
		// 实际只返回固定文件
		filepath := filepath.Join("static", "dzfp_25952000000093739161_特灵科技（中国）有限公司_20250514154937.pdf")
		c.FileAttachment(filepath, "dzfp_25952000000093739161_特灵科技（中国）有限公司_20250514154937.pdf")
	})

	r.Run(":8080")
}
