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
                reject(err);
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

async function InsertCar(model: string, lic_plate: string): Promise<void> {
    return new Promise((resolve, reject) => {
        const query = `INSERT INTO car(model, lic_plate)
                       VALUES (?, ?)`
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
        const query = `UPDATE car
                       SET model=?,
                           lic_plate=?
                       WHERE id = ?`
        connection.query(query, [model, lic_plate, id],
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
        const query = 'DELETE FROM car WHERE id=?'
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

