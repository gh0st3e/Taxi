import {CarStore} from "../store/car";
import {logger} from "../main";

const GetAllCars = async () => {
    logger.info("[GetAllCars] started")
    try {
        const cars = await CarStore.GetAllCars();
        logger.info(cars);
        logger.info("[GetAllCars] ended")
        return cars;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const InsertCar = async (model: string, lic_plate: string) => {
    logger.info("[InsertCar] started")
    try {
        const res = await CarStore.InsertCar(model, lic_plate)
        logger.info(res)
        logger.info("[InsertCar ended]")
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const UpdateCar = async (id: number, model: string, lic_plate: string) => {
    logger.info("[UpdateCar] started")
    try {
        const res = await CarStore.UpdateCar(id, model, lic_plate)
        logger.info(res)
        logger.info("[UpdateCar ended]")
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const DeleteCar = async (id: number) => {
    logger.info("[DeleteCar] started")
    try {
        const res = await CarStore.DeleteCar(id)
        logger.info(res)
        logger.info("[DeleteCar ended]")
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

export const CarService = {
    GetAllCars,
    InsertCar,
    UpdateCar,
    DeleteCar
}