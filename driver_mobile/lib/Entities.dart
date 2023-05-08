import 'package:intl/intl.dart';

class Driver {
  int id;
  String name;
  String phone;
  String password;
  int workExp;

  Driver(
      {required this.id,
      required this.name,
      required this.phone,
      required this.password,
      required this.workExp});

  Map<String, dynamic> toMap() {
    return {
      'id': id,
      'name': name,
      'phone': phone,
      'password': password,
      'work_exp': workExp,
    };
  }

  factory Driver.fromJson(Map<String, dynamic> json) {
    return Driver(
      id: json['ID'],
      name: json['Name'],
      phone: json['Phone'],
      password: json['Password'],
      workExp: json['WorkExp'],
    );
  }
}

class Orders {
  int id;
  String fromCity;
  String toCity;
  String date;
  String time;
  int tickets;
  int state;
  int userID;
  String userName;
  String userPhone;

  Orders(
      {required this.id,
      required this.fromCity,
      required this.toCity,
      required this.date,
      required this.time,
      required this.tickets,
      required this.state,
      required this.userID,
      required this.userName,
      required this.userPhone});

  factory Orders.fromJson(Map<String, dynamic> json) {
    return Orders(
        id: json['ID'],
        fromCity: json['From'],
        toCity: json['To'],
        date: json['Date'],
        time: json['Time'],
        tickets: json['Tickets'],
        state: json["State"],
        userID: json['User']['ID'],
        userName: json['User']['Name'],
        userPhone: json['User']['Phone']);
  }

  Map<String, dynamic> toMap() {
    return {
      'id': id,
      'fromCity': fromCity,
      'toCity': toCity,
      'Date': date,
      'Time': time,
      'Tickets': tickets,
      "State":state,
      'userID': userID,
      'userName': userName,
      'userPhone': userPhone
    };
  }
}
