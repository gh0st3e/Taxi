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

export const CarService = {
    GetAllCars
}