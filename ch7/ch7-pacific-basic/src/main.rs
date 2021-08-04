use std::collections::HashMap;

fn main() {
    let mut capitals = HashMap::new();

    capitals.insert("Cook Islands", "Avarua");
    capitals.insert("Fiji", "Suva");
    capitals.insert("Kiribati", "South Tarawa");
    capitals.insert("Tonga", "Nuku'alofa");
    capitals.insert("Tuvalu", "Funafuti");

    let tongan_capital = capitals["Tonga"];
    let tongan_capital = capitals.get("Tonga");

    println!("Capital of Tonga is: {}", tongan_capital.unwrap());
}