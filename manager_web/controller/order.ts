import express from "express";
import {OrderService} from "../service/order";

const orderRouter = express.Router()

orderRouter.get("/api/orders/:state", async (req, res) => {
    try {
        let state = req.params.state
        const users = await OrderService.GetAllOrders(state)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(users))
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

orderRouter.get("/api/ordersa/:state", async (req, res) => {
    try {
        let state = req.params.state
        const users = await OrderService.GetAllOrdersDateA(state)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(users))
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

orderRouter.get("/api/ordersd/:state", async (req, res) => {
    try {
        let state = req.params.state
        const users = await OrderService.GetAllOrdersDateD(state)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(users))
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

orderRouter.put("/api/orders", async (req, res) => {
    try {
        const {orderID, driverID, carID, date, time} = req.body;
        await OrderService.UpdateOrder(orderID, driverID, carID, date, time); // вызываем функцию обновления заказа
        res.statusCode = 200;
        res.setHeader('Content-Type', 'application/json');
        res.end(JSON.stringify({message: 'Order updated successfully'}));
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

export {orderRouter}