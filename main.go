package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 基础健康检查接口（启动后访问 http://localhost:8080/ping 验证环境是否正常）
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO: 实现社团技术部打卡接口
	// 接口规范要求：
	// 1. 路径：POST /api/v1/checkin
	// 2. 入参：JSON 格式，包含 name（必填）与 student_id（必填，长度 10~12 位）
	// 3. 响应：
	//    - 参数校验成功：返回 HTTP 200 OK，包含打卡成功提示与回显字段
	//    - 参数校验失败：返回 HTTP 400 Bad Request 与具体错误信息

	r.Run(":8080")
}
