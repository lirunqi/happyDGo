package apis

import (
	"github.com/gin-gonic/gin"
	"happyD/util/model"
	"net/http"
)

// 列表数据
func Odds(c *gin.Context) {
	var odd model.Odd

	result, err := odd.Odds()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
	//fmt.Println(c.)
}

// 查询列表数据
func OddF(c *gin.Context) {
	json := make(map[string]interface{})
	c.BindJSON(&json)
	var odd model.Odd
	result, err := odd.OddF(json["param"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}
