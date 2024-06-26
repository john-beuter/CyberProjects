import net

var packet = @["0x00", "0x00", "0x00", "0x00", "0x00", "0x06", "0x01", "0x06", "0x00", "0x04", "0x00", "0x01"]



proc toString(str: seq[string]): string =
  result = newStringOfCap(len(str))
  for ch in str:
    add(result, ch)
    
echo toString(packet)
    #let message: string = "Test"
