import './App.css'
import TaskSend from "./components/Tasks/TaskSend";
import TaskList from "./components/Tasks/TaskList";


function App() {

    return (
        <div className="App">
            <h1>Expression Parser App</h1>
            <TaskSend/>
            <TaskList/>
        </div>
    )
}

export default App
