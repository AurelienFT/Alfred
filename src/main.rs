use std::path::Path;
#[macro_use]
extern crate log;
extern crate env_logger;
extern crate text_io;

use std::io;
use std::io::Write;
use std::process::Command;

fn main() {
    env_logger::init();
    if !Path::new("synapses.json").exists() {
        println!("Train the bot...");
        let output = Command::new("python")
            .args(&["src/neural_network/train_neural_network.py"])
            .status()
            .expect("failed to execute process");
        if !output.success() {
            error!("Traning of bot failed")
        }
        debug!("Successfully trained");
    } else {
        debug!("Bot already trained");
    }
    loop {
        print!("Enter your sentence : ");
        let mut line = String::new();
        io::stdout().flush().unwrap();
        io::stdin()
            .read_line(&mut line)
            .expect("Error getting guess");
        let output = Command::new("python")
            .args(&["src/neural_network/classify.py", &line])
            .output()
            .expect("failed to execute process");
        let word_class = String::from_utf8_lossy(&output.stdout);
        debug!("The class of this word is : {}", word_class);
    }
}
