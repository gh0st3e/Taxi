DELIMITER
//
CREATE PROCEDURE SignUp(
    IN NewName VARCHAR(50),
    IN NewPhone VARCHAR(20),
    IN NewPassword VARCHAR(50),
    IN NewEmail VARCHAR(50)
)
BEGIN
INSERT INTO USER(name,phone,password,email)
VALUES(NewName,NewPhone,NewPassword,NewEmail);
END //
DELIMITER