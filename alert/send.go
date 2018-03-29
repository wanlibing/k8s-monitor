package alert

import (
	"github.com/go-gomail/gomail"
	"fmt"
	"net/http"
	"bytes"
	"strconv"
)


var m *gomail.Message = gomail.NewMessage()

func SendMail(alter map[string]int){
	//link https://blog.csdn.net/qq_30949367/article/details/71076193
	message := ""
	for k,v := range alter{
		message = message + "Warning:" + k + "total restart times is" + strconv.Itoa(v) + "\n"
	}
	m.SetAddressHeader("From","xx@qq.com","xx@qq.com")
	m.SetHeader("To",m.FormatAddress("xx@qq.com","wng"))
	m.SetHeader("Cc",m.FormatAddress("xx@qq.com","litn"))
	m.SetHeader("Cc",m.FormatAddress("zxx@qq.com","zhya"))//抄送
	m.SetHeader("Subject","k8s-monitor-podrestartcount")
	m.SetBody("",message)  //how
	d := gomail.NewDialer("smtp.exmail.qq.com",465,"xx@qq.com","xx@qq.com")
	if err := d.DialAndSend(m);err != nil{
		fmt.Println("email send failed",err)
		return
	}
	fmt.Println("send mail success")
}


func SendDingDing(alter map[string]int){
	message := ""
	for k,v := range alter{
		message = message + "Warning:" + k + "total restart times is" + strconv.Itoa(v) + "\n"
	}
	dingdingUrl := "https://oapi.dingtalk.com/robot/send?access_token=bb7a5e1be3d46ff62f5cb8548c81248d5a6677177901f76823b3"
	formt := `
        {
            "msgtype": "markdown",
            "markdown": {
                "title":"重构环境集群监控",
                "text": "%s"
            }
        }`
	body := fmt.Sprintf(formt,message)
	jsonvalue := []byte(body)
	resp,err := http.Post(dingdingUrl,"application/json",bytes.NewBuffer(jsonvalue))
	if err != nil{
		fmt.Println("send failed")
	}
	fmt.Println("send successed")
	fmt.Println(resp)
}

