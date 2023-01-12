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
	//fmt.Println(result)
}
