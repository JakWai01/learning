use rusqlite::{Connection, Result};

#[derive(Debug)]
struct Task {
    id: u32,
    task: String,
    done: u32,
}

#[derive(Debug)]
struct Count {
    count: u32,
}

fn main() -> Result<()> {
    loop {
        use std::io::{stdin, stdout, Write};
        let mut s = String::new();
        println!("--------------------------");
        println!("Please choose an option:");
        println!("[0] Display todo's");
        println!("[1] Create todo");
        println!("[2] Finish todo");
        println!("[3] Exit");
        println!("--------------------------");

        let _ = stdout().flush();
        stdin()
            .read_line(&mut s)
            .expect("Did not enter a correct string");
        if let Some('\n') = s.chars().next_back() {
            s.pop();
        }
        if let Some('\r') = s.chars().next_back() {
            s.pop();
        }
        if s == "0" {
            display()?;
        } else if s == "1" {
            create()?;
        } else if s == "2" {
            finish()?;
        } else if s == "3" {
            break;
        } else {
            println!("Please make a valid guess");
        }
    }

    Ok(())
}

fn display() -> Result<()> {
    println!("Display");
    let conn = connect(String::from("todo.db"));

    let mut stmt = conn.prepare("SELECT t.id, t.task, t.done from todos t;")?;
    let tasks = stmt.query_map([], |row| {
        Ok(Task {
            id: row.get(0)?,
            task: row.get(1)?,
            done: row.get(2)?,
        })
    })?;

    for task in tasks {
        println!("Found task {:?}", task);
    }

    Ok(())
}

fn create() -> Result<()> {
    println!("Create");
    let conn: rusqlite::Connection = connect(String::from("todo.db"));

    let mut s = String::new();
    print!("Enter new task: ");
    use std::io::{stdin, stdout, Write};
    let _ = stdout().flush();
    stdin()
        .read_line(&mut s)
        .expect("Did not enter a correct string");
    if let Some('\n') = s.chars().next_back() {
        s.pop();
    }
    if let Some('\r') = s.chars().next_back() {
        s.pop();
    }

    let mut stmt = conn.prepare("SELECT MAX(t.id) from todos t;")?;
    let values = stmt.query_map([], |row| Ok(Count { count: row.get(0)? }))?;
    let mut max_value = 0;
    for val in values {
        let current_value = val.unwrap().count;
        max_value = current_value + 1;
    }

    conn.execute(
        "INSERT INTO todos (id, task, done) values (?1, ?2, 0)",
        &[&max_value.to_string(), &s],
    )?;

    Ok(())
}

fn finish() -> Result<()> {
    println!("Finish");
    let conn = connect(String::from("todo.db"));

    let mut s = String::new();
    print!("Enter Id of finished todo: ");

    use std::io::{stdin, stdout, Write};
    let _ = stdout().flush();
    stdin()
        .read_line(&mut s)
        .expect("Did not enter a correct string");
    if let Some('\n') = s.chars().next_back() {
        s.pop();
    }
    if let Some('\r') = s.chars().next_back() {
        s.pop();
    }
    conn.execute("UPDATE todos SET done = 1 WHERE id = ?1", &[&s])?;

    Ok(())
}

fn connect(filename: String) -> rusqlite::Connection {
    let conn = Connection::open(filename);

    let conn = conn.unwrap();

    conn.execute(
        "create table if not exists todos (
            id integer primary key,
            task text not null,
            done integer not null
        )",
        [],
    )
    .expect("Error when creating table");

    conn
}
