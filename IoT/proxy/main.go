/*
 * Proxy which subscribes to a MQTT topic and exposes the data through a REST endpoint
 */

package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	payload := message.Payload()
	value, err := strconv.Atoi(string(payload))

	if err != nil {
		fmt.Printf("Error: received %s\n", payload)
	} else {
		fmt.Printf("Received %s\n", payload)
		data = append(data, value)
	}
}

var data []int
var i int64

func main() {
	data = make([]int, 0)
	c := make(chan os.Signal, 1)
	i = 0
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("signal received, exiting")
		os.Exit(0)
	}()

	connOpts := &MQTT.ClientOptions{
		ClientID:             "un-id",
		CleanSession:         true,
		Username:             "",
		Password:             "",
		MaxReconnectInterval: time.Duration(1 * time.Second),
		KeepAlive:            int64(30 * time.Second),
		TLSConfig:            tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert},
	}
	connOpts.AddBroker("tcp://192.168.43.253:1883")
	connOpts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe("test/lopy", 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", "localhost")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/test/lopy", func(w http.ResponseWriter, req *http.Request) {
		encoded, err := json.Marshal(data)
		if err == nil {
			fmt.Fprintf(w, string(encoded))
		}
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
