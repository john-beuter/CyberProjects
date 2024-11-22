import net

let client: Socket = newSocket()
client.connect("127.0.0.1", Port(5555))

    
let message: string = "Test"
client.send(message)
#client.send(message)

client.close()
