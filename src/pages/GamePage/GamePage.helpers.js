export const getRowCol = (idx) => {
  const row = Math.trunc(idx / 3);
  const col = idx % 3;

  return {
    row,
    col,
  };
};

export const getIdx = (row, col) => {
  const idx = row * 3 + col;
  return idx;
};

export const isCellEmpty = (cell, EMPTY_VALUE) => {
  return cell === EMPTY_VALUE;
}

export const getInitialBoard = (EMPTY_VALUE, size = 9) => {
  return Array(size).fill(EMPTY_VALUE);
}

export const getResultMessage = (hasWon) => {
  return hasWon ? "CPU has won" : "DRAW";
}