use std::fs;

#[allow(dead_code)]
static NAME: &str = "[a-zA-Z_][a-zA-Z_0-9]*";
#[allow(dead_code)]
static NUMBER: &str = "(^|\\s+)((\\-?[0-9]+(\\.[0-9]*)?)|(-?\\.[0-9]+))($|\\s+)";
#[allow(dead_code)]
static OPERATOR: &str = "[\\+\\-\\*/]";

#[derive(Debug)]
struct Node {
    value: String,
    left: Option<Box<Node>>,
    right: Option<Box<Node>>,
}

impl Node {
    #[allow(dead_code)]
    fn new(value: String, left: Option<Box<Node>>, right: Option<Box<Node>>) -> Self {
        Self{value: value, left: left, right: right}
    }
}


fn main() {
    let filename = String::from("test.apl");
    
    let node = Node { value: String::from("1"), left: None, right: None }; 

    println!("{:?}", node);

    println!("{}", filename);

    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");

    println!("With text: \n{}", contents);

}
