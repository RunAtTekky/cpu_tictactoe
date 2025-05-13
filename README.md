# CPU TicTacToe
Well well well, I have made the two player TicTacToe game many times. It is very easy, you only need a little bit of logic to create TWO player TicTacToe game.

But when it comes to creating TicTacToe game when the computer will be making the moves, it becomes exponentially difficult.

## What we want?
Make the CPU unbeatable.

After user makes a move, CPU should respond with the best possible move.

For this we will be using [minimax](https://en.wikipedia.org/wiki/Minimax)

### Minimax
Kind of like chess evaluation.
- +val means White winning
- -val means Black winning
- 0 means draw.

We have current game state.

We have many possible move choice, each choice will give us a score.

This score is like the evaluation.

So we need to have a heuristic function which gives us the evaluation.

### Game State
If we can play a move and win the game.
Then the heuristic function should give value 10.

If opponent can play a move and win the game.
Then the heuristic function should give value -10.

We should take the maximum.

The Opponent should take the minimum.

So let's say we have three move choices and these are the following scores for those
- 10 Leads to victory in just one move
- -10 Leads to loss in one move
- 0 Leads to draw in one move

We obviously want to win, so we will take the maximum. On the contrary, our opponent will choose the minimum.

### Code
This is an example of what our heuristic function looks like, we are also using depth in real code, so that the CPU prolongs the losing position, keep in mind that CPU itself will never arrive at a losing position so we will have to give it a losing position.

```go
func heuristic() {
    if game_won_by_player {
        return 10
    }
    else if game_won_by_opponent {
        return -10
    }
    else {
        return 0
    }
}
```

