import {DriverStore} from "../store/driver";
import {logger} from "../main";

const GetAllDrivers = async () => {
    logger.info("[GetAllDrivers] started")
    try {
        const drivers = await DriverStore.GetAllDrivers();
        logger.info(drivers);
        logger.info("[GetAllDrivers] ended")
        return drivers;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

export const DriverService = {
    GetAllDrivers
}