#include <iostream>
#include <map>
#include <vector>

int main() {
  int t;
  std::cin >> t;
  while (t--) {
    int l;
    std::cin >> l;
    std::vector<char> logs(l);
    for (int i = 0; i < l; i++) {
      std::cin >> logs[i];
    }
    std::map<char, int> occ;
    for (auto log : logs) {
      occ[log]++;
    }
    int sum = 0;
    for (const auto& pair : occ) {
      if (pair.second > (pair.first - 'A')) sum++;
    }
    std::cout << sum << "\n";
  }
  return 0;
}