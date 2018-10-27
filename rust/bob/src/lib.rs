pub fn reply(message: &str) -> &str {
    match message.trim() {
        m if m.is_empty() => "Fine. Be that way!",
        m if is_shouting(m) && m.ends_with('?') => "Calm down, I know what I'm doing!",
        m if m.ends_with('?') => "Sure.",
        m if is_shouting(m) => "Whoa, chill out!",
        _ => "Whatever.",
    }
}

fn is_shouting(message: &str) -> bool {
    message.to_uppercase() == message && message.chars().any(|x| x.is_alphabetic())
}
