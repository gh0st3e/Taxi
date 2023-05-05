import express from "express";
import {DriverService} from "../service/driver";

const driverRouter = express.Router()

driverRouter.get("/api/drivers", async (req, res) => {
    try {
        const drivers = await DriverService.GetAllDrivers()
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(drivers))
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

export {driverRouter}