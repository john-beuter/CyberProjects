import std/net

let socket = newSocket()
socket.connect("google.com", Port(80))

