import std/net

let socket = newSocket()
var packet: array[12, int] = [0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x06, 0x00, 0x04, 0x00, 0x01];
send(socket; "")

