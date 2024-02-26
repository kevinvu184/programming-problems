#include <iostream>
#include <map>
#include <string>

bool isanagrampalindrome(std::string s) {
  std::map<char, int> occ{};
  for (char const& c : s) {
    occ[c]++;
  }
  int odd = 0;
  for (auto const& [key, val] : occ) {
    if (key % 2 == 1) {
      odd++;
    }
  }
  return odd > 1;
}

bool ispalindrome(std::string s) {
  bool palindrome = true;
  int left = 0;
  int right = s.size() - 1;
  while (left <= right && palindrome) {
    palindrome = (s[left] == s[right]);
    left++;
    right--;
  }
  if (palindrome) {
    return true;
  }
  return false;
}

int main() {
  int t;
  std::cin >> t;
  while (t--) {
    std::map<char, int> occ{};

    int n{}, k{};
    std::cin >> n >> k;

    std::string s{};
    std::cin >> s;
    for (char const& c : s) {
      occ[c]++;
    }

    for (auto const& [_, val] : occ) {
      k -= val % 2;
    }
    std::cout << ((k >= -1) ? "YES\n" : "NO\n");
  }
}