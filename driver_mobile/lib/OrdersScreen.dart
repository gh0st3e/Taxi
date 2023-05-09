import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:http/http.dart' as http;

import 'DatabaseHelper.dart';
import 'Entities.dart';

class OrdersScreen extends StatefulWidget {
  final int id;

  OrdersScreen({required this.id});

  @override
  _OrdersScreenState createState() => _OrdersScreenState();
}

class _OrdersScreenState extends State<OrdersScreen> {
  String _selectedDate = DateFormat('dd-MM-yyyy').format(DateTime.now());
  List<Orders> _orders = [];



  @override
  void initState() {
    super.initState();
    _getOrdersFromServer();
    _loadOrders();
  }



  Future<void> _getOrdersFromServer() async {
    print(_selectedDate);

    final url = Uri.parse('http://10.0.2.2:9000/driver/orders');
    final headers = {'Content-Type': 'application/json'};
    final data = {'id': widget.id, 'date': _selectedDate};
    final response =
        await http.post(url, headers: headers, body: jsonEncode(data));

    if (response.statusCode == 200) {
      print("OK");
      print(response.body.toString());
      final jsonResponse = json.decode(response.body);
      List<Orders> ordersList = [];

      if (jsonResponse != null) {
        for (var jsonOrder in jsonResponse) {
          ordersList.add(Orders.fromJson(jsonOrder));
        }
      }



// Вносим данные в базу данных
      final dbHelper = DatabaseHelper.instance;
      for (var orders in ordersList) {
        await dbHelper.insertOrder(orders);
      }
    } else {
      print("NOT OK");
      print(response.body.toString());
    }
  }

  Future<void> _loadOrders() async {
    _getOrdersFromServer();
    final dbHelper = DatabaseHelper.instance;
    final orders = await dbHelper.getAllOrders(_selectedDate);
    setState(() {
      _orders = orders;
    });
  }

  Future<void> deleteThisDay() async {
    final dbHelper = DatabaseHelper.instance;
    final res = dbHelper.deleteOrders(_selectedDate);
  }

  void _onDateSelected(String date) {
    setState(() {
      _selectedDate = date;
      _loadOrders();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Orders for $_selectedDate'),
      ),
      body: Column(
        children: [
          Container(
            height: 80,
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                IconButton(
                  icon: Icon(Icons.arrow_left),
                  onPressed: () {
                    final currentDate =
                        DateFormat('dd-MM-yyyy').parse(_selectedDate);
                    final newDate = currentDate.subtract(Duration(days: 1));
                    _onDateSelected(DateFormat('dd-MM-yyyy').format(newDate));
                  },
                ),
                Text(_selectedDate),
                IconButton(
                  icon: Icon(Icons.refresh),
                  onPressed: () async {
                    await deleteThisDay();
                    await _getOrdersFromServer();
                    await _loadOrders();
                    setState(() {});

                  },
                ),
                IconButton(
                  icon: Icon(Icons.arrow_right),
                  onPressed: () {
                    final currentDate =
                        DateFormat('dd-MM-yyyy').parse(_selectedDate);
                    final newDate = currentDate.add(Duration(days: 1));
                    _onDateSelected(DateFormat('dd-MM-yyyy').format(newDate));
                  },
                ),
              ],
            ),
          ),
          Expanded(
            child: ListView.builder(
              itemCount: _orders.length,
              itemBuilder: (context, index) {
                final order = _orders[index];
                return GestureDetector(
                  onTap: () {
                    setState(() {

                      print(order.state.toString());

                    });
                  },
                  child: Container(
                    color: order.state == -1
                        ? Colors.red
                        : order.state == 1
                            ? Colors.white
                            : Colors.green,
                    // Если элемент выбран, устанавливаем зеленый цвет фона
                    child: ListTile(
                      title: Text('${order.fromCity} - ${order.toCity}'),
                      subtitle: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text('${order.time}'),
                          Text('${order.userName} (${order.userPhone})'),
                          Text('${order.tickets} tickets')
                          // отображение имени и телефона пользователя
                        ],
                      ),
                      trailing: PopupMenuButton<int>(
                        onSelected: (int newStatus) {
                          onOrderStatusChanged(order.id, newStatus);
                        },
                        itemBuilder: (BuildContext context) =>
                            <PopupMenuItem<int>>[
                          PopupMenuItem<int>(
                            value: -1,
                            child: Text('НЕ ПРИШЕЛ'),
                          ),
                          PopupMenuItem<int>(
                            value: 1,
                            child: Text('ОЖИДАЕТ'),
                          ),
                          PopupMenuItem<int>(
                            value: 2,
                            child: Text('ЗАБРАЛ'),
                          ),
                        ],
                        child: Text(
                          '${order.state}',
                          style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontSize: 20.0,
                          ),
                        ),
                      ),
                    ),
                  ),
                );
              },

            ),

          ),

        ],
      ),
    );
  }

  Future<void> onOrderStatusChanged(int orderID, int state) async {
    print(orderID);
    print(state);

    final url = Uri.parse('http://10.0.2.2:9000/driver/state');
    final headers = {'Content-Type': 'application/json'};
    final data = {'id': orderID, 'state': state};
    final response =
        await http.put(url, headers: headers, body: jsonEncode(data));
    if (response.statusCode == 200) {
      print("OK");
      print(response.body.toString());
    } else {
      print("NOT OK");
      print(response.body.toString());
    }

    final dbHelper = DatabaseHelper.instance;
    int res = await dbHelper.updateOrder(orderID, state);
    print(res);

    _loadOrders();
  }


}
