import 'dart:convert';
import 'dart:io';

import 'package:driver_mobile/OrdersScreen.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;

import 'DatabaseHelper.dart';
import 'Entities.dart';

import 'package:connectivity/connectivity.dart';

class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _phoneController = TextEditingController();
  final _passwordController = TextEditingController();

  void _login() async {
    try {
      final phone = _phoneController.text;
      final password = _passwordController.text;

      bool isInternetConnected;
      try {
        final client = new HttpClient();
        client.connectionTimeout = const Duration(seconds: 5);

        final request = await client.get(
          '10.0.2.2',
          9000,
          'main',
        );
        final response =
            await request.close().timeout(const Duration(seconds: 2));
        isInternetConnected = response.statusCode == 200;
      } catch (e) {
        print(e);
        isInternetConnected = false;
      }

      if (isInternetConnected) {
        final url = Uri.parse('http://10.0.2.2:9000/driver/login');
        final headers = {'Content-Type': 'application/json'};
        final data = {'phone': phone, 'password': password};
        final response =
            await http.post(url, headers: headers, body: jsonEncode(data));


        if (response.statusCode == 200) {
          print("OK");
          final jsonResponse = json.decode(response.body);
          final driver = Driver.fromJson(jsonResponse);

          // Вносим данные в базу данных
          final dbHelper = DatabaseHelper.instance;
          await dbHelper.deleteDrivers();
          await dbHelper.insertDriver(driver);

          // Открываем страницу с заказами
          Navigator.push(
            context,
            MaterialPageRoute(
                builder: (context) => OrdersScreen(
                      id: driver.id,
                    )),
          );
        } else {
          print(response.toString());
          final errorResponse = json.decode(response.body);
          final error = errorResponse['error'];
          showDialog(
            context: context,
            builder: (BuildContext context) {
              return AlertDialog(
                title: Text("Ошибка"),
                content: Text(error),
                actions: <Widget>[
                  ElevatedButton(
                    child: Text("Закрыть"),
                    onPressed: () {
                      Navigator.of(context).pop();
                    },
                  ),
                ],
              );
            },
          );
        }
      } else {
        showDialog(
          context: context,
          builder: (BuildContext context) {
            return AlertDialog(
              title: Text("Ошибка"),
              content: Text(
                  "Вы не подключены к интернету.\nВход по последним данным"),
              actions: <Widget>[
                ElevatedButton(
                  child: Text("Закрыть"),
                  onPressed: () async {
                    Navigator.of(context).pop();
                    final dbHelper = DatabaseHelper.instance;
                    final id = await dbHelper.getLocalDriver();
                    if (id == -1) {
                      throw Exception("Данные о водителе не сохранены");
                    } else {
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: (context) => OrdersScreen(
                                  id: id,
                                )),
                      );
                    }
                  },
                ),
              ],
            );
          },
        );
      }
    } catch (e) {
      print(e);
      showDialog(
        context: context,
        builder: (BuildContext context) {
          return AlertDialog(
            title: Text("Ошибка"),
            content: Text(e.toString()),
            actions: <Widget>[
              ElevatedButton(
                child: Text("Закрыть"),
                onPressed: () {
                  Navigator.of(context).pop();
                },
              ),
            ],
          );
        },
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Вход'),
      ),
      body: Padding(
        padding: EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextField(
              controller: _phoneController,
              keyboardType: TextInputType.phone,
              decoration: InputDecoration(
                labelText: 'Телефон',
              ),
            ),
            SizedBox(height: 16.0),
            TextField(
              controller: _passwordController,
              obscureText: true,
              decoration: InputDecoration(
                labelText: 'Пароль',
              ),
            ),
            SizedBox(height: 16.0),
            ElevatedButton(
              onPressed: _login,
              child: Text('Войти'),
            ),
          ],
        ),
      ),
    );
  }
}
