import net
import sequtils

let
    ip = "192.168.13.86"
    port = Port(502)
    
var
    socket: Socket = newSocket()


socket.connect(ip, port)

var packet: array[12, uint8] = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x00, 0x00, 0xFF, 0x00]

let dataSent = packet.toSeq()

socket.send(dataSent)

socket.close()


