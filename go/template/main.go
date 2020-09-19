package main

import (
	"bytes"
	"fmt"
	"text/template"
	"time"
)

// 输出
/*
   	天津生态城生态环境监管温馨提示：【考核事件】施工扬尘于2020年09月02日16时已派发至APP，
     限时60分钟，请您及时确认并反馈执行情况，感谢您的配合
*/
// 结构体
type Message struct {
	ProjectName string
	EventType   string
	EventName   string
	Time        time.Time
	LimitMinute int
}

func (m Message) GetTimeStr() string {
	tf := "2006年01月02日15时"
	return m.Time.Format(tf)
}

func main() {
	m := Message{
		ProjectName: "天津生态环境监管",
		EventType:   "考核事件",
		EventName:   "施工扬尘",
		Time:        time.Now(),
		LimitMinute: 60,
	}
	a := getStrFromSprintf(m)
	b := getStrFromTemplate(m)
	fmt.Println(a)
	fmt.Println(b)

}

func getStrFromSprintf(m Message) string {
	return fmt.Sprintf("%s生态环境监管温馨提示：【%s】%s于%s已派发至APP，限时%d分钟，请您及时确认并反馈执行情况，感谢您的配合",
		m.ProjectName, m.EventType, m.EventName, m.GetTimeStr(), m.LimitMinute)
}

func getStrFromTemplate(m Message) string {
	t, err := template.New("messageTemplate").Parse("{{.ProjectName}}生态环境监管温馨提示：【{{.EventType}}】{{.EventName}}于{{.GetTimeStr}}已派发至APP，限时{{.LimitMinute}}分钟，请您及时确认并反馈执行情况，感谢您的配合")
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer([]byte{})
	err = t.Execute(buffer, m)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
