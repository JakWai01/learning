#[macro_use]
extern crate clap;
use clap::App;
use rusqlite::{Connection, Result};
use std::io::{stdin, stdout, Write};

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
        let yaml = load_yaml!("cli.yml");
        let matches = App::from(yaml).get_matches();
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
    let conn: rusqlite::Connection = connect(String::from("todo.db"));

    let mut s = String::new();
    print!("Enter new task: ");
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

    conn.execute("INSERT INTO todos (task, done) values (?1, 0)", &[&s])?;

    Ok(())
}

fn finish() -> Result<()> {
    let conn = connect(String::from("todo.db"));
    let mut s = String::new();
    print!("Enter Id of finished todo: ");

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
            id integer primary key autoincrement,
            task text not null,
            done integer not null
        )",
        [],
    )
    .expect("Error when creating table");

    conn
}
