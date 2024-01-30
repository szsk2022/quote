package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", getRandomQuote)
	r.POST("/", getRandomQuote)

	r.Static("/public", "./public") // 假设quotes_txt文件在public文件夹下

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
