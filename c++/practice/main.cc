#include "oop.h"
#include "validate.h"
#include <iostream>
#include <stdexcept>
#include <string>
#include <vector>

inline void default_func(int x = 10) { std::cout << x << std::endl; }

// const与重载
void fcn(int *value) { return; }
void fcn(const int *value) { return; }
// void fcn(int *const value) { return; } error

// TODO: 没有看懂constexpr的逻辑
constexpr int constexpr_func(int x) { return x + 3; }

using std::string;

// struct OOP
struct Sales_data {
  Sales_data() = default;
  // Sales_data() {}

  Sales_data(std::string bookNo) { this->bookNo = bookNo; }
  // this 是一个指针
  std::string isbn() { return this->bookNo; }
  std::string bookNo;
};

// class OOP
class C_Sales_data {
public:
  C_Sales_data() = default;
  C_Sales_data(std::string bookNo) { this->bookNo = bookNo; }
  std::string isbn() { return this->bookNo; }

private:
  std::string bookNo;
};

int main(int argc, char **argv) {

  for (decltype(argc) index = 0; index < argc; index++) {
    std::cout << "args[" << index << "]:" << argv[index] << std::endl;
  }

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

  // int c = 10; error
  // 需要有固定地址，如果是int c将放在栈中
  static int c = 10;
  constexpr int *const ptr_cc = &c;

  string line;

  // int q = 10;
  decltype(i) q = 10;

  while (getline(std::cin, line)) {
    std::cout << "Echo: " << line << std::endl;
  }

  std::vector<int> vectors = {1, 2, 3, 4};
  vectors.push_back(1);

  std::cout << "vectors.size() is :" << vectors.size() << std::endl;
  for (auto index = vectors.begin(); index != vectors.end(); ++index) {
    std::cout << *index << std::endl;
  }

  int arrays[] = {1, 2, 3};
  std::cout << "arrays size is :" << sizeof(arrays) / sizeof(arrays[0])
            << std::endl;

  // 强制转换
  // static_cast
  // const_cast
  // reinterpret_cast 提供较低层次上的重新解释
  int n = 5;
  void *vptr = &n;
  auto sc = static_cast<int *>(vptr);

  // 异常处理

  try {
    throw std::exception();
  } catch (std::exception) {
    std::cerr << "catch an exception" << std::endl;
  }

  default_func();

  constexpr_func(3);
  int val = 1;
  while (val < 10) {
    constexpr_func(++val);
  }

  assert(1 == 1);

  Sales_data sd = Sales_data("test");
  // 没有权限隔离
  std::cout << sd.isbn() << sd.bookNo << std::endl;

  Sales_data empty_sd = Sales_data();

  C_Sales_data cd = C_Sales_data("c_test");
  // std::cout <<  cd.bookNo << std::endl; error bookNo 是私有属性，无法被访问

  OOP op = OOP();
  op.draw();

  friend_oop(op);

  string str = "Hello, World!";
  std::cout << str << std::endl;
  return 0;
}
