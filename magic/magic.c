#include<stdio.h>
#include<stdlib.h>

int main(int argc, char* argv[]){

    FILE *fp; // file pointer   
    char BUF[4]; // Magic is 4 bytes
    int n = 0;
    // check for file 
    fp = fopen(argv[1],"r");
    // exit if not available
    if (fp == NULL) exit(2);
    
    n = fread(BUF,1,4,fp);
    printf("bytes read = %d\n",n);

    for (int i = 0;i < n;i++) {
        printf("%X ",BUF[i]);
    }
    printf("\n");

    fclose(fp);

}
