fn main() {
    println!("{:?}", summary_ranges(vec![0, 1, 2, 4, 5, 7]));
    println!("{:?}", summary_ranges(vec![0, 2, 3, 4, 6, 8, 9]));
}

pub fn summary_ranges(nums: Vec<i32>) -> Vec<String> {
    let mut answer = Vec::new();

    let mut i = 0;
    let size = nums.len();

    while i < size {
        let start = nums[i];

        while i < size - 1 && nums[i] + 1 == nums[i + 1] {
            i += 1;
        }

        let mut s = start.to_string();

        if start != nums[i] {
            s += "->";
            s += &nums[i].to_string()
        }

        answer.push(s);

        i += 1
    }

    return answer;
}
