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

const InsertDriver = async (name: string, phone: string, password: string, work_exp: string) => {
    logger.info("[InsertDriver] started")
    try {
        const res =  await DriverStore.InsertDriver(name, phone, password, work_exp)
        logger.info(res)
        logger.info("[InsertDriver] ended")
        return res
    } catch (err) {
        logger.error(err)
        throw err
    }
}

const UpdateDriver = async (id: number, name: string, phone: string, password: string, work_exp: string) => {
    logger.info("[UpdateDriver] started")
    try {
        const res = await DriverStore.UpdateDriver(id, name, phone, password, work_exp)
        logger.info(res)
        logger.info("[UpdateDriver] ended")
        return res
    } catch (err) {
        logger.error(err)
        throw err
    }
}

const DeleteDriver = async (id: number) => {
    logger.info("[DeleteDriver] started")
    try {
        const res = await DriverStore.DeleteDriver(id)
        logger.info(res)
        logger.info("[DeleteDriver] ended")
        return res
    } catch (err) {
        logger.error(err)
        throw err
    }
}

export const DriverService = {
    GetAllDrivers,
    InsertDriver,
    UpdateDriver,
    DeleteDriver
}