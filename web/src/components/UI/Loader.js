import styles from "./Loader.module.css"

function Loader(props) {
    const {slowSpin} = props

    return <div className={[styles.loader, slowSpin? styles.loader_slow:""].join(" ")} ></div>
}

export default Loader
