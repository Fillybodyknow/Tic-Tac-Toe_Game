import 'package:get/get.dart';
import 'package:tictactoe/approute/route.dart';
import 'package:tictactoe/page/home/view.dart';
import 'package:tictactoe/page/tic_tac_toe_board/view.dart';

class pages {
  static List<GetPage> routes = [
    GetPage(
      name: Routes.home,
      page: () => Home(),
    ),
    GetPage(name: Routes.game, page: () => GameBoard())
  ];
}
