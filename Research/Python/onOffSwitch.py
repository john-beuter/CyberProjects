import time
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.12.151', port=502, auto_open=True, debug=False)

c.write_single_coil(3, 0)