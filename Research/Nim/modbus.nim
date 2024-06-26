import std/net

#Packet that I am trying to send
#==0x00, 0x00== | ==0x00, 0x00== | 0x00, 0x06 | 0x01 | 0x05 | 0x00 0x00 (write reg 0) 0xFF 0x00 [Write a one to the first register]


#let socket = newSocket(AF_INET, SOCK_STREAM, IPPROTO_TCP)

let socket = newSocket(AF_INET, SOCK_STREAM, IPPROTO_TCP)
socket.connect("192.168.13.86", Port(502))

var packet: array[12, int] = [0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x00, 0x00, 0xFF, 0x00];

var dataPtr: ptr array = packet.addr


socket.sendTo("192.168.13.86", Port(502), dataPtr, sizeof(packet), AF_INET, 0'i32)


