import 'package:get/get.dart';
import 'package:http/http.dart';
import 'package:tictactoe/approute/route.dart';
import 'package:tictactoe/utility/https.dart';
import 'package:tictactoe/utility/element/alert.dart';
import 'dart:convert';
import 'package:tictactoe/page/tic_tac_toe_board/controller.dart';

class HomeController extends GetxController {
  GameBoardController gameBoardController = Get.put(GameBoardController());
  // ignore: non_constant_identifier_names
  Future<void> Request_StartGame(String Player) async {
    var response = await post(
      Uri.parse(Https.Local + Endpoint.StartGame),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'player': Player}),
    );
    if (response.statusCode != 200) {
      AlertBox('Error', response.body, Get.context!);
    } else {
      await gameBoardController.Request_Getboard();
      Get.toNamed(Routes.game);
    }
  }
}
