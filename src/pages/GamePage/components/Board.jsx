import { useState } from "react"
import { BOARD } from "../../../constants";

import './Board.css'
import { Cell } from "./Cell";
import { mark_spot } from "../GamePage.services";

const getRowCol = (idx) => {
  const row = Math.trunc(idx / 3);
  const col = idx % 3;

  return {
    col,
    row,
  };
}

export const Board = () => {
  const [cells, setCells] = useState(Array(9).fill(BOARD.EMPTY));
  const [xTurn, setXTurn] = useState(true);

  const handleCellClick = (idx) => {
    const newCells = cells.slice();
    if (newCells[idx] != BOARD.EMPTY) return;

    const { row, col } = getRowCol(idx);
    mark_spot(row, col, xTurn);

    newCells[idx] = (xTurn) ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;
    setCells(newCells);

    setXTurn(!xTurn);
  };

  const cellsBtn = cells.map((cell, idx) => (
    <Cell key={idx} value={cell} onClick={handleCellClick} idx={idx} />
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
