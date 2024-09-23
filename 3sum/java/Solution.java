import java.util.*;

public class Solution {
    public static void main(String[] args) {

        System.out.println(threeSum(new int[]{-1, 0, 1, 2, -1, -4}));
        //System.out.println(threeSum(new int[]{0, 1, 1}));
        //System.out.println(threeSum(new int[]{0, 0, 0}));

    }

    public static List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> r = new ArrayList<List<Integer>>();

        for (int i = 0; i < nums.length; i++) {
            if (i > 0 && nums[i] == nums[i-1]) {
                continue;
            }

            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                int ln = nums[left];
                int rn = nums[right];
                int sum = nums[i] + ln + rn;
                
                if (sum < 0) {
                    left++;
                } else if(sum > 0) {
                    right--;
                } else {
                    r.add(Arrays.asList(nums[i], ln, rn));
                    left++;

                    while (nums[left] == nums[left-1] && left < right) {
                        left++;
                    }
                }
            }
        }

        return r;
    }

}