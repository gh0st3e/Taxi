import fs from 'fs'

import express from 'express';

const userList = fs.readFileSync('views/userList.html')
const driverList = fs.readFileSync('views/driverList.html')
const orderList = fs.readFileSync("views/orderList.html")
const carList = fs.readFileSync("views/carList.html")
const updateOrder = fs.readFileSync("views/updateOrder.html")

const viewRouter = express.Router()

viewRouter.get("/view/users", (req, res) => {
    res.end(userList)
})

viewRouter.get("/view/drivers", (req, res) => {
    res.end(driverList)
})

viewRouter.get("/view/orders", (req, res) => {
    res.end(orderList)
})

viewRouter.get("/view/cars", (req, res) => {
    res.end(carList)
})

viewRouter.get("/view/orderupd", (req, res) => {
    res.end(updateOrder)
})

export {viewRouter}