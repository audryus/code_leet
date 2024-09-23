import java.util.*;

public class Solution {
    public static void main(String[] args) {
        System.out.println(triangleNumber(new int[] {2,2,3,4}));
        System.out.println(triangleNumber(new int[] {2,3,4,4}));
    }


    public static int triangleNumber(int[] nums) {
        Arrays.sort(nums);
        int count = 0;
        for (int i = nums.length - 1; i > 1; i--) {
            int left = 0;
            int right = i - 1;

            while(left < right) {
                if (nums[left] + nums[right] > nums[i]) {
                    count += right - left;
                    right--;
                } else {
                    left++;
                }
            }
        }

        return count;
    }
}