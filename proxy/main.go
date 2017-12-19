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

type route struct {
	mqttHandler MQTT.MessageHandler
	httpHandler http.HandlerFunc
}

type sensorPing struct {
	Timestamp int64 `json:"timestamp"`
	Value     int   `json:"value"`
}

var myServerName = "*"
var lastValue *sensorPing
var routes = map[string]route{
	"test/lopy": route{testMQTTHandler, testHTTPHandler},
}

func addRoutes(client MQTT.Client, mux *http.ServeMux) {
	for topic, r := range routes {
		client.Subscribe(topic, 0, r.mqttHandler)
		mux.HandleFunc(fmt.Sprintf("/api/%s", topic), r.httpHandler)
	}
}

func testHTTPHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Received a pull")
	w.Header().Add("Access-Control-Allow-Origin", myServerName)
	data, _ := json.Marshal(lastValue)
	w.Write(data)
}

func testMQTTHandler(client MQTT.Client, message MQTT.Message) {
	payload := message.Payload()
	value, err := strconv.Atoi(string(payload))

	if err != nil {
		fmt.Printf("Error: received %s\n", payload)
	} else {
		fmt.Printf("Received %s\n", payload)
		lastValue = &sensorPing{time.Now().UnixNano(), value}
	}
}

func connectMqtt(user, passwd, server string) MQTT.Client {
	connOpts := &MQTT.ClientOptions{
		ClientID:             "lopy-project",
		CleanSession:         true,
		Username:             user,
		Password:             passwd,
		MaxReconnectInterval: time.Duration(1 * time.Second),
		KeepAlive:            int64(30 * time.Second),
		TLSConfig:            tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert},
	}

	connOpts.AddBroker(server)
	client := MQTT.NewClient(connOpts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func startServer(mux http.Handler, bind string) {
	server := &http.Server{
		Addr:           bind,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("signal received, exiting")
		os.Exit(0)
	}()

	client := connectMqtt("", "", "tcp://127.0.0.1:1883")
	mux := http.NewServeMux()

	addRoutes(client, mux)

	startServer(mux, ":8181")
}
