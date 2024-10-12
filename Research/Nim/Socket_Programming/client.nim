import net

let client: Socket = newSocket()
client.connect("127.0.0.1", Port(5555))

while true:
    var packet: array[12, string]

    packet[0] = "0x00"
    packet[1] = "0x00"
    packet[2] = "0x00"
    packet[3] = "0x00"
    packet[4] = "0x00"
    packet[5] = "0x06"
    packet[6] = "0x01"
    packet[7] = "0x06"
    packet[8] = "0x00"
    packet[9] = "0x04"
    packet[10] = "0x00"
    packet[11] = "0x01"


    echo packet
    let message: string = "Test"
    client.send(message)

client.close()
