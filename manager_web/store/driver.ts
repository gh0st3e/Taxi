import connection from "../config/db.config";
import {MysqlError} from "mysql";
import { RowDataPacket } from 'mysql2';

interface Driver {
    id: number;
    name: string;
    phone: string;
    work_exp: number;
}

async function GetAllDrivers(): Promise<Driver[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT id,name,phone,work_exp FROM driver", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve drivers.'));
            } else {
                const drivers: Driver[] = result.map(row => ({
                    id: row.id,
                    name: row.name,
                    phone: row.phone,
                    work_exp: row.work_exp,
                }));
                resolve(drivers);
            }
        });
    });
}

export const DriverStore = {
    GetAllDrivers
}

