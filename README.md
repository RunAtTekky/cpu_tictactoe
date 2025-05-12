## Algorithm
We want to create an algorithm to give us the best move after user moves

We need to purge other choices.

We can use minimax.

Kind of like chess evaluation. +val means White winning, -val means Black winning and 0 means draw.

So we need to have a heuristic function which tells us the value of current state.

If we can play a move and win the game.
Then the heuristic function should give value 1.

If opponent can play a move and win the game.
Then the heuristic function should give value -1.

Probably like take the average????
NOOOOO!!!!!!!

We should take the maximum.

The Opponent should take the minimum.

So let's say we have three choices
- Leads to victory
- No victory (Opponent wins next turn)
- No victory (Opponent wins next turn)

```go
func heuristic() {
    if game_won_by_player {
        return 1
    }
    else if game_won_by_opponent {
        return -1
    }
    else {
        return 0
    }
}
```

