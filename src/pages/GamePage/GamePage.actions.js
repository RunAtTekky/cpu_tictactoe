import { gameApi } from "./GamePage.services";

export const validatePlacementAction = async (row, col, isXturn) => {
  const { canPlace } = await gameApi.validatePlacement(row, col, isXturn);
  return canPlace;
};