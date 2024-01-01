#include <iostream>
#include <limits>
#include <map>
#include <string>

int main(void) {
  int n{};
  int k{};
  std::cin >> n >> k;
  std::string s{};
  std::cin >> s;

  std::map<char, int> occurences{};
  for (char const& c : s) {
    occurences[c]++;
  }

  bool flag = true;
  for (int i = 0; i < k; ++i) {
    if (occurences.count((char)(i + 65)) == 0) {
      flag = false;
    }
  }

  if (flag) {
    int min = std::numeric_limits<int>::max();
    for (auto const& o : occurences) {
      if (o.second < min) {
        min = o.second;
      }
    }
    std::cout << min * k;
  } else {
    std::cout << 0;
  }

  return 0;
}