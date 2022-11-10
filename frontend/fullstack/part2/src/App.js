import Node from './components/Node'

const App = (props) => {
  const { notes } = props

  return (
    <div>
      <h1>Notes</h1>
      <ul>
        {notes.map((note)=>(<Node key={note.id} note={note}/>))}
      </ul>
    </div>
  )
}

export default App;
