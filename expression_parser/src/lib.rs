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

    pub fn read(&self) -> std::io::Result<()> {
        let mut reader = my_reader::BufReader::open("Cargo.toml")?;
        let mut buffer = String::new();

        while let Some(line) = reader.read_line(&mut buffer) {
            println!("{}", line?.trim());
        }

        Ok(())
    }

    pub fn next_token_or_null(&self, regex: &str) -> Option<String> {
        None
    }

    pub fn next_token(&self, regex: &str) -> Option<String> {
        None
    }
}

// Cant implement Copy for this type

// impl Copy for ProgramScanner {}

// impl Clone for ProgramScanner {
//     fn clone(&self) -> ProgramScanner {
//         *self
//     }
// }

mod my_reader {
    use std::{
        fs::File,
        io::{self, prelude::*},
    };

    pub struct BufReader {
        reader: io::BufReader<File>,
    }

    impl BufReader {
        pub fn open(path: impl AsRef<std::path::Path>) -> io::Result<Self> {
            let file = File::open(path)?;
            let reader = io::BufReader::new(file);

            Ok(Self { reader })
        }

        pub fn read_line<'buf>(
            &mut self,
            buffer: &'buf mut String,
        ) -> Option<io::Result<&'buf mut String>> {
            buffer.clear();

            self.reader
                .read_line(buffer)
                .map(|u| if u == 0 { None } else { Some(buffer) })
                .transpose()
        }
    }
}
