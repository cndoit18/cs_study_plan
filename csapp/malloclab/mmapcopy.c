#include <sys/mman.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/types.h>

int main(int argc, char *argv[])
{
	struct stat stat;
	int fd = open(argv[1], O_RDONLY, 0);
	fstat(fd, &stat);
	void *bufp = mmap(NULL, stat.st_size, PROT_READ, MAP_PRIVATE, fd, 0);
	write(1, bufp, stat.st_size);
	munmap(bufp, stat.st_size);
	return 0;
}