import { useState } from "react"
import { BOARD } from "../../../constants";

import './Board.css'
import { Cell } from "./Cell";
import { mark_spot } from "../GamePage.services";

const getRowCol = (idx) => {
  const row = Math.trunc(idx / 3);
  const col = idx % 3;

  return {
    row,
    col,
  };
}

const getIdx = (row, col) => {
  const idx = row * 3 + col;
  return idx;
}

export const Board = () => {
  const [cells, setCells] = useState(Array(9).fill(BOARD.EMPTY));
  const [xTurn, setXTurn] = useState(true);
  const [result, setResult] = useState("");

  const handleCellClick = async (idx) => {
    const newCells = cells.slice();
    if (newCells[idx] != BOARD.EMPTY) return;

    const { row: userRow, col: userCol } = getRowCol(idx);
    const {
      row: ai_row,
      col: ai_col,
      game_over,
      has_won
    } = await mark_spot(userRow, userCol, xTurn);

    newCells[idx] = xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;

    if (game_over) {
      declareResult(has_won);
      setXTurn(!xTurn);
      setCells(newCells);
      return;
    }

    const ai_idx = getIdx(ai_row, ai_col);

    console.log("HELLO");
    console.log(ai_row, ai_col);


    newCells[ai_idx] = !xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;

    setCells(newCells);
  };

  const declareResult = (has_won) => {
    if (has_won) {
      setResult(xTurn + "has won")
    } else {
      setResult("Draw")
    }
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
      <div className="result">Result: {result}</div>
    </>
  );
}
