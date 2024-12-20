fn main() {
    println!(
        "{}",
        is_subsequence("abc".to_string(), "ahbgdc".to_string())
    );
    println!(
        "{}",
        is_subsequence("axc".to_string(), "ahbgdc".to_string())
    );
}
pub fn is_subsequence(s: String, t: String) -> bool {
    if s.len() == 0 {
        return true;
    }

    let s_len = s.len();
    let t_len = t.len();

    if s_len > t_len {
        return false;
    }

    let mut j = 0;

    let s_chars: Vec<char> = s.chars().collect();
    let t_chars: Vec<char> = t.chars().collect();

    for i in 0..t_len {
        if t_chars[i] == s_chars[j] {
            if j == s_len - 1 {
                return true;
            }
            j += 1;
        }
    }
    return false;
}
