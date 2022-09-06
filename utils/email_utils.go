package utils

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendBugEmail(sessionid string) {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "xx.hshi@qq.com"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{"mosquito@email.cn"}

	em.Subject = "Oops, Code Sheep Error Happen In Golang!"

	text := "Can't Stop the program: " + fmt.Sprint(sessionid)
	em.Text = []byte(text)

	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "xx.hshi@qq.com", "dbgbwpzmhtveebec", "smtp.qq.com"))

	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
