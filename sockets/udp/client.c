#include<stdio.h>
#include<stdlib.h>
#include<string.h>

#include<sys/socket.h>
#include<netinet/ip.h>
#include<arpa/inet.h>

#define HOST "127.0.0.1"
#define BUFLEN 256 
#define PORT 2343

char line[BUFLEN];
struct sockaddr_in server;
int sock, rlen, slen = sizeof(server);

int main() {

    printf("1> create udp socket\n");
    sock = socket(AF_INET,SOCK_DGRAM,IPPROTO_UDP);

    printf("2> populate sockaddr instances with proper args\n");
    memset((char*)&server,0,sizeof(server)); // cast me to bytes zero the memory
    server.sin_family = AF_INET;
    server.sin_port = htons(PORT);
    inet_aton(HOST,&server.sin_addr); 

    printf("4> start client");
    while(1) {
 
        printf("> enter a line : \n");
        fgets(line,BUFLEN,stdin);
        line[strlen(line)-1] = 0;
        printf("> sending line...\n");
        sendto(sock,line,strlen(line),0,(struct sockaddr*)&server,slen);
        memset(line,0,BUFLEN); // zero memory for later reuse
        printf("> receiving reply\n");
        rlen = recvfrom(sock,line,BUFLEN,0,(struct sockaddr*)&server,&slen);
        printf("> reply : \n len = %d : line = %s \n",rlen,line);
    }
}
