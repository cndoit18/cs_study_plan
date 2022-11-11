import Node from "./components/Node";
import { useState, useEffect } from "react";

import axios from "axios";

const App = () => {
  const [notes, setNotes] = useState([]);
  const [newNote, setNewNote] = useState("a new note");
  const [showAll, setShowAll] = useState(true);

  useEffect(() => {
    axios.get("/api/notes").then((response) => {
      setNotes(response.data);
    });
  }, []);

  return (
    <div>
      <h1>Notes</h1>
      <button
        onClick={(event) => {
          setShowAll(!showAll);
        }}
      >
        show {showAll ? "important" : "all"}
      </button>
      <ul>
        {notes
          .filter((note) => showAll || note.important)
          .map((note) => (
            <Node
              key={note.id}
              note={note}
              toggleImportance={() => {
                console.log(note.id);
                note.important = !note.important;
                axios
                  .put(`/api/notes/${note.id}`, note)
                  .then((response) => {
                    setNotes(
                      notes.map((note) =>
                        note.id !== response.data.id ? note : response.data
                      )
                    );
                  });
              }}
            />
          ))}
      </ul>
      <form
        onSubmit={(event) => {
          console.log("submit");
          event.preventDefault();
          const note = {
            content: newNote,
            date: new Date().toISOString(),
            important: Math.random() < 0.5,
          };
          axios
            .post("/api/notes", note)
            .then((response) => {
              setNotes(notes.concat(response.data));
              setNewNote("");
            });
        }}
      >
        <input
          value={newNote}
          onChange={(event) => {
            setNewNote(event.target.value);
          }}
        />
        <button type="submit">save</button>
      </form>
    </div>
  );
};

export default App;
