import time
import sys
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.13.86', port=502, auto_open=True, debug=False)

if(sys.argv[2] == "1"):
    print(sys.argv[1])
    #print(type(int(sys.argv[2])))
    c.write_single_coil(int(sys.argv[1]), 1)
else:
    print(sys.argv[2])
    c.write_single_coil(int(sys.argv[1]), 0)
