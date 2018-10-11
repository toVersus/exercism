extern crate chrono;
use chrono::{DateTime, Duration, Utc};

const GIGASEC: i64 = 1000000000;

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    start
        .checked_add_signed(Duration::seconds(GIGASEC))
        .unwrap()
}
