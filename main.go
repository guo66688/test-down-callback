package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()

	// 静态资源目录
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// 模拟发票页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "invoice.html", gin.H{
			"invoiceNo": "25952000000093739161",
		})
	})

	// 模拟下载接口
	r.GET("/download-pdf/:invoiceNo", func(c *gin.Context) {
		invoiceNo := c.Param("invoiceNo")
		filename := invoiceNo + "_发票.pdf"
		filepath := filepath.Join("static", filename)
		c.FileAttachment(filepath, filename)
	})

	r.Run(":8080")
}
