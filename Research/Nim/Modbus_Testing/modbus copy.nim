import std/[net, times, os]


#[ I want to have functions as follows

ModbusClient(hostIP, port)
Client.blah

]#


let
    ip = ""
    port = Port(502)
    
var
    socket: Socket = newSocket()[]

socket.connect(ip, port)

proc timer(): string =
    const message: string = "Time is "

    let startTime: DateTime = now()
    let countdownTime = initDuration(seconds = 2)

    var elapsedTime: Duration = initDuration()

    while countdownTime > elapsedTime:
        sleep(100)
        let currentTime: DateTime = now()
        elapsedTime = currentTime - startTime
        let timeLeft = countdownTime - elapsedTime
        let totalSeconds: int64 = timeLeft.inSeconds
        let minutes: int64 = totalSeconds div 60
        let seconds: int64 = totalSeconds mod 60
        echo message, minutes, ':', seconds

    return "Completed"



#Create packet to write coil


var packetOn: array[12, uint8] = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x03, 0x23, 0xFF, 0x00];
var packetOff: array[12, uint8] = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x03, 0x23, 0x00, 0x00];

var dataPtr1: ptr array = packetOn.addr
var dataPtr2: ptr array = packetOff.addr
var output: int 

#output = socket.send(dataPtr2, sizeof(packetOff))

output = socket.send(dataPtr1, sizeof(packetOn))

#[while true:
    output = socket.send(dataPtr1, sizeof(packetOn))
    var val: string = timer()
    echo val
    output = socket.send(dataPtr2, sizeof(packetOff))
    var val2: string = timer()

    echo output ]#

