import connection from "../config/db.config";
import {MysqlError} from "mysql";
import {RowDataPacket} from 'mysql2';

interface Order {
    id: number;
    userID: string;
    driverID: string;
    carID: string;
    fromCity: string;
    toCity: string;
    date: string;
    tickets: number;
    state: number;
}

async function GetAllOrders(state: number): Promise<Order[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT * FROM orders WHERE STATE=?", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve orders.'));
            } else {
                const orders: Order[] = result.map(row => ({
                    id: row.id,
                    userID: row.UserID,
                    driverID: row.DriverID,
                    carID: row.CarID,
                    fromCity: row.FromCity,
                    toCity: row.ToCity,
                    date: row.Date,
                    tickets: row.Tickets,
                    state: row.State,
                }));
                resolve(orders);
            }
        });
    });
}

async function UpdateOrder(orderID: number, driverID: number, carID: number, date: string, time: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = "UPDATE orders SET driverID = ?, carID = ?, date = ?, time = ?, state = 1 WHERE orders.id = ?";
        connection.query(query, [driverID, carID, date, time, orderID], (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error(`Failed to update order with ID ${orderID}.`));
            } else {
                resolve();
            }
        });
    });
}

async function GetAllOrdersDateA(state: number): Promise<Order[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT * FROM orders WHERE STATE=? ORDER BY orders.date ASC", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve orders.'));
            } else {
                const orders: Order[] = result.map(row => ({
                    id: row.id,
                    userID: row.UserID,
                    driverID: row.DriverID,
                    carID: row.CarID,
                    fromCity: row.FromCity,
                    toCity: row.ToCity,
                    date: row.Date,
                    tickets: row.Tickets,
                    state: row.State,
                }));
                resolve(orders);
            }
        });
    });
}

async function GetAllOrdersDateD(state: number): Promise<Order[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT * FROM orders WHERE STATE=? ORDER BY orders.date DESC", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve orders.'));
            } else {
                const orders: Order[] = result.map(row => ({
                    id: row.id,
                    userID: row.UserID,
                    driverID: row.DriverID,
                    carID: row.CarID,
                    fromCity: row.FromCity,
                    toCity: row.ToCity,
                    date: row.Date,
                    tickets: row.Tickets,
                    state: row.State,
                }));
                resolve(orders);
            }
        });
    });
}

export const OrderStore = {
    GetAllOrders,
    UpdateOrder,
    GetAllOrdersDateA,
    GetAllOrdersDateD
}

