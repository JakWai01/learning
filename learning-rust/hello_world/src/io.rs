use std::fs;

pub fn io() {
    let filename = "foo.txt";

    println!("In file {}", filename);

    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");

    println!("With text:\n{}", contents);
}
