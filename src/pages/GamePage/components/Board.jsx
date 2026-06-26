import { useState } from "react";
import { BOARD, CUSTOM_ACTIONS } from "../../../constants";

import "./Board.css";
import { Cell } from "./Cell";
import { getInitialBoard, getResultMessage } from "../GamePage.helpers";
import { ACTION_HANDLERS } from "../GamePage.actionHandlers";

export const Board = () => {
  const [cells, setCells] = useState(getInitialBoard(BOARD.EMPTY, 9));
  const [xTurn, setXTurn] = useState(true);
  const [result, setResult] = useState("");

  const onCellClick = async (idx) => {
    ACTION_HANDLERS[CUSTOM_ACTIONS.CELL_CLICK](idx, cells, setCells, xTurn, onDeclareResult, result);
  }

  const onDeclareResult = (has_won) => {
    ACTION_HANDLERS[CUSTOM_ACTIONS.DECLARE_RESULT](has_won, setResult, getResultMessage);
  }

  const onRestartGame = async() => {
    ACTION_HANDLERS[CUSTOM_ACTIONS.RESTART_GAME](setCells, setXTurn, setResult);
  }

  const cellsBtn = cells.map((cell, idx) => (
    <Cell key={idx} value={cell} onClick={onCellClick} idx={idx} />
  ));

  return (
    <>
      <h1 className="heading">AI TicTacToe</h1>
      <div className="board">{cellsBtn}</div>
      <div className="result">Result: {result}</div>
      <button onClick={onRestartGame}>Restart</button>
    </>
  );
};
