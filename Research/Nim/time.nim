import std/[times, os]

# proc multiply(a: int, b: int): int =
#     return a * b

# let x = multiply(3, 4)

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
    echo message, minutes, ':', seconds

