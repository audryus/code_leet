fn main() {
    println!("{}", find_closest_number(vec![-4, -2, 1, 4, 8]));
    println!("{}", find_closest_number(vec![2, -1, 1]));
}
pub fn find_closest_number(nums: Vec<i32>) -> i32 {
    let mut closest = nums[0];
    for &x in nums.iter() {
        if x.abs() < closest.abs() {
            closest = x;
        }
    }

    if closest < 0 && nums.contains(&closest.abs()) {
        return closest.abs();
    }

    return closest;
}
