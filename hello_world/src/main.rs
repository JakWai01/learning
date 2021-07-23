#![allow(warnings, unused)]
mod conditionals;
mod hello;
mod loops;

use std::env;

fn main() {
    let key = "FUNC";

    match env::var(key) {
        Ok(val) => {
            if val == "hello" {
                hello::hello();
            } else if val == "loops" {
                loops::loops();
            } else if val == "conditionals" {
                conditionals::conditionals();
            } else {
                println!("There is no function called {} in this crate", val)
            }
        }
        Err(e) => println!("couldn't interpret {}: {}", key, e),
    }
}
