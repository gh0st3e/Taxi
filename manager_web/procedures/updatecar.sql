DELIMITER //
CREATE PROCEDURE UpdateCar(
    IN NewID INT,
    IN NewModel VARCHAR(100),
    IN NewLicPlate VARCHAR(20)
)
BEGIN
UPDATE car SET model=NewModel, lic_plate=NewLicPlate WHERE car.id=NewID;
END //
DELIMITER