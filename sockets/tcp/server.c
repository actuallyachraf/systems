#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#include<sys/socket.h>
#include<arpa/inet.h>
#include<netdb.h>
#include<unistd.h>

#define MAX 256
#define HOST "localhost"
#define IP "127.0.0.1"
#define PORT 1234

struct sockaddr_in server_addr,client_addr;
int mysock, csock;
int r,len,n;

int start_server(){

    printf("> Starting Server...\n");

    printf(">Creating socket\n");
    mysock = socket(AF_INET,SOCK_STREAM,0);

    if (mysock < 0){
    
        printf(">failed to create socket exiting!\n");
        exit(1);
    } 

    printf(">populate sockaddr structure\n");
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = htonl(INADDR_ANY);
    server_addr.sin_port = htons(PORT);

    printf(">binding socket to server address\n");
    r = bind(mysock,(struct sockaddr*)&server_addr,sizeof(server_addr));
    if (r < 0){
        printf(">failure to bind socket exiting!\n");
        exit(1);
    }

    printf(">server initialized on hostname : %s | port : %d\n",HOST,PORT);
    printf(">listening for incomming connections...\n");
    listen(mysock,5);

}

int main(){

    char line[MAX];
    start_server();
    while(1){
    
        printf(">server : accepting new incoming...\n");
        len = sizeof(client_addr);
        csock = accept(mysock,(struct sockaddr*)&client_addr,&len);

        if (csock < 0) {exit(1);}
        printf(">incoming accepted from :\n");
       // printf(">[+]client : IP=%s | PORT=%s\n",inet_ntoa(client_addr.sin_addr.s_addr),ntohs(client_addr.sin_port));
    
        while(1){
        
            n = read(csock,line,MAX);
            if (n==0){
                printf(">client died...\n");
                close(csock);
                break;
            }

            printf("> read n=%d bytes \n line=%s\n",n,line);
            n = write(csock,line,MAX);
            printf("> server wrote back n=%d bytes\t ECHO=%s\n",n,line);
        
        
        }
    }

}


