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
        connection.query("CALL GetAllCars()", (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                console.log(result)
                const cars: Car[] = result[0].map((row: any) => ({
                    id: row.id,
                    model: row.model,
                    licPlate: row.lic_plate,
                }));
                resolve(cars);
            }
        });
    });
}

async function InsertCar(model: string, lic_plate: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = `CALL InsertCar(?,?)`
        connection.query(query, [model, lic_plate], (err: MysqlError | null, result: RowDataPacket[]) => {
            if (err) {
                reject(err);
            } else {
                resolve();
            }
        })
    })
}

async function UpdateCar(id: number, model: string, lic_plate: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = `CALL UpdateCar(?,?,?)`
        connection.query(query, [id, model, lic_plate],
            (err: MysqlError | null) => {
                if (err) {
                    reject(err)
                } else {
                    resolve();
                }
            })
    })
}

async function DeleteCar(id: number): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = 'CALL DeleteCar(?)'
        connection.query(query, id, (err: MysqlError | null) => {
            if (err) {
                reject(err)
            } else {
                resolve();
            }
        })
    })
}


export const CarStore = {
    GetAllCars,
    InsertCar,
    UpdateCar,
    DeleteCar
}

