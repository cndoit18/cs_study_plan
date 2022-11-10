import { useState } from "react";

const Part = ({ part: { name, exercises } }) => {
  return (
    <p onClick={()=>{console.log(`onClick ${name} ${exercises}`)}}>
      {name} {exercises}
    </p>
  );
};
const Header = ({ course: { name } }) => {
  return <h1>{name}</h1>;
};
const Content = ({ course: { parts } }) => {
  return (
    <>
    {parts.map((part)=>(<Part part={part}/>))}
    </>
  );
};
const Total = ({ course: { parts } }) => {
  let total = 0;
  parts.forEach((element) => {
    total += element.exercises;
  });
  return <p>Number of exercises {total}</p>;
};

const App = () => {
  let [now, setDate] = useState(new Date());
  setInterval(() => setDate(new Date()), 1000);
  const course = {
    name: "Half Stack application development",
    parts: [
      {
        name: "Fundamentals of React",
        exercises: 10,
      },
      {
        name: "Using props to pass data",
        exercises: 7,
      },
      {
        name: "State of a component",
        exercises: 14,
      },
    ],
  };
  return (
    <div>
      <Header course={course} />
      <Content course={course} />
      <Total course={course} />
      <p>{now.toString()}</p>
    </div>
  );
};

export default App;
