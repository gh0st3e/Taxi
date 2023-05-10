import express from "express";
import {CarService} from "../service/car";

const carRouter = express.Router()

carRouter.get("/api/cars", async (req, res) => {
    try {
        const cars = await CarService.GetAllCars()
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(cars))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({error: err.message}))
    }
})

carRouter.post("/api/cars", async (req, res) => {
    try {
        const {model, lic_plate} = req.body
        await CarService.InsertCar(model, lic_plate)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({error: err.message}))
    }
})

carRouter.put("/api/cars", async (req, res) => {
    try {
        const {id, model, lic_plate} = req.body
        console.log(id,model,lic_plate)
        await CarService.UpdateCar(id, model, lic_plate)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({error: err.message}))
    }
})

carRouter.delete("/api/cars", async (req, res) => {
    try {
        const {id} = req.body
        await CarService.DeleteCar(id)
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify("OK"))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({error: err.message}))
    }
})

export {carRouter}