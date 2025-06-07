package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
)

func checkWin() bool {
	winPatterns := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, pattern := range winPatterns {
		a, b, c := pattern[0], pattern[1], pattern[2]
		if board[a] != " " && board[a] == board[b] && board[a] == board[c] {
			return true
		}
	}
	return false
}

func is_draw() {
	for _, cell := range board {
		if cell == " " {
			draw = false
			return
		}
	}
	draw = true
}

func getAvailableMoves() []int {
	availableMoves := []int{}
	for i := 0; i < 9; i++ {
		if board[i] == " " {
			availableMoves = append(availableMoves, i)
		}
	}
	return availableMoves
}

func ai_move(aiPlayer string) {
	availableMoves := getAvailableMoves()
	state := getState()
	key := fmt.Sprintf("('%s', '%s')", state, aiPlayer)

	if len(availableMoves) == 0 {
		// ไม่ควรเกิดขึ้น แต่กันไว้ก่อน
		fmt.Println("No available moves for AI")
		return
	}

	qValues, found := qTable[key]
	if !found {
		// ถ้าไม่เจอ Q-table ให้สุ่ม
		randMove := availableMoves[rand.Intn(len(availableMoves))]
		board[randMove] = aiPlayer
		return
	}

	bestMove := -1
	bestQ := math.Inf(-1)

	for _, move := range availableMoves {
		moveStr := fmt.Sprintf("%d", move)
		q, exists := qValues[moveStr]
		if exists && q > bestQ {
			bestQ = q
			bestMove = move
		}
	}

	if bestMove == -1 {
		// fallback กรณีไม่มี move ที่ match
		bestMove = availableMoves[rand.Intn(len(availableMoves))]
	}

	board[bestMove] = aiPlayer
	currentPlayer = aiPlayer
}

func getState() string {
	state := ""
	for _, cell := range board {
		state += cell
	}
	return state
}

func loadQTable(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read Q-table: %v", err)
		return false
	}
	err = json.Unmarshal(data, &qTable)
	if err != nil {
		log.Fatalf("Failed to unmarshal Q-table: %v", err)
		return false
	}
	return true
}
