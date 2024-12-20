use std::collections::HashMap;

fn main() {
    println!("{}", roman_to_int("III".to_string()));
    println!("{}", roman_to_int("LVIII".to_string()));
    println!("{}", roman_to_int("MCMXCIV".to_string()));
}

pub fn roman_to_int(s: String) -> i32 {
    let chars: Vec<char> = s.chars().collect();
    let mut d = HashMap::new();
    d.insert('I', 1);
    d.insert('V', 5);
    d.insert('X', 10);
    d.insert('L', 50);
    d.insert('C', 100);
    d.insert('D', 500);
    d.insert('M', 1000);

    let n = s.len();
    let mut i = 0;
    let mut summ = 0;

    while i < n {
        if i < n - 1 && d[&chars[i]] < d[&chars[i + 1]] {
            summ += d[&chars[i + 1]] - d[&chars[i]];
            i += 2;
        } else {
            summ += d[&chars[i]];
            i += 1;
        }
    }
    return summ;
}
