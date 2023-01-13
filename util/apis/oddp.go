package apis

import (
	"github.com/gin-gonic/gin"
)

func Oddp(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	//log.Printf("%v", &json)
	//fmt.Println(json["param"])
	//fmt.Println(json["username"])
}
