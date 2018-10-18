pub fn find() -> Option<u32> {
    for a in 1..1000 {
        for b in 1..1000 - a {
            if is_pythagorean_triplet(a, b, 1000 - a - b) {
                return Some(a * b * (1000 - a - b));
            }
        }
    }
    None
}

fn is_pythagorean_triplet(a: u32, b: u32, c: u32) -> bool {
    a.pow(2) + b.pow(2) == c.pow(2)
}
