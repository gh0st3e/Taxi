DELIMITER //
CREATE PROCEDURE GetAllUsers()
BEGIN
SELECT user.id,user.name,user.phone,user.password,user.email,user.telegram
FROM user;
END //
DELIMITER