#include <iostream>
#include <memory>

using namespace std;

shared_ptr<string> func() { return make_shared<string>("from func"); }

int main() {
  auto ptr = func();
  cout << *ptr << endl;

  // 类似malloc 和 free
  auto malloc_ptr = new string("malloc");
  delete malloc_ptr;
  malloc_ptr = nullptr;

  unique_ptr<string> up(new string("unique"));
  up.release();
  assert(up == nullptr);

  unique_ptr<int[]> arrays(new int[10]);

  allocator<string> alloc;
  auto const p = alloc.allocate(2);
  auto q = p;
  // c++17 废弃
  alloc.construct(q++, "test1");
  alloc.construct(q, "test2");
  while (q != (p - 1)) {
    cout << *q << endl;
    alloc.destroy(q--);
  }

  return 0;
}
