DELIMITER //
CREATE PROCEDURE InsertCar(
    IN NewModel VARCHAR(100),
    IN NewLicPlate VARCHAR(20)
)
BEGIN
INSERT INTO car(model,lic_plate) VALUES(NewModel,NewLicPlate);
END //
DELIMITER