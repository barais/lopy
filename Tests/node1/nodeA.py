from network import LoRa
import socket
import time

lora = LoRa(mode=LoRa.LORAWAN, frequency=868100000)

s = socket.socket(socket.AF_LORA, socket.SOCK_RAW)

s.setsockopt(socket.SOL_LORA,socket.SO_DR,5)

s.setblocking(True)

while True:
    print("waiting for data...")
    print(s.recv(256))
    time.sleep(0.1)
