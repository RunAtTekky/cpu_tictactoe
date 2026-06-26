import { BOARD, CUSTOM_ACTIONS } from "../../constants";
import { validatePlacementAction } from "./GamePage.actions";
import { getIdx, getInitialBoard, getRowCol } from "./GamePage.helpers";
import { gameApi } from "./GamePage.services";

const handleDeclareResult = (has_won, setResult, getResultMessage) => {
  setResult(getResultMessage(has_won));
};

const handleCellClick = async (
  idx,
  cells,
  setCells,
  xTurn,
  onDeclareResult,
  result,
) => {
  // Client side check
  const newCells = cells.slice();
  if (newCells[idx] != BOARD.EMPTY) return;
  if (result !== "") return;

  // Server side check
  const { row: userRow, col: userCol } = getRowCol(idx);
  let canPlace = await validatePlacementAction(userRow, userCol, xTurn);
  if (!canPlace) {
    return;
  }

  const {
    row: ai_row,
    col: ai_col,
    game_over,
    has_won,
  } = await gameApi.placeMarkApi(userRow, userCol, xTurn);

  newCells[idx] = xTurn ? BOARD.X_SYMBOL : BOARD.O_SYMBOL;

  if (game_over) {
    onDeclareResult(has_won);

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

const handleRestartGame = async (setCells, setXTurn, setResult) => {
  const { success: restarted } = await gameApi.restart_game();

  if (!restarted) return;

  setCells(getInitialBoard(BOARD.EMPTY, 9));
  setXTurn(xTurn => !xTurn);
  setResult("");
}


export const ACTION_HANDLERS = {
  [CUSTOM_ACTIONS.DECLARE_RESULT]: handleDeclareResult,
  [CUSTOM_ACTIONS.CELL_CLICK]: handleCellClick,
  [CUSTOM_ACTIONS.RESTART_GAME]: handleRestartGame,
};
