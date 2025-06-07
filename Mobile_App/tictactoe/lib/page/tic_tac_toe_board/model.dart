import 'package:get/get.dart';

class GameBoard {
  RxList<String> board = <String>[].obs;
  RxList<int> Available_move = <int>[].obs;
  RxString Winner = "".obs;
  RxBool Draw = false.obs;
  RxString currentPlayer = "".obs;

  GameBoard();

  factory GameBoard.fromJson(Map<String, dynamic> json) {
    return GameBoard()
      ..board.value = List<String>.from(json['board']) ?? []
      ..currentPlayer.value = json['currentPlayer']
      ..Available_move.value = List<int>.from(json['available_moves'])
      ..Winner.value = json['winner'] ?? ''
      ..Draw.value = json['draw'] ?? false;
  }
}
