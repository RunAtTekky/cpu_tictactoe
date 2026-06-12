import { Board } from "./components/Board"

import styles from './GamePage.module.css'

export const GamePage = () => {
  return (
    <div className={styles.gamePage}>
      <Board />
    </div>
  )
}
