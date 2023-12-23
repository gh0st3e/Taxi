DELIMITER //
CREATE PROCEDURE UpdateUser(
    IN NewName VARCHAR(50),
    IN NewPassword VARCHAR(50),
    IN NewEmail VARCHAR(50),
    IN NewID INT
)
BEGIN
UPDATE user SET user.name=NewName, user.password=NewPassword, user.email=NewEmail WHERE user.id=NewID;
END //
DELIMITER