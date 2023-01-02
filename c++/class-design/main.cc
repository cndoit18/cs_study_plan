#include <iostream>

using namespace std;

template <typename _T> inline bool compare(const _T &x, const _T &y) {
  return x == y;
}

class Sales_data {
public:
  Sales_data() = default;
  Sales_data(string bookNo) = delete;
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
  virtual string Name() { return "Sales_data"; }
  template <typename T> string Other() { return "Other"; }

protected:
  string name = "protecetd";

private:
  string bookNo;
};

// typename/class
template <typename T, typename F = int> class Quota final : public Sales_data {
public:
  virtual string Name() { return "Quota"; }
  // virtual void useTemplate() { shared_ptr<T> data; }
  virtual void useTemplate(typename T::quota_type data) {}
};

class QuotaType {
public:
  // typedef int quota_type;
  using quota_type = int;
};

Sales_data::~Sales_data() { cout << "free" << endl; }

string Sales_data::isbn() const {
  // this->bookNo = "123"; error
  return this->bookNo;
}

auto fcn() { return 1 + 1; }

int main(int argc, char **args) {
  compare(1, 2);
  compare<float>(1.1, 2.2);

  unique_ptr<Sales_data> sd(new Sales_data());
  // 浅拷贝，可能会有野指针
  Sales_data clone = *sd;

  Sales_data left = clone;
  Sales_data right = std::move(clone);
  // Sales_data d("test"); error 字符串构造器被拒绝
  sd.release();
  // cout << clone.ptr << endl;

  Sales_data *q = new Quota<QuotaType>();
  // 回避虚函数
  cout << q->Sales_data::Name() << endl;
  q->Other<int>();

  return 0;
}
