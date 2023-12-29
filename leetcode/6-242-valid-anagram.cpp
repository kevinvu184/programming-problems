#include <map>
#include <string>
#include <utility>

class Solution {
 public:
  bool isAnagram(std::string s, std::string t) {
    if (s.size() != t.size()) return false;

    std::map<char, int> occ{};
    for (char const& c : s) ++occ[c];
    for (char const& c : t) --occ[c];
    for (std::pair<const char, int> const& o : occ)
      if (o.second > 0) return false;
    return true;
  }
};