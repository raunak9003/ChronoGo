class Solution {
    public List lexSmallestNegatedPerm(int n, int target) {
        int totalSum = n * (n + 1) / 2;
        
        if (target < -totalSum || target > totalSum) {
            return new ArrayList<>();
        }
        
        int diff = totalSum - target;
        if (diff % 2 != 0) {
            return new ArrayList<>();
        }
        
        int needToNegate = diff / 2;
        
        List result = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            result.add(i);
        }
        
        for (int i = n - 1; i >= 0; i--) {
            if (needToNegate >= result.get(i)) {
                needToNegate -= result.get(i);
                result.set(i, -result.get(i));
            }
            
            if (needToNegate == 0) {
                break;
            }
        }
        
        Collections.sort(result);
        
        return result;
    }
}
