
from network import LoRa
import socket
import machine
import time
import pycom

lora = LoRa(mode=LoRa.LORA, frequency=868100000)
s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)
s.setblocking(True)
#pycom.rgbled(0xff00)

pycom.heartbeat(False)

adc = machine.ADC()             # create an ADC object
apin = adc.channel(pin='P13')   # create an analog pin on P16
val = apin()                    # read an analog value

while True:
    print('start')

    val = apin()
    if val > 900:
        pycom.rgbled(0x7f0000)
    else:
        pycom.rgbled(0x0000ff)  

    print(val)

    s.send(str(val))
    time.sleep(2.0)

    #s.send('coucou')
    #print(s.recv(64))
    #if s.recv(64) == b'Ping':
    #    s.send('Pong')
    #    print('receive ping -> send pong')
