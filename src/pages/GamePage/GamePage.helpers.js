import { BOARD } from "../../constants";

export const getRowCol = (idx) => {
  const row = Math.trunc(idx / 3);
  const col = idx % 3;

  return {
    row,
    col,
  };
};

export const getWinningHighlights = (cells) => {
  const isHighlighted = {};
  const highlightedIdxs = getHighlightedIdxs(cells);

  for (let idx=0; idx<9; idx++) {
    isHighlighted[idx] = highlightedIdxs.includes(idx) ? true : false;
  }

  return isHighlighted;
}

const getHighlightedIdxs = (cells) => {
  let highlightedIdxs = [];

  // Horizontal
  for (let i=0; i<=6; i+=3) {
    if (cells[i] !== BOARD.EMPTY && cells[i] === cells[i+1] && cells[i] === cells[i+2]) {
      highlightedIdxs.push(i, i+1, i+2);
    }
  }

  // Vertical
  for (let i=0; i<3; i++) {
    if (cells[i] !== BOARD.EMPTY && cells[i] === cells[i+3] && cells[i] === cells[i+6]) {
      highlightedIdxs.push(i, i+3, i+6);
    }
  }

  // Diagonal
  if (cells[0] !== BOARD.EMPTY && cells[0] === cells[4] && cells[0] === cells[8]) {
    highlightedIdxs.push(0, 4, 8);
  }

  if (cells[2] !== BOARD.EMPTY && cells[2] === cells[4] && cells[2] === cells[6]) {
    highlightedIdxs.push(2, 4, 6);
  }
  
  return highlightedIdxs;
}

export const getIdx = (row, col) => {
  const idx = row * 3 + col;
  return idx;
};

export const isCellEmpty = (cell, EMPTY_VALUE) => {
  return cell === EMPTY_VALUE;
};

export const getInitialBoard = (EMPTY_VALUE, size = 9) => {
  return Array(size).fill(EMPTY_VALUE);
};

export const getResultMessage = (hasWon) => {
  return hasWon ? "CPU has won" : "DRAW";
};