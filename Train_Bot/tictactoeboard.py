class TicTacToeBoard:
    
    def __init__ (self): # จะถูกเรียกอัตโนมัติเมื่อสร้างอ็อบเจกต์
        self.board = [' '] * 9
    
    def resetgame(self): # รีเกม
        self.board = [' '] * 9

    def get_available_moves(self): #คืนรายการ index ของช่องที่ยังว่าง
        return [i for i, spot in enumerate(self.board) if spot == ' ']
    
    def apply_move(self, move, player): #ให้ผู้เล่นลง X หรือ O ที่ตำแหน่งที่เลือก (ถ้าว่าง)
        if self.board[move] == ' ':
            self.board[move] = player 
            return True
        return False
    
    def win_pattern(self):# ตรวจสอบว่ามีผู้ชนะ
        win_patterns = [
            (0, 1, 2), (3, 4, 5), (6, 7, 8),
            (0, 3, 6), (1, 4, 7), (2, 5, 8),
            (0, 4, 8), (2, 4, 6)
        ]
        return any(
            self.board[a] == self.board[b] == self.board[c] != ' '
            for a, b, c in win_patterns
        )
    
    def if_draw(self):#ช่องเต็มแล้วยังไม่มีผู้ชนะให้เสมอ
        return ' ' not in self.board
    
    def get_state(self):#สร้าง string แสดงสถานะกระดานเตรียม Q-Learning
        return ''.join(self.board)
    
    def print_board(self):
        print("Current board:")
        for i, row in enumerate([self.board[i*3:(i+1)*3] for i in range(3)]):
           print('| ' + ' | '.join(f'{row[j] or " "}' for j in range(3)) + ' |    ' +
                 '| ' + ' | '.join(str(i*3 + j) for j in range(3)) + ' |')