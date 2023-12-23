DELIMITER //
CREATE PROCEDURE UpdatePassengerStatus(
    IN NewID INT,
    IN NewState INT
)
BEGIN
UPDATE orders SET orders.State=NewState WHERE orders.id=NewID;
END //
DELIMITER