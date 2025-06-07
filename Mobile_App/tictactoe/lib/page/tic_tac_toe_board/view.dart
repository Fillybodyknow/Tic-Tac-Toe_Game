import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:get/get.dart';
import 'package:tictactoe/page/tic_tac_toe_board/controller.dart';

class GameBoard extends StatelessWidget {
  final GameBoardController gameBoardController =
      Get.put(GameBoardController());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.lightBlue,
      body: Center(
        child: Obx(() {
          if (gameBoardController.isLoading.value) {
            return const CircularProgressIndicator();
          }

          return Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text(
                "Tic Tac Toe Game",
                style: TextStyle(fontSize: 48, color: Colors.white),
              ),
              Container(
                margin: const EdgeInsets.only(top: 24),
                width: 300,
                height: 300,
                child: GridView.builder(
                  gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 3,
                  ),
                  itemCount: 9,
                  itemBuilder: (context, index) {
                    final value = gameBoardController.gameBoard.board[index];
                    return GestureDetector(
                      onTap: () {
                        if (!gameBoardController.gameBoard.Available_move
                                .contains(index) ||
                            gameBoardController
                                .gameBoard.Winner.value.isNotEmpty) return;

                        gameBoardController.Apply_Move(index);
                      },
                      child: Container(
                        margin: const EdgeInsets.all(4),
                        decoration: BoxDecoration(
                          color: Colors.white,
                          border: Border.all(color: Colors.black),
                        ),
                        child: Center(
                          child: Text(
                            value,
                            style: const TextStyle(
                                fontSize: 32, fontWeight: FontWeight.bold),
                          ),
                        ),
                      ),
                    );
                  },
                ),
              ),
              const SizedBox(height: 24),
              Obx(() => Text(
                    gameBoardController.gameBoard.Winner.value.isNotEmpty
                        ? "Winner: ${gameBoardController.gameBoard.Winner.value}"
                        : gameBoardController.gameBoard.Draw.value
                            ? "Draw"
                            : "",
                    style: const TextStyle(
                        fontSize: 36,
                        fontWeight: FontWeight.bold,
                        color: Colors.white),
                  )),
              const SizedBox(height: 24),
              TextButton(
                onPressed: () async {
                  await gameBoardController.Reset_Game();
                },
                style: ButtonStyle(
                  backgroundColor: WidgetStateProperty.all(Colors.white),
                  padding: WidgetStateProperty.all(
                      const EdgeInsets.symmetric(horizontal: 50, vertical: 20)),
                  shape: WidgetStateProperty.all(RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(20))),
                ),
                child: const Text(
                  "Reset",
                  style: TextStyle(fontSize: 18, color: Colors.lightBlue),
                ),
              ),
            ],
          );
        }),
      ),
    );
  }
}
