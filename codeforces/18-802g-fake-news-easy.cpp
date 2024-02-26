#include <iostream>
#include <string>

int main() {
  std::string s;
  std::cin >> s;
  std::string target = "heidi";
  size_t index = 0;
  bool flag = false;
  for (char c : s) {
    if (c == target[index]) {
      ++index;
      if (index == target.size()) {
        flag = true;
      }
    }
  }
  std::cout << (flag ? "YES" : "NO");
  return 0;
}
