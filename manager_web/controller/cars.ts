import express from "express";
import {CarService} from "../service/car";

const carRouter = express.Router()

carRouter.get("/api/cars", async (req, res) => {
    try {
        const cars = await CarService.GetAllCars()
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(cars))
    } catch (err) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(err))
    }
})

export {carRouter}