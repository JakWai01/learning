#[macro_use]
extern crate clap;
use clap::App;
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
    let yaml = load_yaml!("cli.yml");
    let matches = App::from(yaml).get_matches();

    if matches.is_present("list") {
        display()?;
    } else if matches.is_present("new") {
        create(String::from(matches.value_of("new").unwrap()))?;
    } else if matches.is_present("done") {
        finish(
            String::from(matches.value_of("done").unwrap())
                .parse::<i32>()
                .unwrap(),
        )?;
    } else if matches.is_present("undo") {
        undo(
            String::from(matches.value_of("undo").unwrap())
                .parse::<i32>()
                .unwrap(),
        )?;
    } else {
        println!("Please make a valid guess");
    }

    Ok(())
}

fn display() -> Result<()> {
    let conn = connect();

    let mut stmt = conn.prepare("SELECT t.id, t.task, t.done from todos t;")?;
    let tasks = stmt.query_map([], |row| {
        Ok(Task {
            id: row.get(0)?,
            task: row.get(1)?,
            done: row.get(2)?,
        })
    })?;

    for task in tasks {
        let current_task = task.unwrap();
        if current_task.done == 0 {
            println!("({:0>2}) [ ] {}", current_task.id, current_task.task);
        } else {
            println!("({:0>2}) [X] {}", current_task.id, current_task.task);
        }
    }

    Ok(())
}

fn create(task: String) -> Result<()> {
    let conn: rusqlite::Connection = connect();

    conn.execute("INSERT INTO todos (task, done) values (?1, 0)", &[&task])?;

    Ok(())
}

fn finish(id: i32) -> Result<()> {
    let conn = connect();

    conn.execute("UPDATE todos SET done = 1 WHERE id = ?1", &[&id])?;

    Ok(())
}

fn undo(id: i32) -> Result<()> {
    let conn = connect();

    conn.execute("UPDATE todos set done = 0 WHERE id = ?1", &[&id])?;

    Ok(())
}

fn connect() -> rusqlite::Connection {
    let conn = Connection::open("todo.db");
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
