import std/net

let
    ip = "192.168.13.86"
    port = Port(502)
    

var
    socket: Socket = newSocket()

socket.connect(ip, port)

#Create packet to write coil
var packet: array[12, uint8] = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x00, 0x00, 0x00, 0x00];
var dataPtr: ptr array = packet.addr

var output: int 

output = socket.send(dataPtr, sizeof(packet))

echo output


#socket.sendTo(ip, port, dataPtr, sizeof(packet), AF_INET, 0'i32)
