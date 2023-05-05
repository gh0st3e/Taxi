import connection from "../config/db.config";
import {MysqlError} from "mysql";
import {RowDataPacket} from 'mysql2';

interface Car {
    id: number;
    model: string;
    licPlate: string;
}

async function GetAllCars(): Promise<Car[] | Error> {
    return new Promise((resolve, reject) => {
        connection.query("SELECT id,model,lic_plate FROM car", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(new Error('Failed to retrieve cars.'));
            } else {
                const cars: Car[] = result.map(row => ({
                    id: row.id,
                    model: row.model,
                    licPlate: row.lic_plate,
                }));
                resolve(cars);
            }
        });
    });
}

export const CarStore = {
    GetAllCars
}

