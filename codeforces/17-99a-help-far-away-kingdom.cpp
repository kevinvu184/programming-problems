#include <iostream>
#include <string>

int main(void) {
  std::string s{};
  std::cin >> s;

  int flag = 0;
  for (std::string::size_type i = 0; i < s.size(); ++i) {
    if (s[i] == '.') {
      if (s[i - 1] == '9') {
        flag = 1;
      } else if (s[i + 1] < '5') {
        flag = 2;
      } else if (s[i + 1] >= '5') {
        flag = 3;
      }
      break;
    }
  }

  if (flag == 0) {
    std::cout << "WRONG!";
  } else if (flag == 1) {
    std::cout << "GOTO Vasilisa.";
  } else if (flag == 2) {
    for (char const& c : s) {
      if (c == '.') {
        break;
      }
      std::cout << c;
    }
  } else if (flag == 3) {
    for (char const& c : s) {
      if (c == '.') {
        break;
      }
      if (*(&c + 1) == '.') {
        std::cout << (char)(c + 1);
      } else {
        std::cout << c;
      }
    }
  }

  return 0;
}