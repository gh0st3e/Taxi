import connection from "../config/db.config";
import {MysqlError} from "mysql";
import { RowDataPacket } from 'mysql2';

interface User {
    id: number;
    name: string;
    phone: string;
    password: string;
    email: string;
    telegram: string;
}

async function GetAllUsers(): Promise<User[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT * FROM user", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve users.'));
            } else {
                const users: User[] = result.map(row => ({
                    id: row.id,
                    name: row.name,
                    phone: row.phone,
                    password: row.password,
                    email: row.email,
                    telegram: row.telegram
                }));
                resolve(users);
            }
        });
    });
}

export const UserStore = {
    GetAllUsers
}
