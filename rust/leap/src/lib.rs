pub fn is_leap_year(year: i32) -> bool {
    let modulo = |x: i32| year % x == 0;
    modulo(4) && (!modulo(100) || modulo(400))
}
