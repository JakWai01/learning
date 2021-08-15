// Attempting to modify an iterator while iterating over it
fn main() {
    let mut letters = vec![
        "a", "b", "b"
    ];

    for letter in letters {
        println!{"{}", letter};
        letters.push(letter.clone());
    }
}
