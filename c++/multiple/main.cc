#include <iostream>
#include <string>

class Base1 {

public:
  Base1() = default;
  void echo() { std::cout << "from Base1" << std::endl; };
};

class Base2 : public virtual Base1 {

public:
  Base2() = default;
  void echo() { std::cout << "from Base2" << std::endl; };
};

class D2 : public virtual Base1, public Base2 {
  using Base1::Base1;
  using Base2::Base2;

public:
  D2(D2 &clone) : Base2(clone){};
  D2() = default;
};

int main() {
  D2 d;
  d.Base1::echo();
  std::string *sp = new std::string("value");
  // 如何直接使用namespace来调用函数
  sp->std::string::~string();
}
