use std::collections::HashSet;

pub fn raindrops(n: u32) -> String {
    if n == 1 {
        return String::from("1");
    }

    let mut factors = HashSet::new();
    for x in 1..n + 1 {
        if n % x == 0 {
            factors.insert(x);
        }
    }

    let mut out = String::new();
    let mut matched = false;
    if factors.contains(&3) {
        out.push_str("Pling");
        matched = true;
    }
    if factors.contains(&5) {
        out.push_str("Plang");
        matched = true;
    }
    if factors.contains(&7) {
        out.push_str("Plong");
        matched = true;
    }

    if !matched {
        out.push_str(&*n.to_string());
    }

    return out;
}
