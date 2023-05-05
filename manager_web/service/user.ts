import {UserStore} from "../store/user";
import {logger} from "../main";

const GetAllUsers = async () => {
    logger.info("[GetAllUsers] started")
    try {
        const users = await UserStore.GetAllUsers();
        logger.info(users);
        logger.info("[GetAllUsers] ended")
        return users;
    } catch (err) {
        logger.error(err);
        throw err;
    }
}

export const UserService = {
    GetAllUsers
}