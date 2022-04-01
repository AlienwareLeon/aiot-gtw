package service

type Message struct {

}

func (message Message) TableName() string{
	//绑定数据库表
	return "message"
}

func validMessage() () {
	return
}

