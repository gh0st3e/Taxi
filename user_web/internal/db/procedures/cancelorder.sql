DELIMITER //
CREATE PROCEDURE CancelOrder(
    IN NewID INT
)
BEGIN
UPDATE orders SET orders.State=-1 WHERE orders.id=NewID;
END //
DELIMITER