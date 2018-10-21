#[cfg(not(any(feature = "facebook")))]
pub fn display_message(message: String) {
    println!("test {}", message);
}

#[cfg(feature = "facebook")]
pub fn display_message(message: String) {
    println!("test fb {}", message);
}