import React from 'react';
import Question from "./Question"
import './App.css';


// What do we need?
// [ ] feature to add question in the app

function App() {
  return (
    <div>
      <h1>Quiz</h1>
      <Question title="What's a color?" answer="Color" wrong1="Car" wrong2="Wheel" wrong3="ATTACK"/>
      <Question title="Will you marry me?" answer="Yes" wrong1="No" wrong2="Who are you" wrong3="Are you seriously asking a random person on the internet you don't even know if they want to marry you?!?!?"/>
      <Question title="What's 2 + 1?" answer="3" wrong1="Yes" wrong2="No" wrong3="2"/>
    </div>
    
  );
}

export default App;
