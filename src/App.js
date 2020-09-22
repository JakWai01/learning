import React, {useEffect, useState} from "react";
import "./App.css";

function App() {

  const[joke, setJoke] = useState("")

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch("https://api.chucknorris.io/jokes/random");
      const data = await res.json();
          
      setJoke(data.value);
    }

    fetchData()
  }, []);

  return (
    <div>{joke}</div>
  )

}

export default App;