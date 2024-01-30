package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("欢迎使用SZSK—QuoteAPI")
	fmt.Println("官方网址：https://www.sunzishaokao.com")
	fmt.Println("Ver:1.0.0")

	// 根据环境变量设置 Gin 运行模式
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	// 如果环境变量未设置或无效，则默认设置为 Release 模式
	if gin.Mode() != gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true                                                               // 允许所有源（仅用于开发环境，请在生产环境中指定具体的允许域名）
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}                   // 允许的方法
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"} // 允许的头信息
	corsConfig.AllowCredentials = true                                                              // 允许携带凭据（如cookies）

	// 添加CORS中间件
	r.Use(cors.New(corsConfig))

	r.GET("/", getRandomQuote)
	r.POST("/", getRandomQuote)

	r.Static("/public", "./public")
	//r.StaticFile("/favicon.ico", "./favicon.ico") // cdn探测，负载均衡使用

	r.Run(":8080")
}

func getRandomQuote(c *gin.Context) {
	lang := c.Query("lang") // 获取请求中的lang参数，默认为空字符串
	var content []byte
	var err error

	switch lang {
	case "en":
		content, err = ioutil.ReadFile("public/quotes_en.txt") // 请确保此路径正确指向英文一言文件
	case "cn":
		fallthrough // 如果没有提供lang参数或者参数是"cn"，则默认读取中文一言文件
	default:
		content, err = ioutil.ReadFile("public/quotes_cn.txt")
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "无法读取一言文件"})
		return
	}

	lines := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(lines))
	quote := lines[randomIndex]
	quote = strings.TrimSpace(quote)

	c.JSON(http.StatusOK, gin.H{"quote": quote})
}
