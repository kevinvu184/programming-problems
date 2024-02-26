#include <iostream>
#include <map>
#include <string>

int main(void) {
  std::string f{};
  std::string s{};
  std::string t{};
  std::cin >> f >> s >> t;
  std::map<char, char> layout{};
  for (std::string::size_type i = 0; i < f.size(); ++i) {
    layout[f[i]] = s[i];
  }
  for (char const& c : t) {
    if ('A' <= c && c <= 'Z') {
      std::cout << (char)(layout[c - 'A' + 'a'] - 'a' + 'A');
    } else if ('a' <= c && c <= 'z') {
      std::cout << layout[c];
    } else {
      std::cout << c;
    }
  }
  return 0;
}