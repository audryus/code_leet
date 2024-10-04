import java.util.*;
import java.util.stream.*;

public class Solution {
    public static void main(String[] args) {
        var arr = new int[] {0,1,0,3,12};
        moveZeroes(arr) ;
        System.out.println(Arrays.toString(arr));
        //System.out.println(moveZeroes(new int[] {0}));
    }
    
    public static void moveZeroes(int[] nums) {
        int nextNonZero  = 0;
        for (int i = 0; i< nums.length; i++) {
            if (nums[i] != 0) {
                var tmp = nums[i];
                nums[i] = nums[nextNonZero];
                nums[nextNonZero] = tmp;
                nextNonZero++;
            }
        }
    }
}