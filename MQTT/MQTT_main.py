from mqtt import MQTTClient
from network import WLAN
import machine
import time


i = 0



def sub_cb(topic, msg):
   print(msg)

wlan = WLAN(mode=WLAN.STA)
wlan.connect("FD-51", auth=(WLAN.WPA2, "fromage2chevre"), timeout=5000)

while not wlan.isconnected():
    machine.idle()
print("Connected to Wifi\n")

client = MQTTClient("lopy", "192.168.43.253",port=1883)

client.set_callback(sub_cb)
client.connect()

while True:
    print("Sending value")
    client.publish(topic="test/lopy", msg=str(i))
    i+=1
    time.sleep(1)
