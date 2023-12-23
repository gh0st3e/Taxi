DELIMITER //
CREATE PROCEDURE InsertDriver(
    IN NewName VARCHAR(50),
    IN NewPhone VARCHAR(20),
    IN NewPassword VARCHAR(50),
    IN NewWorkExp INT
)
BEGIN
INSERT INTO driver(name,phone,password,work_exp) VALUES(NewName,NewPhone,NewPassword,NewWorkExp);
END //
DELIMITER