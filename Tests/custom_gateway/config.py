""" LoPy LoRaWAN Nano Gateway configuration options """

GATEWAY_ID = '1188227733664455' # specified in when registering your gateway

SERVER = '172.168.43.105' # server address & port to forward received data to
PORT = 1700

NTP = "pool.ntp.org" # NTP server for getting/setting time
NTP_PERIOD_S = 3600 # NTP server polling interval

WIFI_SSID = 'FD-51'
WIFI_PASS = 'fromage2chevre'

LORA_FREQUENCY = 868100000 # check your specifc region for LORA_FREQUENCY and LORA_DR (datarate)
LORA_DR = "SF7BW125"   # DR_5
