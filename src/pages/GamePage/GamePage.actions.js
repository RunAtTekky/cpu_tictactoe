import { validate_placement } from "./GamePage.services";

export const validatePlacement = async (row, col, isXturn) => {
  const { canPlace } = await validate_placement(row, col, isXturn);
  return canPlace;
};