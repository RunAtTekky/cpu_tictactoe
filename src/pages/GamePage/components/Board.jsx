import { useState } from "react"
import { BOARD } from "../../../constants";

import './Board.css'
import { Cell } from "./Cell";

export const Board = () => {
  const [cells, setCells] = useState(Array(9).fill(BOARD.EMPTY));

  const onCellClick = () => {
    console.log("Button clicked");
    const newCells = cells.slice();
    newCells[0] = BOARD.X_SYMBOL;
    setCells(newCells);
  };

  const cellsBtn = cells.map((cell, idx) => <Cell key={idx} value={cell} onClick={onCellClick} />);

  return (
    <>
      <h1>AI TicTacToe</h1>
      {cellsBtn}
    </>
  );
}
