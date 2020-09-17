package api

import (
	"github.com/gin-gonic/gin"
	"go_yangzhou/model"
	"net/http"
)

func VideoCompose(c *gin.Context)  {
	var video model.VideoCompose
	if err := c.ShouldBind(&video); err == nil{
		video := model.NewVideoCompose(video.ID,video.Code,video.Timestamp,video.Sign)
		video.Redirection(c)
		c.JSON(http.StatusOK,
			model.NewVideoCompose(video.ID,video.Code,video.Timestamp,video.Sign),
			)
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// video合成回调
func VideoCallBack(c *gin.Context)  {

}