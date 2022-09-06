package middles

import (
	"CodeSheep-runcode/configs"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetSessionId() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("GSESSIONID")
		if err != nil {
			sessionid := generateID()
			c.SetCookie("GSESSIONID", sessionid, 0, "/", configs.Domain, false, true)
		}
	}
}

func GetSessionId(c *gin.Context) string {
	id, err := c.Cookie("GSESSIONID")
	if err != nil {
		id = generateID()
		c.SetCookie("GSESSIONID", id, 0, "/", configs.Domain, false, true)
	}
	return id
}

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func generateID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	number := make([]byte, 32)
	if _, err := rand.Read(number); err != nil {
		log.Fatal("wrong can't generate sessionid: ", err.Error())
	}
	return Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(string(number)))
}
