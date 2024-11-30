fn main() {
    println!(
        "{}",
        merge_alternately("abc".to_string(), "pqr".to_string())
    );
    println!(
        "{}",
        merge_alternately("ab".to_string(), "pqrs".to_string())
    );
    println!(
        "{}",
        merge_alternately("abcd".to_string(), "pq".to_string())
    );
}
pub fn merge_alternately(word1: String, word2: String) -> String {
    let a_len = word1.len();
    let b_len = word2.len();

    let mut a = 0;
    let mut b = 0;

    let mut word = 1;

    let mut s = Vec::new();

    while a < a_len && b < b_len {
        if word == 1 {
            s.push(word1.chars().nth(a).unwrap());
            a = a + 1;
            word = 2;
        } else {
            s.push(word2.chars().nth(b).unwrap());
            b = b + 1;
            word = 1;
        }
    }

    while a < a_len {
        s.push(word1.chars().nth(a).unwrap());
        a = a + 1;
    }

    while b < b_len {
        s.push(word2.chars().nth(b).unwrap());
        b = b + 1;
    }

    return s.into_iter().collect();
}
