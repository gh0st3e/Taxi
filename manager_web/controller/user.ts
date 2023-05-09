import express from "express";
import {UserService} from "../service/user";

const userRouter = express.Router()

userRouter.get("/api/users", async (req, res) => {
    try {
        const users = await UserService.GetAllUsers()
        res.statusCode = 200
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify(users))
    } catch (err: any) {
        res.statusCode = 500
        res.setHeader("Content-Type", "application/json")
        res.end(JSON.stringify({ error: err.message }))
    }
})

export {userRouter}