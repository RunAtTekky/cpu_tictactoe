export const BOARD = {
    EMPTY: '$',
    X_SYMBOL: 'X',
    O_SYMBOL: 'O'
}

export const API = {
    BASE_URL: "http://localhost:8090",
    PLACE: "/place",
    HAS_WON: "/has_won",
    GAME_OVER: "/game_over",
    VALIDATE: "/validate",
    RESTART: "/restart",
}

export const CONTENT_TYPE = {
    'Content-Type': 'application/json'
}

export const CUSTOM_ACTIONS = {
    CELL_CLICK: 'CELL_CLICK',
    DECLARE_RESULT: 'DECLARE_RESULT',
    RESTART_GAME: 'RESTART_GAME',
}