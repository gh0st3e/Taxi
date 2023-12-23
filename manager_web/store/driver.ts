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
        connection.query("CALL GetAllDrivers()", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const drivers: Driver[] = result[0].map((row: any) => ({
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
        const query = `CALL InsertDriver(?,?,?,?)`
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
        const query = `CALL UpdateDriver(?,?,?,?,?)`
        connection.query(query, [id, name, phone, password, work_exp],
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
        const query = 'CALL DeleteDriver(?)'
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

