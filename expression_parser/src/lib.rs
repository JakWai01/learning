use std::fs::File;
use std::io::{self, prelude::*, BufReader};

#[allow(dead_code)]
static ANYTHING: &str = ".+";

#[allow(dead_code)]
pub struct ProgramScanner {
    filename: String,
    linenumber: i32,
    line: String 
}


impl ProgramScanner {
    pub fn new(filename: String) -> Self {
        ProgramScanner {filename: filename, linenumber: 0, line: String::from("")}
    }

    pub fn next_token_or_null(&self, regex: &str) -> Option<String> {
        None
    }

    pub fn next_token(&self, regex: &str) -> Option<String> {
        None
    }

    pub fn next_line(&mut self) -> io::Result<()> {
        let file = File::open("Cargo.toml")?;
        let reader = BufReader::new(file);
        let mut current_line_number: i32 = 0;

        for line in reader.lines() {
            if current_line_number < self.linenumber {
                current_line_number += 1;
            } else {
                println!("{}", line?);
                break;
            }
        }
        self.linenumber += 1;

        Ok(())
    }
}