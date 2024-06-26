#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <netdb.h>
#include <string.h>
#include <arpa/inet.h>
#include <sys/socket.h>

//Made possible from reference to Thiago Alves defcon23 talk and code

/***
I could expand this to be, give people the idea for how to read the value from the register then tell them
to find out how to write

*/

int main(int argc, char **argv)
{   
    int socket_fd, port = 502;
	struct sockaddr_in server_addr; //What is this?? 
	struct hostent *server;
	int data_len;
	socklen_t cli_len;
    char *hostaddr;


    //Modbus packet to send to the PLC and turn LED on
    char packetOn[] = {0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x00, 0x00, 0xFF, 0x00};
    
    //Modbus packet to send to the PLC and turn LED off
    char packetOff[] = {0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x05, 0x00, 0x00, 0x00, 0x00};

    server = gethostbyname("192.168.13.86");
	socket_fd = socket(AF_INET, SOCK_STREAM, 0);

    bzero((char *) &server_addr, sizeof(server_addr)); //zeros struct
	server_addr.sin_family = AF_INET;
	server_addr.sin_port = htons(port);
	bcopy((char *)server->h_addr, (char *)&server_addr.sin_addr.s_addr, server->h_length);

    while (connect(socket_fd, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0)
	{
			perror("error connecting to server");
			//return 0;
	}

	if (socket_fd<0)
	{
		perror("Server: error creating stream socket");
		exit(1);
	}

    if(atoi(argv[1]) == 1)
    {
        data_len = sendto(socket_fd, packetOn, sizeof(packetOn), 0, (struct sockaddr *)&server_addr, sizeof(server_addr));
        if (data_len < 0)
            {
                printf("Error sending data on socket %d\n", socket_fd);
                perror("error: ");
            }
    }
    else
    {
        data_len = sendto(socket_fd, packetOff, sizeof(packetOff), 0, (struct sockaddr *)&server_addr, sizeof(server_addr));
        if (data_len < 0)
        {
            printf("Error sending data on socket %d\n", socket_fd);
            perror("error: ");
        }
    }
}
