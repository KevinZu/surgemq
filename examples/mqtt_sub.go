package main

import (
	"fmt"
	//"github.com/surge/netx"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	//"log"
	"os"
	"time"
)

var done chan struct{}

func main() {

	done = make(chan struct{})

	// Instantiates a new Client
	c := &service.Client{}

	// Creates a new MQTT CONNECT message and sets the proper parameters
	//msg := message.NewConnectMessage()
	//msg.SetWillQos(2)
	//msg.SetVersion(4)
	//msg.SetCleanSession(true)
	//msg.SetClientId([]byte("surgemq"))
	//msg.SetKeepAlive(10)
	//msg.SetWillTopic([]byte("will"))
	//msg.SetWillMessage([]byte("send me home"))
	//msg.SetUsername([]byte("surgemq"))
	//msg.SetPassword([]byte("verysecret"))

	msg := message.NewConnectMessage()
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte(fmt.Sprintf("pingmqclient%d%d", os.Getpid(), time.Now().Unix())))
	msg.SetKeepAlive(300)

	// Connects to the remote server at 127.0.0.1 port 1883
	c.Connect("tcp://127.0.0.1:1883", msg)

	// Creates a new SUBSCRIBE message to subscribe to topic "abc"
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("abc"), 0)

	subhello := message.NewSubscribeMessage()
	subhello.AddTopic([]byte("helloworld"), 0)

	// Subscribes to the topic by sending the message. The first nil in the function
	// call is a OnCompleteFunc that should handle the SUBACK message from the server.
	// Nil means we are ignoring the SUBACK messages. The second nil should be a
	// OnPublishFunc that handles any messages send to the client because of this
	// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
	c.Subscribe(submsg, nil, onPublish)
	c.Subscribe(subhello, nil, onHelloPublish)

	// Disconnects from the server
	//c.Disconnect()
	<-done
	c.Disconnect()
}

func onPublish(msg *message.PublishMessage) error {
	//pr := &netx.PingResult{}
	//if err := pr.GobDecode(msg.Payload()); err != nil {
	//	log.Printf("Error decoding ping result: %v\n", err)
	//	return err
	//}
	fmt.Printf("%s\n", msg.Payload())
	//log.Println(pr)
	return nil
}

func onHelloPublish(msg *message.PublishMessage) error {
	//pr := &netx.PingResult{}
	//if err := pr.GobDecode(msg.Payload()); err != nil {
	//	log.Printf("Error decoding ping result: %v\n", err)
	//	return err
	//}
	fmt.Printf("%s\n", msg.Payload())
	//log.Println(pr)
	return nil
}
