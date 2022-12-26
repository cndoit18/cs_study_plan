#include <iostream>

int main()
{
	int start, end;
	std::cin >> start >> end;

	while(start < end) {
		std::cout << start << std::endl;
		start ++;
	}
	return 0;
}
