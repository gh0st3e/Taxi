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
        connection.query("CALL GetAllOrders(?)", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const orders: Order[] = result[0].map((row: any) => ({
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

async function UpdateOrder(orderID: number, driverID: number, carID: number, date: string, time: string, state: number): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = "CALL UpdateManagerOrder(?,?,?,?,?,?)";
        connection.query(query, [orderID, driverID, carID, date, time, state], (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                resolve();
            }
        });
    });
}

async function GetAllOrdersDateA(state: number): Promise<Order[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("CALL GetAllOrdersDateAsc(?)", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const orders: Order[] = result[0].map((row: any) => ({
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
        connection.query("CALL GetAllOrdersDateDesc(?)", state, (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const orders: Order[] = result[0].map((row: any) => ({
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

