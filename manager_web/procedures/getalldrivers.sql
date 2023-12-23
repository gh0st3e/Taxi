DELIMITER //
CREATE PROCEDURE GetAllDrivers()
BEGIN
SELECT driver.id,driver.name,driver.phone,driver.password,driver.work_exp
FROM driver;
END //
DELIMITER