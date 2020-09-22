import React, { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [dogs, setDogs] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch("https://api.thedogapi.com/v1/images/search");
      const data = await res.json();

      setDogs(data)
    }

    fetchData()
  }, [])

  return (
    <>
      {dogs.map((dog, index) => <img src={dog.url} alt={`Dog ${index}`} key={index} />)}
    </>
  )
}

export default App;