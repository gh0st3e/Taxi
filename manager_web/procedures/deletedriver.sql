DELIMITER //
CREATE PROCEDURE DeleteDriver(
    IN NewID INT
)
BEGIN
DELETE FROM driver WHERE driver.id=NewID;
END //
DELIMITER