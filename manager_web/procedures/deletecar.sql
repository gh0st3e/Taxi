DELIMITER //
CREATE PROCEDURE DeleteCar(
    IN NewID INT
)
BEGIN
DELETE FROM car WHERE car.id=NewID;
END //
DELIMITER