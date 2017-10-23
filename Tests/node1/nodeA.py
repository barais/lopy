from network import LoRa
import socket
import time

lora = LoRa(mode=LoRa.LORA, frequency=868100000)
s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)
s.setblocking(True)

while True:
    print('ok')
    s.send('coucou')
    #print(s.recv(64))
    #if s.recv(64) == b'Ping':
    #    s.send('Pong')
    #    print('receive ping -> send pong')
    time.sleep(0.1)
