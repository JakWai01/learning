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
    loop {
        use std::io::{stdin, stdout, Write};
        let mut s = String::new();
        println!("--------------------------");
        println!("Please choose an option:");
        println!("[0] Display todo's");
        println!("[1] Create todo");
        println!("[2] Finish todo");
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
            println!("Display");
        } else if s == "1" {
            println!("Create");
            let conn = Connection::open("todo.db")?;

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

            for (task, done) in &todos {
                let mut stmt = conn.prepare("SELECT MAX(t.id) from todos t;")?;
                let values = stmt.query_map([], |row| Ok(Count { count: row.get(0)? }))?;
                let mut max_value = 0;
                for val in values {
                    let current_value = val.unwrap().count;
                    max_value = current_value + 1;
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
        } else if s == "2" {
            println!("Finish");
        } else {
            println!("Please make a valid guess");
        }
    }
}
