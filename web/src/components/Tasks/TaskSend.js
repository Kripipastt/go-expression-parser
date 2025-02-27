import Button from "../UI/Button";
import InputData from "../UI/InputData";
import styles from "./TaskSend.module.css";
import {useState} from "react";
import ExpressionService from "../../API/ExpressionService.ts";

function TaskSend() {
    const [expression, setExpression] = useState("")

    const postExpression = async () => {
        await ExpressionService.sendExpression(expression)
       setExpression("")
    }

    const changeExpression = e => {
        setExpression(e.target.value);
    }

    return <span>
        <div className={styles.div_input}>
        <InputData value={expression} placeholder={"Введите выражение"} type={"text"} onInput={changeExpression}></InputData>
        <Button onClick={postExpression}>Посчитать</Button>
        </div>
    </span>
}

export default TaskSend;
