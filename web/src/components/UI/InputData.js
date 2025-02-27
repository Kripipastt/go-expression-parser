import styles from './InputData.module.css'

function InputData(props) {
    // const {type, placeholder, handleInput, textValue} = props

    return <input {...props} className={styles.input}></input>
}

export default InputData
