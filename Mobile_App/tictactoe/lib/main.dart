import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:tictactoe/approute/pages.dart';
import 'package:tictactoe/approute/route.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  // อาจมี await บางอย่าง
  runApp(Main());
}

class Main extends StatelessWidget {
  const Main({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      debugShowCheckedModeBanner: false,
      initialRoute: Routes.home,
      getPages: pages.routes,
    );
  }
}
