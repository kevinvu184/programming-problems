#include <iostream>
#include <string>

int main() {
  std::string s{};
  std::cin >> s;

  int c = 0;
  for (std::string::size_type i = 0; i < s.size(); ++i) {
    if (c == 6) {
      break;
    }
    if (i + 1 < s.size() && s[i] == s[i + 1]) {
      c++;
    } else if (i + 1 != s.size()) {
      c = 0;
    }
  }
  std::cout << (c >= 6 ? "YES" : "NO");
  return 0;
}
