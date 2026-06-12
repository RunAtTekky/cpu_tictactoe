import { API } from "../../constants"

export const mark_spot = async (row, col, isXTurn) => {
    const response = await fetch(API.BASE_URL + API.PLACE, {
        method: "POST",
        body: JSON.stringify({
            row: row,
            col: col,
            x_turn: isXTurn
        })
    });
    if (!response.ok) return;

    const body = await response.json();

    return {
        row: body.row,
        col: body.col,
    }
}