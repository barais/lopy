from mqttclient import MQTTClient
from network import WLAN,LoRa
import socket
import machine
import time

if __name__=='__main__':
    #WLAN
    wlan = WLAN(mode=WLAN.STA)
    wlan.connect("FD-51", auth=(WLAN.WPA2, "fromage2chevre"), timeout=5000)

    while not wlan.isconnected():
        machine.idle()
    print("Connected to Wifi")

    #MQTT
    mqttServerAddr = "192.168.43.253"
    client = MQTTClient("lopy", mqttServerAddr,port=1883)
    client.connect()
    print("Connected to MQTT at : {}".format(mqttServerAddr))

    #LoRa
    lora = LoRa(mode=LoRa.LORA, frequency=868100000)
    s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)
    s.setsockopt(socket.SOL_LORA,socket.SO_DR,5)
    s.setblocking(True)
    print("Listening to LoRaMAC")

    while True:
        print("waiting for data to send")
        loraClientData = int(s.recv(256))
        client.publish(topic="test/lopy", msg=str(loraClientData))
        print("data sent : {}".format(loraClientData))
        time.sleep(0.1)
