using ll = long long;

class Solution {
public:
    vector<ll> countStableSubarrays(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        int q = queries.size();

        vector<ll> prefix(n);
        prefix[0] = 1;

        int streak = 1;
        for (int i = 1; i < n; i++) {
            if (nums[i] >= nums[i - 1])
                streak++;
            else
                streak = 1;
            
            prefix[i] = prefix[i - 1] + streak; 
        }

        vector<int> inversion(n);
        inversion[n - 1] = n;

        for (int i = n - 2; i >= 0; i--) 
            if (nums[i] > nums[i + 1])
                inversion[i] = i + 1;
            else 
                inversion[i] = inversion[i + 1];

        vector<ll> ans(q);
        for (int i = 0; i < q; i++) {
            int l = queries[i][0], r = queries[i][1];

            int b = inversion[l];
            if (b > r) {
                int sizeLR = r - l + 1;
                ans[i] = (sizeLR * ll(sizeLR + 1)) / 2;
            } else {
                int sizeLB = b - l; 
                ll countBR = prefix[r] - prefix[b - 1];

                ans[i] = ((sizeLB * ll(sizeLB + 1)) / 2) + countBR;
            }
        }

        return ans; 
    }
};
