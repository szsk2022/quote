package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"yiyan/conf"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client // 全局Redis客户端实例

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.C.Redis.Address,  // Redis服务器地址和端口
		Password: conf.C.Redis.Password, // 如果设置了密码，请填写
		DB:       conf.C.Redis.Database, // 数据库索引，默认是0
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("连接Redis失败: %v", err)
	}
}
func main() {
	fmt.Println("欢迎使用SZSK—QuoteAPI")
	fmt.Println("官方网址：https://www.sunzishaokao.com")
	fmt.Println("Ver:1.0.0")

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

	err := r.Run(conf.C.Web.Address)
	if err != nil {
		log.Fatalf("Run Serve err: " + err.Error())
		return
	}
}

func getRandomQuote(c *gin.Context) {
	lang := c.Query("lang") // 获取请求中的lang参数，默认为空字符串
	var content []byte
	var cacheKey string

	switch lang {
	case "en":
		cacheKey = "quotes_en"
	case "cn":
		fallthrough
	default:
		cacheKey = "quotes_cn"
	}

	content, err := rdb.Get(context.Background(), cacheKey).Bytes()
	if err == redis.Nil { // 缓存未命中
		var fileContent []byte
		switch lang {
		case "en":
			fileContent, err = ioutil.ReadFile("public/quotes_en.txt")
		case "cn":
			fallthrough
		default:
			fileContent, err = ioutil.ReadFile("public/quotes_cn.txt")
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "无法读取一言文件"})
			return
		}

		err = rdb.Set(context.Background(), cacheKey, fileContent, 0).Err() // 设置永不过期
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "缓存写入失败"})
			return
		}

		content = fileContent
	} else if err != nil { // 其他Redis错误
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "获取缓存时出错"})
		return
	}

	lines := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(lines))
	quote := lines[randomIndex]
	quote = strings.TrimSpace(quote)

	c.JSON(http.StatusOK, gin.H{"quote": quote})
}
