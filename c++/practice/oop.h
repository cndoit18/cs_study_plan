#ifndef OOP_HH
#define OOP_HH
#include <iostream>
#include <string>

class OOP {
  friend void friend_oop(OOP oop);

public:
  // OOP() = default;
  // 类似于super("demo")
  OOP() : OOP("demo") { std::cout << "delegating " << this->name << std::endl; }
  OOP(std::string name) : name(name) {}
  void draw() { std::cout << this->name << std::endl; }

private:
  std::string name;
};

// 友元函数的目的是为了访问类里面的私有属性
void friend_oop(OOP oop) {
  std::cout << "from friend: " << oop.name << std::endl;
}

#endif
