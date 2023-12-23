DELIMITER //
CREATE PROCEDURE UpdateDriver(
    IN NewID INT,
    IN NewName VARCHAR(50),
    IN NewPhone VARCHAR(20),
    IN NewPassword VARCHAR(50),
    IN NewWorkExp INT
)
BEGIN
UPDATE driver SET name=NewName, phone=NewPhone, password=NewPassword, work_exp=NewWorkExp WHERE driver.id=NewID;
END //
DELIMITER