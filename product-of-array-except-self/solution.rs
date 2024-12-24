fn main() {
    println!("{:?}", product_except_self(vec![1, 2, 3, 4]));
    println!("{:?}", product_except_self(vec![-1, 1, 0, -3, 3]));
}

pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let mut l_mult = 1;
    let mut r_mult = 1;

    let n = nums.len();

    let mut prefix = vec![0; n];
    let mut suffix = vec![0; n];

    for i in 0..n {
        let j = n - i - 1;
        prefix[i] = l_mult;
        suffix[j] = r_mult;

        l_mult *= nums[i];
        r_mult *= nums[j];
    }

    let mut ans = vec![0; n];
    for i in 0..n {
        ans[i] = prefix[i] * suffix[i];
    }

    ans
}
