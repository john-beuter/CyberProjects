import net

let client: Socket = newSocket()
client.connect("192.168.13.86", Port(502))

var packet = @["0x00", "0x00", "0x00", "0x00", "0x00", "0x06", "0x01", "0x06", "0x00", "0x04", "0x00", "0x01"]

proc toString(str: seq[string]): string =
  result = newStringOfCap(len(str))
  for ch in str:
    add(result, ch)
    
var message = toString(packet)
    #let message: string = "Test"

echo message
client.send(message)

client.close()

