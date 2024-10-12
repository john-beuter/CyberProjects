import net

let server: Socket = newSocket()
server.bindAddr(Port(5555))
server.listen()
stdout.writeLine("Server: started. Listening to new connections on port 5555...")

var client: Socket = new(Socket)
server.accept(client)
stdout.writeLine("Server: client connected")

while true:
  let message: string = client.recv(9)

  if message == "":
    break

  stdout.writeLine("Server: received from client: ", message)

server.close()