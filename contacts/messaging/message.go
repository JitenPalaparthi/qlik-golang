package messaging

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// func GetConnection(url string) (*nats.Conn, error) {
// 	nc, err := nats.Connect(url)
// 	return nc, err
// }

// type MessageWithOutConn struct {
// 	Conn *nats.Conn
// 	MSG  Message
// }

type IMessage interface {
	Publish() error
}

type Message struct {
	URL     string
	Subject string
	Data    []byte
	//Conn    *nats.Conn
}

func (m *Message) Publish() error {
	nc, err := nats.Connect(m.URL)
	defer nc.Close()
	if err != nil {
		return err
	}
	fmt.Println(m.Data)
	if nc != nil {
		if err := nc.Publish(m.Subject, m.Data); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("no connection to the message broker")
	}

}

func (m *Message) Subscribe() error {
	nc, err := nats.Connect(m.URL)
	fmt.Println(nc.ConnectedAddr())
	//defer nc.Close()
	if err != nil {
		return err
	}
	fmt.Println(m.Subject)
	//sub, err := nc.SubscribeSync(m.Subject)

	_, err = nc.Subscribe(m.Subject, func(m *nats.Msg) {

		//contact := new(models.Contact)
		//json.Unmarshal(m.Data, contact)
		fmt.Println("--->", string(m.Data))
		// todo stuff here
		//m.Respond([]byte("answer is 42"))
		//m.Sub.Unsubscribe()
	})

	return err
}

// 1- How do you creat connection to nats?
// 2- Do you try to connect to nats during traction?
