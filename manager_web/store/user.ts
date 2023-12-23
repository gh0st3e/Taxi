import connection from "../config/db.config";
import {MysqlError} from "mysql";
import {RowDataPacket} from "mysql2"

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
        connection.query("CALL GetAllUsers()", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                const users: User[] = result[0].map((row: any) => ({
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

