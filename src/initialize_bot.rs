#[cfg(not(any(feature = "facebook")))]
pub fn initialize_bot() {
    return;
}

#[cfg(feature = "facebook")]
#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[cfg(feature = "facebook")]
pub fn initialize_bot() {
    rocket::ignite().mount("/facebook", routes![index]).launch();
}