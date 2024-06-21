import time
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.13.86', port=502, auto_open=True, debug=False)

c.write_single_coil(0, 0)