package model

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	encryptKey string = "ShkGStna8IV1iij"
	sessionID string = "GOSESSID"
	redisSessionPrefix string = "GOREDIS_SESSION::"
)
type VideoCompose struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Code        uint64 `json:"code" form:"code" binding:"required"`
	Timestamp   uint64 `json:"timestamp" form:"timestamp" binding:"required"`
	Sign        string `json:"sign" form:"sign" binding:"required"`
	encryptSign string
}

// 返回实例化对象
func NewVideoCompose(id,code,timestamp uint64,sign string) *VideoCompose  {
	return &VideoCompose{
		ID: id,
		Code: code,
		Timestamp: timestamp,
		Sign: sign,
	}
}

func (v *VideoCompose) Redirection(c *gin.Context) {
	v.createMd5()
	if v.verifySign() {
		c.SetCookie(sessionID, v.encryptSign[4:14] + v.encryptSign[11:21], 3600 * 24, "/", "127.0.0.1", false, true)
		c.Redirect(http.StatusMovedPermanently,"127.0.0.1")
	}
	//重定向到外部网页
	c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
}

// 检测sign是否合法
func (v *VideoCompose) verifySign() bool {
	if v.encryptSign == v.Sign {
		return true
	}
	return false
}

func (v *VideoCompose)  createMd5()  {
	composerStr := strconv.FormatUint(v.Timestamp,64) + encryptKey
	md5Byte := md5.Sum([]byte(composerStr))
	v.encryptSign = fmt.Sprintf("%x",md5Byte)
}

