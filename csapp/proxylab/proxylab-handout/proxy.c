#include "csapp.h"

void clienterror(int fd, char* cause, char* errnum,
    char* shortmsg, char* longmsg);
void doit(int fd);
void parse_uri(char* uri, char* hostname, char* port, char* path);
void* thread(void* vargp);

struct entry {
    char* key;
    char* value;
};

struct hashmap {
    uint64_t (*hashcode)(const char* key);
    size_t length;
    size_t capactiy;
    struct entry* data;
    pthread_mutex_t lock;
};

static uint64_t hashcode(const char* key);
void map_init(struct hashmap* map);
char* map_get(struct hashmap* map, const char* key);
void map_set(struct hashmap* map, const char* key, const char* value);
void map_destroy(struct hashmap* map);

/* Recommended max cache and object sizes */
#define MAX_CACHE_SIZE 1049000
#define MAX_OBJECT_SIZE 102400

#define FNV_OFFSET 14695981039346656037UL
#define FNV_PRIME 1099511628211UL

/* You won't lose style points for including this long line in your code */
static const char* user_agent_hdr = "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:10.0.3) Gecko/20120305 Firefox/10.0.3\r\n";
static struct hashmap cache;

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
    map_init(&cache);
    while (1) {
        clientlen = sizeof(clientaddr);
        connfdp = Malloc(sizeof(int));
        *connfdp = Accept(listenfd, (struct sockaddr*)&clientaddr, &clientlen);
        Getnameinfo((struct sockaddr*)&clientaddr, clientlen, hostname, MAXLINE, port, MAXLINE, 0);
        printf("Accepted connection from (%s, %s)\n", hostname, port);

        Pthread_create(&tid, NULL, thread, connfdp);
    }
    map_destroy(&cache);
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
    char *buffer;
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
    if ((buffer = map_get(&cache, uri)) != NULL) {
        Rio_writen(fd, buffer, strlen(buffer));
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

    buffer = Malloc(MAX_OBJECT_SIZE);
    memset(buffer, '\0', MAX_OBJECT_SIZE);
    Rio_readinitb(&servrio, servfd);
    while ((n = Rio_readlineb(&servrio, buf, MAXLINE)) != 0) { // line:netp:readhdrs:checkterm
        strncat(buffer, buf, n);
        Rio_writen(fd, buf, n);
    }
    map_set(&cache, uri, buffer);
    free(buffer);
    
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

static uint64_t hashcode(const char* key)
{
    uint64_t hash = FNV_OFFSET;
    for (const char* p = key; *p; p++) {
        hash ^= (uint64_t)(unsigned char)(*p);
        hash *= FNV_PRIME;
    }

    return hash;
}

void map_init(struct hashmap* map)
{
    map->capactiy = 16;
    map->length = 0;
    map->hashcode = &hashcode;
    map->data = Malloc(map->capactiy * sizeof(struct entry));
    memset(map->data, 0, map->capactiy * sizeof(struct entry));
    pthread_mutex_init(&map->lock, NULL);
}

char* map_get(struct hashmap* map, const char* key)
{
    if (map == NULL || map->hashcode == NULL || key == NULL)
        return NULL;

    pthread_mutex_lock(&map->lock);
    uint64_t hash = (map->hashcode)(key);
    size_t index = (size_t)hash & (map->capactiy - 1);

    while (map->data[index].key != NULL) {
        if (strcmp(key, map->data[index].key) == 0) {
            pthread_mutex_unlock(&map->lock);
            return map->data[index].value;
        }

        index++;
        if (index >= map->capactiy) {
            index = 0;
        }
    }

    pthread_mutex_unlock(&map->lock);
    return NULL;
}

void map_set(struct hashmap* map, const char* key, const char* value)
{
    if (value == NULL || key == NULL || map == NULL) {
        return;
    }

    pthread_mutex_lock(&map->lock);
    if (map->length >= map->capactiy / 2) {
        size_t new_capactiy = map->capactiy * 2;
        if (new_capactiy < map->capactiy) {
            pthread_mutex_unlock(&map->lock);
            return;
        }

        map->data = (struct entry*)Realloc(map->data, new_capactiy * sizeof(struct entry));
        map->capactiy = new_capactiy;
    }

    uint64_t hash = (map->hashcode)(key);
    size_t index = (size_t)hash & (map->capactiy - 1);

    while (map->data[index].key != NULL) {
        if (strcmp(key, map->data[index].key) == 0) {
            free(map->data[index].value);
            map->data[index].value = strdup(value);
            pthread_mutex_unlock(&map->lock);
            return;
        }

        index++;
        if (index >= map->capactiy) {
            index = 0;
        }
    }

    map->length++;
    map->data[index].key = strdup(key);
    map->data[index].value = strdup(value);
    pthread_mutex_unlock(&map->lock);
}

void map_destroy(struct hashmap* map)
{
    if (map == NULL)
        return;
    pthread_mutex_lock(&map->lock);
    for (int i = 0; i < map->capactiy && map->data[i].key != NULL; i++) {
        free(map->data[i].key);
        free(map->data[i].value);
    }

    free(map->data);
    map->capactiy = 0;
    map->data = NULL;
    map->length = 0;
    pthread_mutex_unlock(&map->lock);
    pthread_mutex_destroy(&map->lock);
}