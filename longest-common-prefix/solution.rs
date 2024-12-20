fn main() {
    println!(
        "{}",
        longest_common_prefix(vec![
            "flower".to_string(),
            "flow".to_string(),
            "flight".to_string()
        ])
    );
    println!(
        "{}",
        longest_common_prefix(vec![
            "dog".to_string(),
            "racecar".to_string(),
            "car".to_string()
        ])
    );
    println!("{}", longest_common_prefix(vec!["a".to_string()]));
}

pub fn longest_common_prefix(strs: Vec<String>) -> String {
    return strs
        .into_iter()
        .reduce(|acc, cur| {
            acc.chars()
                .zip(cur.chars())
                .take_while(|(a, c)| a == c)
                .map(|(c, _)| c)
                .collect()
        })
        .unwrap();
}
