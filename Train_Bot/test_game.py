from tictactoeboard import TicTacToeBoard
import pickle
import random

def get_ai_move(board, Q, ai_player):
    state = board.get_state()
    moves = board.get_available_moves()
    if (state, ai_player) in Q:
        q_moves = Q[(state, ai_player)]
        valid_q_moves = {k: v for k, v in q_moves.items() if k in moves}
        if valid_q_moves:
            return max(valid_q_moves, key=lambda x: valid_q_moves[x])
    return random.choice(moves)

def get_human_move(board, human_player):
    moves = board.get_available_moves()
    while True:
        try:
            move = int(input(f"Player {human_player}, choose a move (0-8): "))
            if move in moves:
                return move
            else:
                print("Invalid move. Try again.")
        except ValueError:
            print("Invalid input. Please enter a number.")

def play_game():
    try:
        with open("Train_Bot/q_table.pkl", "rb") as f:
            Q = pickle.load(f)
    except FileNotFoundError:
        print("Q-table not found. กรุณารัน q_learning() เพื่อฝึก AI ก่อน.")
        return

    board = TicTacToeBoard()

    try:
        choose = int(input("Do you want to play as 'X' (1) or 'O' (2)? "))
        if choose not in (1, 2):
            raise ValueError
    except ValueError:
        print("Invalid choice. Please enter 1 or 2.")
        return

    human_player = 'X' if choose == 1 else 'O'
    ai_player = 'O' if human_player == 'X' else 'X'
    current_player = 'X'

    while True:
        board.print_board()
        moves = board.get_available_moves()
        print(f"{current_player}'s turn. Available moves: {moves}")

        if current_player == ai_player:
            move = get_ai_move(board, Q, ai_player)
            print(f"AI chooses move: {move}")
        else:
            move = get_human_move(board, human_player)

        board.apply_move(move, current_player)

        if board.win_pattern():
            board.print_board()
            print(f"Player {current_player} wins!")
            break

        if board.if_draw():
            board.print_board()
            print("It's a draw!")
            break

        current_player = 'O' if current_player == 'X' else 'X'

if __name__ == "__main__":
    play_game()
