use rusqlite::{Connection, Result};
use std::collections::HashMap;

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
    let mut stmt = conn.prepare("SELECT max(t.id) from todos t;")?;

    for (task, done) in &todos {
        let mut max_value = 0;

        let values = stmt.query_map([], |row| Ok(Count { count: row.get(0)? }))?;
        for val in values {
            // println!("{:?}", val.unwrap().count);
            let current_value = val.unwrap().count;
            if current_value > max_value {
                max_value = current_value + 1;
            }
        }
        conn.execute(
            "INSERT INTO todos (id, task, done) values (?1, ?2, ?3)",
            &[&max_value.to_string(), &task, &done.to_string()],
        )?;
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
