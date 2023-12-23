DELIMITER //
CREATE PROCEDURE RetrieveById(
    IN NewID INT
)
BEGIN
SELECT id,name,phone,password,email,telegram FROM user WHERE user.id=NewID;
END //
DELIMITER