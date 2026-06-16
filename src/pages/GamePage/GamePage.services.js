import { API, CONTENT_TYPE } from "../../constants"

export const mark_spot = async (row, col, isXTurn) => {
    const response = await fetch(API.BASE_URL + API.PLACE, {
        method: "POST",
        headers: CONTENT_TYPE,
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
        game_over: body.game_over,
        has_won: body.has_won,
        success: body.success,
        errMsg: body.errMsg
    }
}

export const validate_placement = async (row, col, isXTurn) => {
    const response = await fetch(API.BASE_URL + API.VALIDATE, {
        method: "POST",
        headers: CONTENT_TYPE,
        body: JSON.stringify({
            row: row,
            col: col,
            x_turn: isXTurn
        })
    });
    if (!response.ok) return;

    const body = await response.json();

    return {
        canPlace: body.can_place,
    }
}

export const restart_game = async () => {
    const response = await fetch(API.BASE_URL + API.RESTART, {
        method: "GET",
        headers: CONTENT_TYPE,
    });

    if (!response.ok) return;
    
    const body = await response.json();

    return {
        success: body.success
    }
}