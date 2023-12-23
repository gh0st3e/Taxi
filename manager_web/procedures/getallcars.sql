DELIMITER //
CREATE PROCEDURE GetAllCars()
BEGIN
SELECT car.id,car.model,car.lic_plate
FROM car;
END //
DELIMITER