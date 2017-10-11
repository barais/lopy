from network import LoRa
import socket
import time

lora = LoRa(mode=LoRa.LORA, frequency=863000000)
s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)
s.setblocking(False)

while True:
    print('ok')
    if s.recv(64) == b'Ping':
        s.send('Pong')
        print('receive ping -> send pong')
    time.sleep(5)
