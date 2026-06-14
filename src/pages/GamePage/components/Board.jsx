import { useState } from "react";
import { BOARD } from "../../../constants";

import "./Board.css";
import { Cell } from "./Cell";
import { mark_spot, validate_placement } from "../GamePage.services";

const getRowCol = (idx) => {
  const row = Math.trunc(idx / 3);
  const col = idx % 3;

  return {
    row,
    col,
  };
};

const getIdx = (row, col) => {
  const idx = row * 3 + col;
  return idx;
};

export const Board = () => {
  const [cells, setCells] = useState(Array(9).fill(BOARD.EMPTY));
  const [xTurn, setXTurn] = useState(true);
  const [result, setResult] = useState("");

  const handleCellClick = async (idx) => {
    // Client side check
    const newCells = cells.slice();
    if (newCells[idx] != BOARD.EMPTY) return;

    // Server side check
    const { row: userRow, col: userCol } = getRowCol(idx);
    let canPlace = await validatePlacement(userRow, userCol, xTurn);
    if (!canPlace) {
      return;
    }

    const {
      row: ai_row,
      col: ai_col,
      game_over,
      has_won,
    } = await mark_spot(userRow, userCol, xTurn);

    newCells[idx] = xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;

    if (game_over) {
      setXTurn(!xTurn);
      declareResult(has_won);

      if (has_won) {
        const ai_idx = getIdx(ai_row, ai_col);
        newCells[ai_idx] = !xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;
      }
    } else {
      const ai_idx = getIdx(ai_row, ai_col);
      newCells[ai_idx] = !xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;
    }

    setCells(newCells);
  };

  const declareResult = (has_won) => {
    if (has_won) {
      setResult("CPU has won");
    } else {
      setResult("Draw");
    }
  };

  const cellsBtn = cells.map((cell, idx) => (
    <Cell key={idx} value={cell} onClick={handleCellClick} idx={idx} />
  ));

  return (
    <>
      <h1 className="heading">AI TicTacToe</h1>
      <div className="board">{cellsBtn}</div>
      <div className="result">Result: {result}</div>
    </>
  );
};

const validatePlacement = async (row, col, isXturn) => {
  const { canPlace } = await validate_placement(row, col, isXturn);
  return canPlace;
};