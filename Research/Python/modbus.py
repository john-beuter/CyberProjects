from pymodbus.client.sync import ModbusTcpClient

# Modbus TCP server address and port
server_ip = "192.168.14.236" 
server_port = 502 

# Modbus addresses of the coils to write to
coil_addresses = [0] 

# Connect to the Modbus server
client = ModbusTcpClient(server_ip, port=server_port)
client.connect()

try:
    while True:
        # Manual Override true, Dropboomgate false
        new_coil_values = [False] 
       
        for coil_address, new_coil_value in zip(coil_addresses, new_coil_values):
            response = client.write_coil(coil_address, new_coil_value, unit=1)  # Unit ID is typically 1

            if response.isError():
                print(f"Modbus Write Coil Error: {response.getExceptionDescription()}")
            else:
                print(f"Coil {coil_address} set to {new_coil_value}")

except Exception as e:
    print(f"Error: {e}")

finally:
    # Disconnect from the Modbus server
    client.close()
