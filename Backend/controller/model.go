package controller

var board = [9]string{
	" ", " ", " ",
	" ", " ", " ",
	" ", " ", " ",
}

var currentPlayer = "X"

var human_player = ""
var ai_player = ""
var winner = ""
var draw = false
var is_started = false

var qTable map[string]map[string]float64
