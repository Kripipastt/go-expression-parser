import styles from './Button.module.css'

function Button(props) {
  const { children, disabled = false, onClick} = props

  return (
    <button {...props} className={styles.button} disabled={disabled} onClick={onClick}>
      {children}
    </button>
  )
}

export default Button
