#include <iostream>
#include <map>
#include <string>

int main(void) {
  int t{};
  std::cin >> t;
  while (t--) {
    std::string s{};
    std::cin >> s;
    std::map<char, int> occ{};
    for (char const& c : s) occ[c]++;

    for (std::string::size_type i = 0; i < s.size(); i++) {
      if (--occ[s[i]] == 0) {
        std::cout << s.substr(i) << "\n";
        break;
      }
    }
  }
  return 0;
}