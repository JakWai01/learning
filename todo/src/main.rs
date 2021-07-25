use rusqlite::{Connection, Result};
use std::collections::HashMap;

#[derive(Debug)]
struct Task {
    id: u32,
    task: String,
    done: u32,
}

fn main() -> Result<()> {
    let conn = Connection::open("todo.db")?;

    conn.execute(
        "create table if not exists cat_colors (
             id integer primary key,
             name text not null unique
         )",
        [],
    )?;

    conn.execute(
        "create table if not exists todos (
            id integer primary key,
            task text not null,
            done integer not null
        )",
        [],
    )?;
    let mut todos = HashMap::new();
    todos.insert(String::from("Clean Room"), 0);
    todos.insert(String::from("Code"), 1);

    // Get maximum ID instead
    let mut count = 0;
    for (task, done) in &todos {
        conn.execute(
            "INSERT INTO todos (id, task, done) values (?1, ?2, ?3)",
            &[&count.to_string(), &task, &done.to_string()],
        )?;

        count += 1;
    }
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
