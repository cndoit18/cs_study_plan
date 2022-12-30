#include "validate.h"
#include <iostream>
#include <string>
#include <vector>

using std::string;

int main() {
  // initialization
  int x = 1;

  std::cout << x << std::endl; // 1

  extern int y;
  extern const int z;

  std::cout << y << std::endl; // 2
  std::cout << z << std::endl; // 3

  unsigned int ui = 10;
  int i = -42;

  /* 总是会隐性转化为无符号? */
  std::cout << ui + i << std::endl; // 4294967264
  std::cout << i + ui << std::endl; // 4294967264

  std::cout << i + i << std::endl; // -84

  // 作用域

  {
    int i = 20;
    std::cout << i << std::endl; // 20
  }

  int &ri = i;
  std::cout << i << std::endl; // -42

  ri = 20;
  std::cout << ri << i << std::endl; // 20 20
  // 引用并不产生新的内存空间
  std::cout << (&ri == &i) << std::endl; // 1

  int *ptr = nullptr;
  // int *ptr = NULL;
  std::cout << (ptr == NULL) << std::endl; // 1

  const int *ptr_i = &i;
  // *ptr_i = 20; error
  // ptr_i = ptr; no error

  // const int *const ptr_ci = &i;
  // *ptr_ci = 20; error
  // ptr_ci = ptr; error
  int *const ptr_c = &i;
  // ptr_c = ptr; error
  // *ptr_c = 20; no error

  string line;

  while (getline(std::cin, line)) {
    std::cout << "Echo: " << line << std::endl;
  }

  std::vector<int> vectors = {1, 2, 3, 4};
  vectors.push_back(1);

  std::cout << "vectors.size() is :" << vectors.size() << std::endl;
  for (auto index = vectors.begin(); index != vectors.end(); ++index) {
    std::cout << *index << std::endl;
  }

  string str = "Hello, World!";
  std::cout << str << std::endl;
  return 0;
}
