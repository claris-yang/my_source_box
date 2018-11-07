public class Main {

    public static void main(String[] args) {
        int A [] = new int[] {1,2,3,4};
        Main main = new Main();
        System.out.println(main.numberOfArithmeticSlices(A));

    }

    public int numberOfArithmeticSlices(int[] A) {
        if(A == null || A.length == 0)
            return 0;
        int n = A.length;
        int[] dp = new int[n];
        for(int i=2;i<n;i++){
            if(A[i] - A[i-1] == A[i-1] - A[i-2]){
                dp[i] = dp[i-1] + 1;
                System.out.println(dp[i]);
            }
        }
        int total = 0;
        for(int c : dp){
            System.out.println(c);
            total += c;
        }
        System.out.println("------");
        for(int i = 0 ; i < A.length; i++) {
            System.out.println(dp[i]);
        }
        return total;

    }

    public String removeKdigits(String num, int k) {
        if (k <= 0) {
            return num;
        }
        if (num.length() == 0 || k >= num.length()) {
            return "0";
        }
        int digits = num.length() - k;// 输出字符串的长度
        char[] stk = new char[num.length()];// 模拟桟
        int top = 0;// 记录栈顶元素的下一个位置

        for (int i = 0; i < num.length(); i++) {// 遍历所有元素
            char c = num.charAt(i);
            System.out.printf("%d - %d\n", top, k);
            while (top > 0 && stk[top - 1] > c && k > 0) {// 如果当前元素比栈顶元素小，则出栈
                top--;
                k--;
            }
            System.out.println(c);
            System.out.println(new String(stk));
            stk[top++] = c; // 将当前元素压桟
        }
        // 从头开始查找头个不为0的元素位置

        int idx = 0;
        while (idx < digits && stk[idx] == '0') {
            idx++;
        }

        System.out.println(new String(stk, idx, digits - idx));
        return idx == digits ? "0" : new String(stk, idx, digits - idx);

    }
}
