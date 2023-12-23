DELIMITER //
CREATE PROCEDURE SignIn(
    IN NewPhone VARCHAR(20),
    IN NewPassword VARCHAR(50)
)
BEGIN
SELECT id,name,phone,password,email FROM user WHERE user.phone=NewPhone AND password=NewPassword;
END //
DELIMITER