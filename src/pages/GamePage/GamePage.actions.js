import { gameApi } from "./GamePage.services";

export const validatePlacement = async (row, col, isXturn) => {
  const { canPlace } = await gameApi.validate_placement(row, col, isXturn);
  return canPlace;
};