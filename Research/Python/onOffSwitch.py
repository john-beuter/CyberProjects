import time
import sys
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.13.86', port=502, auto_open=True, debug=False)

if(sys.argv[1] == "1"):
    print(sys.argv[1])
    c.write_single_coil(0, 1)
else:
    print(sys.argv[1])
    c.write_single_coil(0, 0)
