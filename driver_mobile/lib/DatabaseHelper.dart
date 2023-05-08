import 'dart:async';
import 'package:path/path.dart';
import 'package:sqflite/sqflite.dart';

import 'Entities.dart';

class DatabaseHelper {
  static final DatabaseHelper _instance = DatabaseHelper._internal();
  static Database? _database;

  static DatabaseHelper get instance => _instance;

  factory DatabaseHelper() {
    return _instance;
  }

  DatabaseHelper._internal();

  Future<Database> get database async {
    if (_database != null) {
      return _database!;
    }

    _database = await initDatabase();
    return _database!;
  }

  Future<Database> initDatabase() async {
    final String databasesPath = await getDatabasesPath();
    final String path = join(databasesPath, 'driver.db');

    return openDatabase(
      path,
      version: 1,
      onCreate: _onCreate,
    );
  }

  void _onCreate(Database db, int version) async {
    await db.execute('''
      CREATE TABLE Driver (
        id INTEGER PRIMARY KEY,
        name TEXT,
        phone TEXT,
        password TEXT,
        work_exp INTEGER
      )
    ''');

    await db.execute('''
      CREATE TABLE Orders (
        id INTEGER PRIMARY KEY,
        fromCity TEXT,
        toCity TEXT,
        Date TEXT,
        Time TEXT,
        Tickets INTEGER,
        State INT,
        userID INTEGER,
        userName TEXT,
        userPhone TEXT
      )
    ''');
  }

  Future<int> insertDriver(Driver driver) async {
    final db = await database;
    return db.insert(
      'Driver',
      driver.toMap(),
      conflictAlgorithm: ConflictAlgorithm
          .ignore, // добавлен параметр для обработки конфликтов
    );
  }

  Future<int> insertOrder(Orders order) async {
    final db = await database;
    return db.insert(
      'Orders', order.toMap(),
      conflictAlgorithm: ConflictAlgorithm
          .ignore, // добавлен параметр для обработки конфликтов
    );
  }

  Future<int> deleteOrders() async {
    final db = await database;
    return db.delete('Orders');
  }

  Future<int> updateOrder(int orderID, int state) async {
    final db = await database;
    return db.update(
      'Orders',
      {'state': state}, // 1 if true, 0 if false
      where: 'id = ?',
      whereArgs: [orderID],
    );
  }

  Future<List<Orders>> getAllOrders(String date) async {
    final db = await database;
    final List<Map<String, dynamic>> maps = await db.query(
      'Orders',
      where: 'Date = ?',
      whereArgs: [date],
    );
    return List.generate(maps.length, (i) {
      return Orders(
          id: maps[i]['id'],
          fromCity: maps[i]['fromCity'],
          toCity: maps[i]['toCity'],
          date: maps[i]['Date'],
          time: maps[i]['Time'],
          tickets: maps[i]['Tickets'],
          state: maps[i]['State'],
          userID: maps[i]['userID'],
          userName: maps[i]['userName'],
          userPhone: maps[i]['userPhone']);
    });
  }
}
