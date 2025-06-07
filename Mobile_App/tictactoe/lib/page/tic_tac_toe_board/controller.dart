import 'dart:convert';

import 'package:get/get.dart';
import 'package:http/http.dart';
import 'package:tictactoe/approute/route.dart';
import 'package:tictactoe/page/tic_tac_toe_board/model.dart';
import 'package:tictactoe/utility/element/alert.dart';
import 'package:tictactoe/utility/https.dart';

class GameBoardController extends GetxController {
  GameBoard gameBoard = GameBoard();
  RxBool isLoading = true.obs;

  @override
  void onInit() {
    super.onInit();
    Request_Getboard(); // โหลดตอนเปิด
  }

  Future<void> Request_Getboard() async {
    isLoading.value = true;
    try {
      var response = await get(Uri.parse(Https.Local + Endpoint.Get_Board));
      if (response.statusCode != 200) {
        AlertBox('Error', response.body, Get.context!);
      } else {
        final newData = GameBoard.fromJson(jsonDecode(response.body));
        gameBoard.board.value = newData.board;
        gameBoard.Available_move.value = newData.Available_move;
        gameBoard.Winner.value = newData.Winner.value;
        gameBoard.Draw.value = newData.Draw.value;
      }
    } finally {
      isLoading.value = false;
    }
  }

  Future<void> Apply_Move(int index) async {
    var response = await post(Uri.parse(Https.Local + Endpoint.Apply_Move),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'position': index}));
    if (response.statusCode != 200) {
      AlertBox('Error', response.body, Get.context!);
    } else {
      await Request_Getboard();
    }
  }

  Future<void> Reset_Game() async {
    var response = await post(Uri.parse(Https.Local + Endpoint.Reset));
    if (response.statusCode != 200) {
      AlertBox('Error', response.body, Get.context!);
    } else {
      await Request_Getboard();
      Get.toNamed(Routes.home);
    }
  }
}
