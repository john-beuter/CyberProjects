import time
from pyModbusTCP.client import ModbusClient

# init
c = ModbusClient(host='192.168.16.164', port=502, auto_open=True, debug=False)
bit = False





# main loop
while True: ## The first value is the coil and the second is the value
    is_ok = c.write_single_coil(1, 1)
    #c.write_single_coil(0,1)
    #c.write_single_coil(1,1)

    if is_ok:
        print('coil #%s: write to %s' % (2, 0))
    else:
        print('coil #%s: unable to write %s' % (2, 0))

    

