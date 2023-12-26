#include <iostream>
#include <map>
#include <vector>

using namespace std;

int main() {
  int t;
  cin >> t;
  while (t--) {
    int l;
    cin >> l;
    vector<char> logs(l);
    for (int i = 0; i < l; i++) {
      cin >> logs[i];
    }
    map<char, int> occ;
    for (int i = 0; i < l; i++) {
      occ[logs[i]]++;
    }
    int sum = 0;
    for (auto const& pair : occ) {
      if (pair.second > (pair.first - 'A')) sum++;
    }
    cout << sum << "\n";
  }
  return 0;
}