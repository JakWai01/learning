import React, { useState } from 'react';

import './App.css';

// [x]Alles was eingegeben wird soll in eine Variable speichern
// [ ]Beim Add drücken soll diese Variable in einen "stack" hinzugefügt werden und soll das Todo printen
// [ ]Haken checken removed das oberste Item

function App() {
  const [todoArr, setTodoArr] = useState([])
  const [todo, setTodo] = useState("");
  console.log({todo})
  console.log({todoArr})
  
  return (
    <div>
      <h1>TODOJO</h1>
      <h2>Note: This Todo app works like a stack</h2>
      <input type="text" onChange={ e => setTodo(e.target.value)}></input>
      
      <button onClick={(todo) => setTodoArr(todoArr => [...todoArr, todo])}>ADD</button>
      
      <button>DONE</button>
    </div>
    

  )
}



export default App;
