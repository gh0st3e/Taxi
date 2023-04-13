import express from "express";
import mysql, {MysqlError, Pool} from 'mysql'
import fs from 'fs'

const app = express();
const PORT = process.env.PORT || 9000
const page = fs.readFileSync('index.html')

const connection: Pool = mysql.createPool({
    host: 'MYSQL8003.site4now.net',
    user: 'a97166_taxiexp',
    password: 'Denis_2003_db',
    database: 'db_a97166_taxiexp',
});

app.get("/", (req, res) => res.end(page))

app.get("/user/get", (req, res) => {
    connection.query("SELECT * FROM user", (err: MysqlError | null, result: any) => {
        if (err) {
            res.send(err)
            res.status(500).end()
            return
        }
        console.log(result)
        res.send(result)
        res.status(200).end()
    })
})

app.get("/driver/get", (req, res) => {
    connection.query("SELECT * FROM driver", (err: MysqlError | null, result: any) => {
        if (err) {
            res.send(err)
            res.status(500).end()
            return
        }
        console.log(result)
        res.send(result)
        res.status(200).end()
    })
})

app.listen(PORT, () => console.log(`Server is running on ${PORT}`))