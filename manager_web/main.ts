import {userRouter} from "./controller/user";
import {driverRouter} from "./controller/driver";
import {orderRouter} from "./controller/order";
import {carRouter} from "./controller/cars";
import {viewRouter} from "./controller/views"


import express from "express";
import bodyParser from "body-parser";
import log4js from 'log4js'

const PORT = process.env.PORT || 9000

log4js.configure('./log4js.json')
const logger = log4js.getLogger()

const app = express();
app.use(bodyParser.json())

app.use(userRouter)
app.use(driverRouter)
app.use(orderRouter)
app.use(carRouter)
app.use(viewRouter)

// app.get("/", (req, res) => res.end(page))
//
// app.get("/user/get", (req, res) => {
//     connection.query("SELECT * FROM user", (err: MysqlError | null, result: any) => {
//         if (err) {
//             res.send(err)
//             res.status(500).end()
//             return
//         }
//         console.log(result)
//         res.send(result)
//         res.status(200).end()
//     })
// })
//
// app.get("/driver/get", (req, res) => {
//     connection.query("SELECT * FROM driver", (err: MysqlError | null, result: any) => {
//         if (err) {
//             res.send(err)
//             res.status(500).end()
//             return
//         }
//         console.log(result)
//         res.send(result)
//         res.status(200).end()
//     })
// })

app.listen(PORT, () => console.log(`Server is running on ${PORT}`))

export {app, logger}