import { useState, useEffect } from "react";
import ExpressionService from "../../API/ExpressionService.ts";
import Task from "./Task";
import styles from "./TaskList.module.css"

function TaskList() {
    const [expressions, setExpressions] = useState([])

    useEffect(() => {
        const interval = setInterval(async () => {
            await loadTask()
        }, 2000);

        return () => setInterval(interval)
    }, []);

    const loadTask = async () => {
        setExpressions(await ExpressionService.getAllExpressions())
    }
    if (expressions.length === 0) {return }
    return <div align={"center"} className={styles.div}>
        {expressions.map((el) => (<Task key={el.id} expression={el} />))}
    </div>
}

export default TaskList;
