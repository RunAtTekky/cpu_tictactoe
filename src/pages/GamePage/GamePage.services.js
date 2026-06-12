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

    const body = response.json();

    console.log(body);
    // const { row, col, success, errMsg } = body;
    // console.log(row);
    // console.log(col);
    // console.log(success);
    // console.log(errMsg);
}