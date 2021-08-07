use lib::*;
pub mod lib;

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

    fn expr(&self, psc: &ProgramScanner) -> Option<Box<Self>> {
        let mut token: Option<String> = psc.next_token_or_null(NUMBER);
        if token != None {
           let node: Box<Node> = Box::new(Self{value: token.unwrap(), left: None, right: None});
           Some(node)
        } else {
            token = psc.next_token_or_null("\\(");
            let left: Option<Box<Node>> = self.expr(psc);
            token = psc.next_token(OPERATOR);
            let value: String = token.unwrap();
            let right: Option<Box<Node>> = self.expr(psc);
            token = psc.next_token_or_null("\\)");
            
            let node: Box<Node> = Box::new(Self{value: token.unwrap(), left: None, right: None});
            Some(node)
        }
    }

    fn evaluate(&self) -> f64 {
        let mut result: f64 = 0.0;
        match self.left {
            None => result = self.value.parse::<f64>().unwrap(),
            _    => {
                let l: f64 = self.left.as_ref().unwrap().evaluate();
                let r: f64 = self.right.as_ref().unwrap().evaluate();
                match self.value.as_str() {
                    "+" => result = l + r,
                    "-" => result = l - r,
                    "*" => result = l * r,
                    "/" => result = l / r,
                    _   => println!("Invalid operation")
                }
            }
        } 
        result
    }

    fn show(&self) -> () {
        match self.left {
            None => println!("{}", self.value),
            _    => {
                self.left.as_ref().unwrap().show();
                println!("{}", self.value);
                self.right.as_ref().unwrap().show();
            }
        } 
    }
}


fn main() {
    let filename = String::from("test.apl");

    let psc: ProgramScanner = ProgramScanner::new(filename);

    psc.read();
}
    
