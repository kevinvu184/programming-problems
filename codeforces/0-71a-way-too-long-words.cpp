#include <iostream>
#include <format>

int main() {
  int tc;
  std::cin >> tc;
  while (tc--) {
    std::string s;
    std::cin >> s;
    int len = s.length();
    std::string r;
    if (len > 10) {
      r = s.front() + std::to_string(s.length()-2) + s.back();
    } else {
      r = s;
    }
    std::cout << r << "\n";
  }
  return 0;
}