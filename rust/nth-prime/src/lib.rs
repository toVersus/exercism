pub fn nth(n: u32) -> u32 {
    (2..)
        .filter(is_prime)
        .take((n + 1) as usize)
        .last()
        .unwrap()
}

pub fn is_prime(n: &u32) -> bool {
    !(2..*n).any(|x| n % x == 0)
}
