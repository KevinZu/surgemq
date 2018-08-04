package main

import (
	"fmt"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	"time"
)

func main() {

	// Instantiates a new Client
	c := &service.Client{}
	//cnt := 0
	//tick := time.NewTicker(10 * time.Second)

	// Creates a new MQTT CONNECT message and sets the proper parameters

	msg := message.NewConnectMessage()
	msg.SetWillQos(2)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(300)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	//	msg.SetUsername([]byte("surgemq"))
	//	msg.SetPassword([]byte("verysecret"))

	// Connects to the remote server at 127.0.0.1 port 1883
	c.Connect("tcp://127.0.0.1:1883", msg)
	/*
		// Creates a new SUBSCRIBE message to subscribe to topic "abc"
		submsg := message.NewSubscribeMessage()
		submsg.AddTopic([]byte("abc"), 0)

		// Subscribes to the topic by sending the message. The first nil in the function
		// call is a OnCompleteFunc that should handle the SUBACK message from the server.
		// Nil means we are ignoring the SUBACK messages. The second nil should be a
		// OnPublishFunc that handles any messages send to the client because of this
		// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
		c.Subscribe(submsg, nil, nil)
	*/
	// Creates a new PUBLISH message with the appropriate contents for publishing

	for {
		//if cnt != 0 {
		//	<-tick.C
		//}

		pubmsg := message.NewPublishMessage()
		//pubmsg.SetPacketId(2)
		pubmsg.SetTopic([]byte("abc"))

		pubmsg.SetQoS(0)

		pubmsg.SetPayload( /*make([]byte, 1024)*/ []byte("hello world!"))
		// Publishes to the server by sending the message
		c.Publish(pubmsg, onComplate)

		pubmsg1 := message.NewPublishMessage()
		//pubmsg.SetPacketId(2)
		pubmsg1.SetTopic([]byte("helloworld"))

		pubmsg1.SetQoS(0)

		pubmsg1.SetPayload( /*make([]byte, 1024)*/ []byte("hehehe!"))
		// Publishes to the server by sending the message
		c.Publish(pubmsg1, nil)

		time.Sleep(time.Second * 5)

	}
	// Disconnects from the server

	c.Disconnect()

}

func onComplate(msg, ack message.Message, err error) error {
	//pr := &netx.PingResult{}
	//if err := pr.GobDecode(msg.Payload()); err != nil {
	//	log.Printf("Error decoding ping result: %v\n", err)
	//	return err
	//}
	fmt.Println("complate")
	//log.Println(pr)
	return nil
}
