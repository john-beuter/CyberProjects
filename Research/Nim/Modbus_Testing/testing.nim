import std/net 
import strutils

echo "What coil would you like to write to?"
var coil = readLine(stdin).parseInt()

let coilString = coil.toHex()

# proc removeLeadingZeros(num: string): string =
 
#     # traverse the entire string
#     for i in 0..<num.len:
#         # check for the first non-zero character
#         if num[i] != '0':
#             # return the remaining string
#             let res = num[i..^1];
#             return res;
         
#     # If the entire string is traversed
#     # that means it didn't have a single
#     # non-zero character, hence return "0"
#     return "0";

proc trimUpper(num: string): string =
  if num.len >= 2:
    return num[^4..^3]  # Slicing the last two characters
  else:
    return num  # Return the entire string if it's less than 2 characters

proc trimLower(num: string): string =
  if num.len >= 2:
    return num[^2..^1]  # Slicing the last two characters
  else:
    return num  # Return the entire string if it's less than 2 characters

echo trimUpper(coilString)
echo trimLower(coilString)
