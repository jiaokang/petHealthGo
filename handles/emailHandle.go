package handles

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"petHealthTool/common"
	"time"

	"gopkg.in/gomail.v2"
)

type EmailHandle struct {
}

func (e *EmailHandle) SendVerifyCode(emailTo string, verifyCode string) error {
	// 加载数据库配置文件
	cfg, err := common.LoadConfig("E:/petHealthGo/config.yml")
	if err != nil {
		panic("faild to load config.yml file")
	}
	fmt.Println(cfg.Mail)
	// 配置邮件发送者信息
	from := cfg.Mail.User
	password := cfg.Mail.Pass
	to := emailTo
	smtpHost := cfg.Mail.Host
	smtpPort := cfg.Mail.Port

	// 读取 HTML 文件
	htmlContent, err := os.ReadFile("resource/verifycode_email_template.html")
	if err != nil {
		log.Fatal("Failed to read HTML file:", err)
		return err
	}

	// 定义替换的数据
	data := struct {
		VerificationCode string
		ExpirationTime   string
		CurrentYear      string
	}{
		VerificationCode: verifyCode,                // 替换为实际的验证码
		ExpirationTime:   "5",                       // 替换为实际的过期时间
		CurrentYear:      time.Now().Format("2006"), // 获取当前年份
	}

	// 解析 HTML 模板并替换占位符
	tmpl, err := template.New("email").Parse(string(htmlContent))
	if err != nil {
		log.Fatal("Failed to parse HTML template:", err)
		return err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		log.Fatal("Failed to execute template:", err)
		return err
	}

	// 创建邮件消息
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "验证码邮件")
	m.SetBody("text/html", body.String())

	// 创建 SMTP 客户端
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Failed to send email:", err)
		return err
	}
	log.Println("Email sent successfully!")
	return nil
}
