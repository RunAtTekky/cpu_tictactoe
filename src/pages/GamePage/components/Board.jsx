import { useState } from "react";
import { BOARD } from "../../../constants";

import "./Board.css";
import { Cell } from "./Cell";
import { mark_spot, restart_game } from "../GamePage.services";
import { getIdx, getInitialBoard, getRowCol } from "../GamePage.helpers";
import { validatePlacement } from "../GamePage.actions";

export const Board = () => {
  const [cells, setCells] = useState(getInitialBoard(BOARD.EMPTY, 9));
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

  const restartGame = async () => {
    const { success: restarted } = await restart_game();

    if (!restarted) return;

    alert("Restarted");

    setCells(getInitialBoard(BOARD.EMPTY, 9));
    setXTurn(xTurn => !xTurn);
  }

  const cellsBtn = cells.map((cell, idx) => (
    <Cell key={idx} value={cell} onClick={handleCellClick} idx={idx} />
  ));

  return (
    <>
      <h1 className="heading">AI TicTacToe</h1>
      <div className="board">{cellsBtn}</div>
      <div className="result">Result: {result}</div>
      <button onClick={restartGame}>Restart</button>
    </>
  );
};
