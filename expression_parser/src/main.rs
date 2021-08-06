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

    fn expr(&self, psc: ProgramScanner) -> Self {
        let token: Option<String> = psc.nextTokenOrNull(NUMBER);
        if token != None {
           Self{value: token.unwrap(), left: None, right: None} 
        } else {
            token = psc.nextTokenOrNull("\\(");
            let left: Option<Box<Node>> = self.expr(psc);
            token = psc.nextToken(OPERATOR);
            let value: String = token.unwrap();
            let right: Option<Box<Node>> = self.expr(psc);
            token = psc.nextTokenOrNull("\\)");
            
            Self{value: value, left: left, right: right}
        }
    }

    fn evaluate(&self) -> f64 {
        let mut result: f64 = 0.0;
        match self.left {
            None => self.value.parse().unwrap(),
            _    => {
                let l: f64 = self.left.unwrap().evaluate();
                let r: f64 = self.right.unwrap().evaluate();
                match self.value.as_str() {
                    "+" => result = l + r,
                    "-" => result = l - r,
                    "*" => result = l * r,
                    "/" => result = l / r,
                }
                result
            }
        } 
    }

    fn show(&self) -> () {
        match self.left {
            None => println!("{}", self.value),
            _    => {
                self.left.unwrap().show();
                println!("{}", self.value);
                self.right.unwrap().show();
            }
        } 
    }
}


fn main() {
    let filename = String::from("test.apl");

    let psc: ProgramScanner = ProgramScanner::new(filename).expect("File not found!");
    let result: Node = Node.expr(psc);

    println!("Success!");
    println!("Result = {}", result.evaluate());
    result.show();
}
