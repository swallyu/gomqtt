package broker

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang/packets"
	"log"
	"net"
)

type BrokerConfig struct {
	port uint
	host string
}

type MqBroker struct {
}

func NewBroker() *MqBroker {
	return &MqBroker{}
}

func (*MqBroker) ListenAndServe() error {

	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		// handle error
		return nil
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}

	return nil
}

func handleConnection(conn net.Conn) {

	packet, err := packets.ReadPacket(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(packet.String())

	msg, ok := packet.(*packets.ConnectPacket)

	if !ok {
		log.Println("received msg that was not Connect")
		return
	}

	log.Println(msg.Dup)

	connack := packets.NewControlPacket(packets.Connack).(*packets.ConnackPacket)
	connack.SessionPresent = msg.CleanSession
	connack.ReturnCode = msg.Validate()

	if connack.ReturnCode != packets.Accepted {
		err = connack.Write(conn)
		if err != nil {
			log.Fatal("send connack error, ", msg.ClientIdentifier)
			return
		}
		return
	}
	err = connack.Write(conn)
	if err != nil {
		log.Fatal("send connack error, ", msg.ClientIdentifier)
		return
	}
}
