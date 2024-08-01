package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type Config struct {
	Email      string `json:"email"`      //账号
	Name       string `json:"name"`       //发送者名字
	Password   string `json:"password"`   //邮箱授权码
	Mailserver string `json:"mailserver"` //邮件服务器
	Port       string `json:"port"`       //服务器端口
}

// 发送邮件信息
type Msg struct {
	Tmail   string
	Title   string
	Content string
}

func main() {

	// 加载配置文件，登录至邮箱
	config := LoadConfig("./config.json")
	fmt.Println(config)

	flag.Usage = flagUsage

	// msg := LoadConfig("./msg.json")
	to := flag.String("To", "845256503@qq.com", "要发送给谁")
	title := flag.String("title", "xxx", "主题")
	content := flag.String("content", "心情很好", "发送内容")

	flag.Parse()

	msg := &Msg{
		Tmail:   *to,
		Title:   *title,
		Content: *content,
	}
	fmt.Println(msg)
	if *to != "" && *title != "" && *content != "" {
		// SendMail(config, msg)

		// fmt.Printf("%v\n %v \n %v \n", *to, *title, *content)
	} else {
		panic("to,title,content can't be null!")
	}
}

// 加载配置文件
func LoadConfig(configPath string) (config *Config) {
	// 读取配置文件
	// data, err := ioutil.ReadFile(configPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := os.ReadFile(configPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(data)
	// 初始化用户信息
	config = &Config{}
	err = json.Unmarshal(data[:n], &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("打印config")
	fmt.Println(config)
	return config
}

func SendMail(config *Config, ms *Msg) {
	auth := smtp.PlainAuth("", config.Email, config.Password, config.Mailserver)

	to := []string{ms.Tmail} //接收用户
	user := config.Email
	nickname := config.Name

	subject := ms.Title
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := ms.Content
	msg := "To:" + strings.Join(to, ",") + "\r\nFrom: "
	msg += nickname + "<" + user + ">\r\nSubject: " + subject
	msg += "\r\n" + content_type + "\r\n\r\n" + body

	server := func(serverName, port string) string {

		var buffer bytes.Buffer

		buffer.WriteString(serverName)
		buffer.WriteString(":")
		buffer.WriteString(port)

		return buffer.String()

	}(config.Mailserver, config.Port)

	// 发送邮件
	err := smtp.SendMail(server, auth, user, to, []byte(msg))
	if err != nil {
		fmt.Printf("send mail error:%v\n", err)
	}

	// fmt.Println(server)
	fmt.Println(msg, "\n")
	// fmt.Printf("%v\n", auth)
}

func flagUsage() {

	usageText := `Usage mailTo [OPTION]
Usage parameter:

  -To      		 default: yourself
  -title         default: yourName
  -content       default: Hello`

	fmt.Fprintf(os.Stderr, "%s\n", usageText)
}
