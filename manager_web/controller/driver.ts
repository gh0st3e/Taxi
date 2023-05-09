import express from "express";
import {DriverService} from "../service/driver";

const driverRouter = express.Router()

driverRouter.get("/api/drivers", async (req, res) => {
    try {
        const drivers = await DriverService.GetAllDrivers()
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(drivers))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({ error: err.message }))
    }
})

driverRouter.post("/api/drivers", async (req, res) => {
    try {
        const {name, phone, password, work_exp} = req.body
        await DriverService.InsertDriver(name, phone, password, work_exp)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({ error: err.message }))
    }
})

driverRouter.put("/api/drivers", async (req,res)=>{
    try {
        const {id, name, phone, password, work_exp} = req.body
        console.log(id, name, phone, password, work_exp)
        await DriverService.UpdateDriver(id, name, phone, password, work_exp)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({ error: err.message }))
    }
})

driverRouter.delete("/api/drivers", async (req,res)=>{
    try {
        const {id} = req.body
        console.log(id)
        await DriverService.DeleteDriver(id)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({ error: err.message }))
    }
})

export {driverRouter}