from tictactoeboard import TicTacToeBoard
import pickle
import random
import os

def q_learning():
    Q = {}
    learning_rate = 0.2
    discount_factor = 0.9
    epsilon = 0.2
    episodes = 1000000

    x_win = 0
    o_win = 0
    draw = 0

    for episode in range(episodes):
        game = TicTacToeBoard()
        state = game.get_state()
        current_player = "X"
        history = []  # เก็บ (state, player, move)

        os.system('cls')
        print(f"Episode: {episode + 1}")


        done = False

        initial_epsilon = 0.4
        final_epsilon = 0.05
        epsilon = initial_epsilon * ((episodes - episode) / episodes) + final_epsilon * (episode / episodes)
        while not done:
            if (state, current_player) not in Q:
                Q[(state, current_player)] = {a: 0 for a in game.get_available_moves()}

            if random.random() < epsilon:
                move = random.choice(game.get_available_moves())
            else:
                move = max(Q[(state, current_player)], key=Q[(state, current_player)].get)

            game.apply_move(move, current_player)
            next_state = game.get_state()

            history.append((state, current_player, move))
            done = game.win_pattern() or game.if_draw()

            # game.print_board()

            state = next_state
            current_player = 'O' if current_player == 'X' else 'X'

        # เมื่อเกมจบ -> สร้าง reward table
        if game.win_pattern():
            winner = 'O' if current_player == 'X' else 'X'  # คนก่อนหน้าคือผู้ชนะ
            if winner == 'X':
                x_win += 1
            elif winner == 'O':
                o_win += 1
            # print(f"Player {winner} wins!")
        elif game.if_draw():
            winner = "Draw"
            draw += 1
            # print("It's a draw!")
        else:
            winner = None

        # Backward update ทุกการเดิน
        for state, player, move in reversed(history):
            if (state, player) not in Q:
                Q[(state, player)] = {a: 0 for a in range(9) if a not in state}

            if winner == "Draw":
                reward = 0.5
            elif winner == player:
                reward = 1
            else:
                reward = -1

            old_q = Q[(state, player)].get(move, 0)
            Q[(state, player)][move] = old_q + learning_rate * (reward - old_q)
            reward = discount_factor * reward  # ลดค่ารางวัลลงในตาที่ย้อนกลับไป

    with open("q_table.pkl", "wb") as f:
        pickle.dump(Q, f)

    print("Training complete. Q-table saved to q_table.pkl")
    print(f"X wins: {x_win}")
    print(f"O wins: {o_win}")
    print(f"Draws: {draw}")

if __name__ == "__main__":
    q_learning()
