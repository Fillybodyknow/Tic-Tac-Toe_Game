import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:tictactoe/approute/route.dart';
import 'package:tictactoe/page/home/controller.dart';

class Home extends StatelessWidget {
  HomeController homeController = Get.put(HomeController());

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.lightBlue,
      body: SizedBox(
          width: double.infinity,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text(
                "Tic Tac Toe Game",
                style: TextStyle(fontSize: 48, color: Colors.white),
              ),
              const SizedBox(
                height: 30,
              ),
              TextButton(
                onPressed: () async {
                  await homeController.Request_StartGame('X');
                },
                style: ButtonStyle(
                  backgroundColor: WidgetStateProperty.all(Colors.white),
                  padding: WidgetStateProperty.all(
                      const EdgeInsets.symmetric(horizontal: 50, vertical: 20)),
                  shape: WidgetStateProperty.all(RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(20))),
                ),
                child: const Text(
                  "เล่นเป็น X",
                  style: TextStyle(fontSize: 18, color: Colors.lightBlue),
                ),
              ),
              const SizedBox(
                height: 30,
              ),
              TextButton(
                onPressed: () async {
                  await homeController.Request_StartGame('O');
                },
                style: ButtonStyle(
                  backgroundColor: WidgetStateProperty.all(Colors.white),
                  padding: WidgetStateProperty.all(
                      const EdgeInsets.symmetric(horizontal: 50, vertical: 20)),
                  shape: WidgetStateProperty.all(RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(20))),
                ),
                child: const Text(
                  "เล่นเป็น O",
                  style: TextStyle(fontSize: 18, color: Colors.lightBlue),
                ),
              )
            ],
          )),
    );
  }
}
