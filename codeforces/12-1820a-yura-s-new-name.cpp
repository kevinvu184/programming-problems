#include <iostream>
#include <string>

int main(void) {
  int t{};
  std::cin >> t;
  while (t--) {
    std::string str{};
    std::cin >> str;

    if (str == "^") {
      std::cout << 1 << "\n";
      continue;
    }

    int sum{};
    int dash{};
    for (const char& c : str) {
      if (c == '_') {
        dash++;
      } else if (dash > 0) {
        sum += dash - 1;
        dash = 0;
      }
    }

    if (str[0] != '^') {
      sum += 1;
    }
    if (str[str.size() - 1] != '^') {
      sum += dash;
    }

    std::cout << sum << "\n";
  }
  return 0;
}