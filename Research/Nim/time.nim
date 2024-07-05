import std/[times, os]
# Simple benchmarking
let time = cpuTime()
sleep(100) # Replace this with something to be timed
echo "Time taken: ", cpuTime() - time

# Current date & time
let now1 = now()     # Current timestamp as a DateTime in local time
let now2 = now().utc # Current timestamp as a DateTime in UTC
let now3 = getTime() # Current timestamp as a Time

# Arithmetic using Duration
echo "One hour from now      : ", now() + initDuration(hours = 1)
# Arithmetic using TimeInterval
echo "One year from now      : ", now() + 1.years
echo "One month from now     : ", now() + 1.months