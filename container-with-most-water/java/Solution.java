

public class Solution {
    public static void main(String[] args) {
        int[] array = new int[] {1,8,6,2,5,4,8,3,7};

        System.out.println(maxArea(array));
    }


    public static int maxArea(int[] height) {
        int left = 0;
        int right = height.length - 1;
        int max = 0;

        while(left < right) {
            int w = right - left;
            int h = Math.min(height[left], height[right]);
            int area = w * h;

            max = Math.max(max, area);

            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }

        }

        return max;
    }

}