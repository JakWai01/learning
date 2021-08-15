import React, { useState, useCallback } from 'react';

import './App.css';

// [x]Alles was eingegeben wird soll in eine Variable speichern
// [x]Beim Add drücken soll diese Variable in einen "stack" hinzugefügt werden und soll das Todo printen
// [x]Done removed das oberste Item

function App() {

  const [todo, setTodo] = useState("");
  const [todoArr, setTodoArr] = useState([]);

  console.log(todo)

  // useCallback && preventdefault 
  const formSubmitted = useCallback((event) => {
    event.preventDefault();
    console.log("A form was submitted")
    setTodoArr([{ id: todoArr.length + 1, content: todo }, ...todoArr
    ])
    setTodo("")
  }, [todo, todoArr])

  const removeForm = useCallback((event) => {
    event.preventDefault();
    console.log("Element got removed");
    setTodoArr(todoArr.splice(1));
  }, [todoArr])

  console.log(todoArr)

  return (
    <div>
      <form >
        <h1>TODOJO</h1>
        <h2>Note: This Todo app works like a stack</h2>
        {/* value */}
        <input type="text" value={todo} onChange={(e) => (setTodo(e.target.value))}></input>


        <button onClick={formSubmitted}>ADD</button>
        <button onClick={removeForm}>DONE</button> 
      </form>

      <ul>
        {todoArr.map((todo) => <li key={todo.id}>{todo.content}</li>)}
      </ul>
    </div>


  )
}



export default App;
