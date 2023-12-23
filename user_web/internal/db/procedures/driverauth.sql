DELIMITER //
CREATE PROCEDURE DriverAuth(
    IN NewPhone VARCHAR(20),
    IN NewPassword VARCHAR(50)
)
BEGIN
SELECT id,name,phone,password,work_exp FROM driver WHERE driver.phone=NewPhone AND driver.password=NewPassword;
END //
DELIMITER