import {OrderStore} from "../store/order";
import {logger} from "../main";

const GetAllOrders = async (state: string) => {
    logger.info("[GetAllOrders] started")
    try {
        let intState = Number(state)
        const orders = await OrderStore.GetAllOrders(intState);
        logger.info(orders);
        logger.info("[GetAllOrders] ended")
        return orders;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const UpdateOrder = async (orderID: number, driverID: number, carID: number) => {
    logger.info("[UpdateOrder] started")
    try {
        const res = await OrderStore.UpdateOrder(orderID, driverID, carID);
        logger.info(res);
        logger.info("[UpdateOrder] ended")
        return res;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const GetAllOrdersDateA = async (state: string) => {
    logger.info("[GetAllOrdersDateA] started")
    try {
        let intState = Number(state)
        const orders = await OrderStore.GetAllOrdersDateA(intState);
        logger.info(orders);
        logger.info("[GetAllOrdersDateA] ended")
        return orders;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

const GetAllOrdersDateD = async (state: string) => {
    logger.info("[GetAllOrdersDateD] started")
    try {
        let intState = Number(state)
        const orders = await OrderStore.GetAllOrdersDateD(intState);
        logger.info(orders);
        logger.info("[GetAllOrdersDateD] ended")
        return orders;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

export const OrderService = {
    GetAllOrders,
    UpdateOrder,
    GetAllOrdersDateA,
    GetAllOrdersDateD
}