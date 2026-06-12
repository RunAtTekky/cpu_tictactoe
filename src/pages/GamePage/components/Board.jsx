import { useState } from "react"
import { BOARD } from "../../../constants";

import './Board.css'
import { Cell } from "./Cell";

export const Board = () => {
  const [cells, setCells] = useState(Array(9).fill(BOARD.EMPTY));

  const onCellClick = (idx) => {
    console.log("Button clicked");
    const newCells = cells.slice();
    newCells[idx] = BOARD.X_SYMBOL;
    setCells(newCells);
  };

  const cellsBtn = cells.map((cell, idx) => (
    <Cell key={idx} value={cell} onClick={onCellClick} idx={idx} />
  ));

  return (
    <>
      <h1 className="heading">AI TicTacToe</h1>
      <div className="board">
        {cellsBtn}
      </div>
    </>
  );
}
