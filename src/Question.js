import React, { useState } from "react";


const shuffle = array => array.sort(() => Math.random() - 0.5)

function Question({ title = "default", answer = "1", wrong1 = "2", wrong2 = "3", wrong3 = "4" }) {
    const [colorFirst, setColorFirst] = useState("");
    const [colorSecond, setColorSecond] = useState("");
    const [colorThird, setColorThird] = useState("");
    const [colorFourth, setColorFourth] = useState("");

    const rightFirst = () => setColorFirst("rgb(67, 201, 67)");
    const rightSecond = () => setColorSecond("rgb(67, 201, 67)");
    const rightThird = () => setColorThird("rgb(67, 201, 67)");
    const rightFourth = () => setColorFourth("rgb(67, 201, 67)");

    const [answers] = useState(shuffle([answer, wrong1, wrong2, wrong3]))

    return (
        <div>
            <h2>{title}</h2>
            <button onClick={answers[0] === answer ? rightFirst : console.log("Wrong")} style={{ backgroundColor: colorFirst }}>{answers[0]}</button>
            <button onClick={answers[1] === answer ? rightSecond : console.log("Wrong")} style={{ backgroundColor: colorSecond }}>{answers[1]}</button>
            <button onClick={answers[2] === answer ? rightThird : console.log("Wrong")} style={{ backgroundColor: colorThird }}>{answers[2]}</button>
            <button onClick={answers[3] === answer ? rightFourth : console.log("Wrong")} style={{ backgroundColor: colorFourth }}>{answers[3]}</button>
        </div>
    )
}
// <button onClick={right} style={{backgroundColor: colorRight}}>{answer}</button>
export default Question