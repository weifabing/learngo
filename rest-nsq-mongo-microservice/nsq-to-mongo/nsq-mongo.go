package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"gopkg.in/mgo.v2"
)

const(
	NSQHost  = "192.168.56.100:4161"
	NSQTopic = "jobs-topic1"
	hosts      = "192.168.56.100:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "jobs"
)

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Salary      string `json:"salary"`
}

type MongoStore struct {
	session *mgo.Session
}

// 消费者
type ConsumerT struct{
	DataServerAddrs  map[string]time.Time  //保存dataServer进程发过来的服务器地址和接收时间
	rwLocker         sync.RWMutex  //防止多线程同时读写
}

var mongoStore = MongoStore{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	saveJobToMongoDB(msg.Body)
	return nil
}


// 保存数据到Mongodb
func saveJobToMongoDB(bytes []byte) {
	fmt.Println("Save to MongoDB")

	col := mongoStore.session.DB(database).C(collection)

	//Save data into Job struct
	var _job Job
	err := json.Unmarshal(bytes, &_job)
	if err != nil {
		panic(err)
	}

	//Insert job into MongoDB
	errMongo := col.Insert(_job)
	if errMongo != nil {
		panic(errMongo)
	}

	fmt.Printf("Saved to MongoDB : %s\n", bytes)
}



func main()  {
	// Create MongoDB session
	session := initialiseMongo()
	mongoStore.session = session
	
	fmt.Printf("订阅NSQ消息开始")
	InitNsqConsumer(NSQTopic, "C1", NSQHost)

	for {
		time.Sleep(time.Second * 10)
	}
}


func initialiseMongo() (session *mgo.Session) {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	return session
}

func InitNsqConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}

	c.SetLogger(nil, 0)  //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 添加消费者接口

	//建立 NSQLookupd 连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}
}
