package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Move struct {
	Position int `json:"position"`
}

type StartGame struct {
	Player string `json:"player"`
}

func Start_game(c *gin.Context) {
	var startGame StartGame
	if err := c.ShouldBindJSON(&startGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if startGame.Player != "X" && startGame.Player != "O" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player"})
		return
	} else if startGame.Player == "X" {
		ai_player = "O"

	} else {
		ai_player = "X"
		if !loadQTable("q_table.json") {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load Q-table"})
			return
		}
		ai_move(ai_player)
		currentPlayer = "O"
	}

	human_player = startGame.Player
	is_started = true

	c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": draw})
}

func Get_board(c *gin.Context) {
	if !is_started {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game is not started"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": draw})
}

func Apply_move(c *gin.Context) {
	var move Move
	if !is_started {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game is not started"})
		return
	}
	if err := c.ShouldBindJSON(&move); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if move.Position < 0 || move.Position > 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid position"})
		return
	}
	if board[move.Position] != " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Position is already taken"})
		return
	}
	board[move.Position] = human_player
	if human_player == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
	if checkWin() {
		currentPlayer, winner = human_player, human_player
		c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": draw})
		return
	}
	ai_move(ai_player)
	if ai_player == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
	if checkWin() {
		currentPlayer, winner = ai_player, ai_player
		c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": draw})
		return
	}
	is_draw()
	c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": false})
}

func Reset_game(c *gin.Context) {
	winner = ""
	draw = false
	board = [9]string{
		" ", " ", " ",
		" ", " ", " ",
		" ", " ", " ",
	}
	currentPlayer = "X"
	c.JSON(http.StatusOK, gin.H{"board": board, "available_moves": getAvailableMoves(), "currentPlayer": currentPlayer, "winner": winner, "draw": false})
}
