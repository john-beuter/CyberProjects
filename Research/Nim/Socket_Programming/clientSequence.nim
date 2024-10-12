import net
import sequtils

let client: Socket = newSocket()
client.connect("127.0.0.1", Port(5555))

while true: 
    var packet = @["0x00", "0x00", "0x00", "0x00", "0x00", "0x06", "0x01", "0x06", "0x00", "0x04", "0x00", "0x01"]
    var result: string = newStringOfCap(len(packet))
    for m in packet:
        add(result, m)
    
    echo packet
    #let message: string = "Test"
    client.send(packet)

client.close()