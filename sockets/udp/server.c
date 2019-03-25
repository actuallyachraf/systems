#include<stdio.h>
#include<stdlib.h>
#include<string.h>

#include<sys/socket.h>
#include<netinet/ip.h>
#include<arpa/inet.h>

#define BUFLEN 256 
#define PORT 2343

char line[BUFLEN];
struct sockaddr_in me,client;
int sock, rlen, clen = sizeof(client);

int main() {

    printf("1> create udp socket\n");
    sock = socket(AF_INET,SOCK_DGRAM,IPPROTO_UDP);

    printf("2> populate sockaddr instances with proper args\n");
    memset((char*)&me,0,sizeof(me)); // cast me to bytes zero the memory
    me.sin_family = AF_INET;
    me.sin_port = htons(PORT);
    me.sin_addr.s_addr = htonl(INADDR_ANY); 

    printf("3> bind the socket to me\n");
    bind(sock, (struct sockaddr*)&me, sizeof(me));

    printf("4> wait and reply");
    while(1) {
    
        memset(line, 0, BUFLEN);
        printf("UDP Server : waiting ...\n");
        // recvfrom sock -> push data into line -> populate client struct
        rlen = recvfrom(sock,line,BUFLEN,0,(struct sockaddr*)&client,&clen);
        printf("> received from [host:port] = [%s:%d]\n",inet_ntoa(client.sin_addr),ntohs(client.sin_port));
        printf("> rlen = %d : line= %s \n",rlen,line);
        // Echo back line
        sendto(sock,line,rlen,0,(struct sockaddr*)&client,clen);
    }
    

}
