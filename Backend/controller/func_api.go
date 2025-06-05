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

	c.JSON(http.StatusOK, gin.H{"board": board, "currentPlayer": currentPlayer})
}

func Get_board(c *gin.Context) {
	if !is_started {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Game is not started"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"board":         board,
		"currentPlayer": currentPlayer,
	})
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
	currentPlayer = human_player
	if checkWin() {
		c.JSON(http.StatusOK, gin.H{"board": board, "winner": currentPlayer})
		return
	} else if is_draw() {
		c.JSON(http.StatusOK, gin.H{"board": board, "draw": true})
		return
	}
	ai_move(ai_player)
	if checkWin() {
		c.JSON(http.StatusOK, gin.H{"board": board, "winner": currentPlayer})
		return
	} else if is_draw() {
		c.JSON(http.StatusOK, gin.H{"board": board, "draw": true})
		return
	}
	c.JSON(http.StatusOK, gin.H{"board": board})
}

func Reset_game(c *gin.Context) {
	is_started = false
	board = [9]string{
		" ", " ", " ",
		" ", " ", " ",
		" ", " ", " ",
	}
	currentPlayer = "X"
	c.JSON(http.StatusOK, gin.H{"board": board, "currentPlayer": currentPlayer})
}
