#include <stdio.h>
#include <stdbool.h>
#include <getopt.h>
#include <stdlib.h>
#include "cachelab.h"

int steps = 0;

int hit_counts = 0, miss_counts = 0, eviction_counts = 0;
bool verbose = false;

void usage(char *bin)
{
    printf(
        "Usage: %s [-hv] -s <num> -E <num> -b <num> -t <file>\n"
        "Options:\n"
        "  -h         Print this help message.\n"
        "  -v         Optional verbose flag.\n"
        "  -s <num>   Number of set index bits.\n"
        "  -E <num>   Number of lines per set.\n"
        "  -b <num>   Number of block offset bits.\n"
        "  -t <file>  Trace file.\n\n"
        "Examples:\n"
        "  linux>  %s -s 4 -E 1 -b 4 -t traces/yi.trace\n"
        "  linux>  %s -v -s 8 -E 2 -b 4 -t traces/yi.trace\n",
        bin, bin, bin);
}

typedef struct
{
    int timedata;
    unsigned long address;
    bool vaild;
} cacheline;

typedef struct
{
    int s;
    int b;
    int E;
    unsigned long setmask;
    unsigned long tagmask;
    cacheline *lines;
} cache;

void init_cache(cache *c)
{
    int size = c->E * (1 << c->s);
    c->lines = (cacheline *)malloc(sizeof(cacheline) * size);
    for (int i = 0; i < size; i++)
    {
        c->lines[i].address = 0;
        c->lines[i].timedata = 0;
        c->lines[i].vaild = false;
    }
}

void close_cache(cache *c)
{
    free(c->lines);
    c->lines = NULL;
}

void update_cache(cache *c, unsigned long address)
{
    int begin = ((c->setmask & address) >> c->b) * c->E;
    int end = begin + c->E;
    for (int i = begin; i < end; i++)
    {
        if (c->lines[i].vaild && (c->lines[i].address & c->tagmask) == (address & c->tagmask))
        {
            hit_counts++;
            c->lines[i].timedata = ++steps;
            return;
        }
    }
    for (int i = begin; i < end; i++)
    {
        if (!c->lines[i].vaild)
        {
            miss_counts++;
            c->lines[i].timedata = ++steps;
            c->lines[i].address = address;
            c->lines[i].vaild = true;
            return;
        }
    }

    miss_counts++;
    eviction_counts++;

    cacheline *overlay = NULL;
    for (int i = begin; i < end; i++)
    {
        if (overlay == NULL || overlay->timedata > c->lines[i].timedata)
        {
            overlay = &c->lines[i];
        }
    }
    overlay->address = address;
    overlay->vaild = true;
    overlay->timedata = ++steps;
}

int main(int argc, char *argv[])
{
    int s = 0, E = 0, b = 0;
    FILE *t = NULL;
    int opt;
    while (-1 != (opt = getopt(argc, argv, "hvs:E:b:t:")))
    {
        switch (opt)
        {
        case 'h':
            usage(argv[0]);
            return 0;
        case 't':
            t = fopen(optarg, "r");
            break;
        case 'v':
            verbose = true;
            break;
        case 'E':
            E = atoi(optarg);
            break;
        case 's':
            s = atoi(optarg);
            break;
        case 'b':
            b = atoi(optarg);
            break;
        default:
            break;
        }
    }
    if (s == 0 || E == 0 || b == 0 || t == NULL)
    {
        printf("%s: Missing required command line argument\n", argv[0]);
        usage(argv[0]);
        return 1;
    }

    cache c = {
        s,
        b,
        E,
        (-1 << b) ^ (-1 << (b + s)),
        -1 << (b + s),
        NULL,
    };
    init_cache(&c);
    char operation;
    unsigned long address;
    int size;
    while (fscanf(t, " %c %lx,%d", &operation, &address, &size) > 0)
    {
        switch (operation)
        {
        case 'L':
            update_cache(&c, address);
            break;
        case 'M':
            update_cache(&c, address);
        case 'S':
            update_cache(&c, address);
            break;
        default:
            break;
        }
    }
    fclose(t);
    close_cache(&c);
    printSummary(hit_counts, miss_counts, eviction_counts);
    return 0;
}