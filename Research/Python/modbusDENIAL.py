import time
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.13.86', port=502, auto_open=True, debug=False)
bit = False





# main loop
while True:
    is_ok = c.write_single_coil(1, 1)
    if is_ok:
        print('coil #%s: write to %s' % (1, 1))
    else:
        print('coil #%s: unable to write %s' % (1, 1))

    

