#include <iostream>
#include <string>

int main() {
  int size{};
  std::cin >> size;

  std::string aabb{"aabb"};
  for (int i = 0; i < size; ++i) {
    std::cout << aabb[i % aabb.size()];
  }
  std::cout << "\n";

  return 0;
}