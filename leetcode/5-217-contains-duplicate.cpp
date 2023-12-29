#include <map>
#include <vector>

class Solution {
 public:
  bool containsDuplicate(std::vector<int>& nums) {
    std::map<int, int> occ{};
    for (int const& num : nums)
      if (++occ[num] > 1) return true;
    return false;
  }
};