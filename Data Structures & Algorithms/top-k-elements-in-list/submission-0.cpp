class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
         unordered_map<int, int> count;
        for (int n : nums) {
            count[n]++;
        }

        vector<pair<int, int>> a;
        for (const auto& p : count) {
            a.push_back({p.second, p.first});
        }
        sort(a.rbegin(), a.rend());

        vector<int> res;
        for (int i = 0; i < k; ++i) {
            res.push_back(a[i].second);
        }
        return res;
    }
};
    