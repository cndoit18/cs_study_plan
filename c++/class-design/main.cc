#include <iostream>

using namespace std;

class Sales_data {
public:
  Sales_data() = default;
  Sales_data(string bookNo);
  // Sales_data(const Sales_data &clone) = default;
  // Sales_data(Sales_data &&clone) { cout << "right" << endl; }

  ~Sales_data();
  string isbn() const;

  Sales_data &operator=(Sales_data &clone) & {
    cout << "left" << endl;
    return *this;
  }
  Sales_data &operator=(Sales_data &clone) && {
    cout << "right" << endl;
    return *this;
  }

private:
  string bookNo;
};

Sales_data::Sales_data(string bookNo) : bookNo(bookNo) {}

Sales_data::~Sales_data() { cout << "free" << endl; }

string Sales_data::isbn() const {
  // this->bookNo = "123"; error
  return this->bookNo;
}

int main(int argc, char **args) {
  unique_ptr<Sales_data> sd(new Sales_data());
  // 浅拷贝，可能会有野指针
  Sales_data clone = *sd;

  Sales_data left = clone;
  Sales_data right = std::move(clone);
  sd.release();
  // cout << clone.ptr << endl;
  return 0;
}
