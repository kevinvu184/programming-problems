#include <iostream>

int main() {
  int t{};
  std::cin >> t;
  while (t--) {
    int n{};
    std::cin >> n;

    std::string x{};
    std::cin >> x;

    int c = 0;
    int left = 0;
    int right = n - 1;
    while (left < right) {
      if (x[left] != x[right]) {
        if (c > 1) {
          c = 3;
          break;
        } else {
          c = 1;
        }
      } else if (c) {
        c = 2;
      }
      left++;
      right--;
    }

    std::cout << (c < 3 ? "YES" : "NO") << "\n";
  }
}