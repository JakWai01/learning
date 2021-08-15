use std::thread;

// Race condition
fn main() {
    let mut data = 100;

    thread::spawn(|| { data = 500; });
    thread::spwan(|| { data = 1000; });

    println!("{}", data)
}