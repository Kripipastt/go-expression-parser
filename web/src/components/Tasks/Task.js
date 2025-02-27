import styles from './Task.module.css';
import Loader from "../UI/Loader";

function Task(props) {
    const {expression} = props

    return <div className={styles.div} data-status={expression.status}>
        {/*<div>{expression.id}</div>*/}
        <div
            className={styles.div_expression}>{expression.expression}{expression.status === "finish" ? ` = ${expression.result}` : ""}</div>
        {expression.status === "create" || expression.status === "pending" ?
            <div className={styles.div_loader}>{expression.status === "create" ?
                <Loader slowSpin={true}/> : <Loader/>}</div> : undefined}
    </div>
}

export default Task;
