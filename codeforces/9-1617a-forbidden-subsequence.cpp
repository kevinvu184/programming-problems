#include <iostream>
#include <map>
#include <string>

int main() {
  int t{};
  std::cin >> t;

  while (t--) {
    std::string s1{};
    std::cin >> s1;
    std::string s2{};
    std::cin >> s2;

    for (std::string::size_type i = 0; i < s1.size(); ++i) {
      for (std::string::size_type j = i + 1; j < s1.size(); ++j) {
        if (s1[i] > s1[j]) {
          int tmp = s1[i];
          s1[i] = s1[j];
          s1[j] = tmp;
        }
      }
    }

    std::map<char, int> occ{};
    for (char c : s1) {
      if (c == 'a' || c == 'b' || c == 'c') {
        occ[c]++;
      }
    }

    std::string::size_type d = s1.size();
    for (std::string::size_type i = 0; i < s1.size(); ++i) {
      if (s1[i] != 'a' && s1[i] != 'b' && s1[i] != 'c') {
        d = i;
        break;
      }
    }

    if (occ['a'] == 0 || occ['b'] == 0 || occ['c'] == 0 || s2 != "abc") {
      std::cout << s1 << "\n";
      continue;
    }
    std::cout << std::string(occ['a'], 'a');
    std::cout << std::string(occ['c'], 'c');
    std::cout << std::string(occ['b'], 'b');
    std::cout << s1.substr(d, s1.size()) << "\n";
  }
  return 0;
}