#include "csapp.h"

void clienterror(int fd, char* cause, char* errnum,
    char* shortmsg, char* longmsg);
void doit(int fd);
void parse_uri(char* uri, char* hostname, char* port, char* path);
void* thread(void* vargp);

/* Recommended max cache and object sizes */
#define MAX_CACHE_SIZE 1049000
#define MAX_OBJECT_SIZE 102400

/* You won't lose style points for including this long line in your code */
static const char* user_agent_hdr = "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:10.0.3) Gecko/20120305 Firefox/10.0.3\r\n";

int main(int argc, char* argv[])
{
    int listenfd, *connfdp;
    char hostname[MAXLINE], port[MAXLINE];
    socklen_t clientlen;
    struct sockaddr clientaddr;
    pthread_t tid;

    if (argc != 2) {
        fprintf(stderr, "usage: %s <port>\n", argv[0]);
        exit(1);
    }

    listenfd = Open_listenfd(argv[1]);

    while (1) {
        clientlen = sizeof(clientaddr);
        connfdp = Malloc(sizeof(int));
        *connfdp = Accept(listenfd, (struct sockaddr*)&clientaddr, &clientlen);
        Getnameinfo((struct sockaddr*)&clientaddr, clientlen, hostname, MAXLINE, port, MAXLINE, 0);
        printf("Accepted connection from (%s, %s)\n", hostname, port);

        Pthread_create(&tid, NULL, thread, connfdp);
    }

    return 0;
}

/* Thread routine */
void* thread(void* vargp)
{
    int connfd = *((int*)vargp);
    Pthread_detach(pthread_self()); // line:conc:echoservert:detach
    Free(vargp); // line:conc:echoservert:free
    doit(connfd);
    Close(connfd);
    return NULL;
}
/* $end echoservertmain */

void doit(int fd)
{
    char buf[MAXLINE], method[MAXLINE], version[MAXLINE], uri[MAXLINE];
    char hostname[MAXLINE], port[MAXLINE], path[MAXLINE];
    rio_t clientrio, servrio;
    int servfd, n;

    Rio_readinitb(&clientrio, fd);
    if (!Rio_readlineb(&clientrio, buf, MAXLINE))
        return;
    printf("%s", buf);

    // GET http://www.cmu.edu/hub/index.html HTTP/1.1
    sscanf(buf, "%s %s %s", method, uri, version);
    if (strcasecmp(method, "GET")) {
        clienterror(fd, method, "501", "Not Implemented",
            "Proxy does not implement this method");
        return;
    }

    parse_uri(uri, hostname, port, path);

    // proxy
    servfd = Open_clientfd(hostname, port);

    sprintf(buf, "%s %s %s\r\n", method, path, version);
    Rio_writen(servfd, buf, strlen(buf));
    Rio_writen(servfd, (void*)user_agent_hdr, strlen(user_agent_hdr));
    sprintf(buf, "Host: %s\r\n", hostname);
    Rio_writen(servfd, buf, strlen(buf));
    sprintf(buf, "Connection: close\r\n");
    Rio_writen(servfd, buf, strlen(buf));
    sprintf(buf, "Proxy-Connection: close\r\n\r\n");
    Rio_writen(servfd, buf, strlen(buf));

    Rio_readinitb(&servrio, servfd);
    while ((n = Rio_readlineb(&servrio, buf, MAXLINE)) != 0) { // line:netp:readhdrs:checkterm
        Rio_writen(fd, buf, n);
    }

    Close(servfd);
}

/*
 * clienterror - returns an error message to the client
 */
/* $begin clienterror */
void clienterror(int fd, char* cause, char* errnum,
    char* shortmsg, char* longmsg)
{
    char buf[MAXLINE];

    /* Print the HTTP response headers */
    sprintf(buf, "HTTP/1.0 %s %s\r\n", errnum, shortmsg);
    Rio_writen(fd, buf, strlen(buf));
    sprintf(buf, "Content-type: text/html\r\n\r\n");
    Rio_writen(fd, buf, strlen(buf));

    /* Print the HTTP response body */
    sprintf(buf, "<html><title>Tiny Error</title>");
    Rio_writen(fd, buf, strlen(buf));
    sprintf(buf, "<body bgcolor="
                 "ffffff"
                 ">\r\n");
    Rio_writen(fd, buf, strlen(buf));
    sprintf(buf, "%s: %s\r\n", errnum, shortmsg);
    Rio_writen(fd, buf, strlen(buf));
    sprintf(buf, "<p>%s: %s\r\n", longmsg, cause);
    Rio_writen(fd, buf, strlen(buf));
    sprintf(buf, "<hr><em>The Tiny Web server</em>\r\n");
    Rio_writen(fd, buf, strlen(buf));
}
/* $end clienterror */

void parse_uri(char* uri, char* hostname, char* port, char* path)
{
    char *ptr, *sign;
    strcpy(port, "80");
    if ((ptr = strstr(uri, "//")) != NULL) {
        ptr += 2;
        for (int offset = 0; offset < strlen(ptr); offset++) {
            if (*(ptr + offset) != '/') {
                continue;
            }
            strncpy(hostname, ptr, offset);
            if ((sign = index(hostname, ':')) != NULL) {
                strcpy(port, sign + 1);
                size_t l = (size_t)(sign - hostname);
                memset(hostname, '\0', MAXLINE);
                strncpy(hostname, ptr, l);
            }
            strcpy(path, ptr + offset);
            return;
        }
    }
}