import { API, CONTENT_TYPE } from "../../constants";

export const gameApi = {
  placeMarkApi: async (row, col, isXTurn) => {
    const response = await fetch(API.BASE_URL + API.PLACE, {
      method: "POST",
      headers: CONTENT_TYPE,
      body: JSON.stringify({
        row: row,
        col: col,
        x_turn: isXTurn,
      }),
    });
    if (!response.ok) throw new Error("Failed to make mark");

    return await response.json();
  },

  validate_placement: async (row, col, isXTurn) => {
    const response = await fetch(API.BASE_URL + API.VALIDATE, {
      method: "POST",
      headers: CONTENT_TYPE,
      body: JSON.stringify({
        row: row,
        col: col,
        x_turn: isXTurn,
      }),
    });
    if (!response.ok) return;

    const body = await response.json();

    return {
      canPlace: body.can_place,
    };
  },

  restart_game: async () => {
    const response = await fetch(API.BASE_URL + API.RESTART, {
      method: "GET",
      headers: CONTENT_TYPE,
    });

    if (!response.ok) return;

    const body = await response.json();

    return {
      success: body.success,
    };
  },
};
