# lopy


### Knowledge Base

Lopy devices come with an internal filesystem mounted on '/flash'. The file /flash/boot.py is executed at boot-up, then '/flash/main.py' is. When using the "Sync" feature from the Pymakr plugin, the filesystem is overwritten. However, the serial terminal is initiated in the file 'boot.py', so when erasing it, be aware that the serial communication will be disabled if done carelessely.

The initial 'boot.py' file is this one :

```python
# boot.py -- run on boot-up
import os
from machine import UART
uart = UART(0, 115200)
os.dupterm(uart)
```
## Useful commands
sudo chown root:wheel /dev/ttyUSB*
sudo chmod ou+rwx /dev/ttyUSB*
