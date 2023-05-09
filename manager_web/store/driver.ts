import connection from "../config/db.config";
import {MysqlError} from "mysql";
import {RowDataPacket} from 'mysql2';

export interface Driver {
    id: number;
    name: string;
    phone: string;
    password: string;
    work_exp: number;
}

async function GetAllDrivers(): Promise<Driver[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT id,name,phone,password,work_exp FROM driver", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const drivers: Driver[] = result.map(row => ({
                    id: row.id,
                    name: row.name,
                    phone: row.phone,
                    password: row.password,
                    work_exp: row.work_exp,
                }));
                resolve(drivers);
            }
        });
    });
}

async function InsertDriver(name: string, phone: string, password: string, work_exp: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = `INSERT INTO driver(name, phone, password, work_exp)
                       VALUES (?, ?, ?, ?)`
        connection.query(query, [name, phone, password, work_exp], (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                resolve();
            }
        })
    })
}

async function UpdateDriver(id: number, name: string, phone: string, password: string, work_exp: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = `UPDATE driver
                       SET name=?,
                           phone=?,
                           password=?,
                           work_exp=?
                       WHERE id = ?`
        connection.query(query, [name, phone, password, work_exp, id],
            (err: MysqlError | null) => {
                if (err) {
                    reject(err)
                } else {
                    resolve();
                }
            })
    })
}

async function DeleteDriver(id: number): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = 'DELETE FROM driver WHERE id=?'
        connection.query(query, id, (err: MysqlError | null) => {
            if (err) {
                reject(err)
            } else {
                resolve();
            }
        })
    })
}

export const DriverStore = {

    GetAllDrivers,
    InsertDriver,
    UpdateDriver,
    DeleteDriver
}

