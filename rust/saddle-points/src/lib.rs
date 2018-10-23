pub fn find_saddle_points(input: &[Vec<u64>]) -> Vec<(usize, usize)> {
    input
        .iter()
        .enumerate()
        .flat_map(|(r, row)| row.iter().enumerate().map(move |(c, &v)| (r, c, v)))
        .filter_map(|(r, c, v)| {
            if input.iter().all(|row| v <= row[c]) && input[r].iter().all(|&x| v >= x) {
                Some((r, c))
            } else {
                None
            }
        })
        .collect()
}
