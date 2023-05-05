import * as dotenv from 'dotenv'
import mysql, {Pool} from "mysql";

dotenv.config()

const DB_CONFIG = {
    HOST: process.env.DB_HOST,
    USER: process.env.DB_USER,
    PASSWORD: process.env.DB_PASSWORD,
    DB: process.env.DB_NAME
}

const connection: Pool = mysql.createPool({
    host: DB_CONFIG.HOST,
    user: DB_CONFIG.USER,
    password: DB_CONFIG.PASSWORD,
    database: DB_CONFIG.DB,
});

export default connection