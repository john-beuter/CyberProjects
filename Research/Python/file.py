calc = input()
blacklist = [ ';', '"', 'os', '_', '\\', '`', '!', '[', ']', 'import',
              'eval', 'banner', 'echo', 'cat', '%', 
              '&', '>', '<', 'b', 's', 
              'lower', 'upper', 'system', '}', '{' ]


print(type(calc))


if any(char in calc for char in blacklist):
    print("blocked!")
    exit

else:
    calc_result = eval(calc)
    print(calc_result)
