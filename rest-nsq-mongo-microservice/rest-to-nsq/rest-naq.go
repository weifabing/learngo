package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nsqio/go-nsq"
)

type Job struct {
	Title     string `json:"title"`
	scription string `json:"description"`
	Company   string `json:"company"`
	Salary    string `json:"salary"`
}

var (
	producer *nsq.Producer
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/jobs", jobsPostHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", router))
}

func jobsPostHandler(w http.ResponseWriter, r *http.Request) {
	//Retrieve body from http request
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	//Save data into Job struct
	var _job Job
	err = json.Unmarshal(b, &_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// 将数据发送到 MQ 队列
	go saveJobToNSQ(_job)

	//Convert job struct into json
	jsonString, err := json.Marshal(_job)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Set content-type http header
	w.Header().Set("content-type", "application/json")

	//Send back data as response
	w.Write(jsonString)
}

// 将提交的数据发布到NSQ系统
func saveJobToNSQ(job Job) error {
	fmt.Printf("save to nsq")

	jsonData, err := json.Marshal(job)

	if err != nil {
		panic(err)
	}

	// Produce messages to topic (asynchronously)
	//创建一个生产者，这里的RanDomGetServer()是自定义的一个工具，用来随机获取一个nsqd地址
	ProducerAddr := "192.168.56.100:4150"
	producer, err := nsq.NewProducer(ProducerAddr, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	topic := "jobs-topic1"
	err = producer.Publish(topic, jsonData)
	if err != nil {
		return fmt.Errorf("new nsq producer client is err: %v", err)
	}

	return nil
}
