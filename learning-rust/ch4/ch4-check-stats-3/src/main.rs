#[derive(Debug)]
struct CubeSat {
    id: u64,
    mailbox: Mailbox
}

#[derive(Debug)]
enum StatusMessage {
    Ok,
}

#[derive(Debug)]
struct Mailbox {
    messages: Vec<Message>,
}

type Message = String;

struct GroundStation;

fn main() {
    GroundStation {};
}
