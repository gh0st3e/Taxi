DELIMITER //
CREATE PROCEDURE UpdateManagerOrder(
    IN NewOrderID INT,
    IN NewDriverID INT,
    IN NewCarID INT,
    IN NewDate VARCHAR(60),
    IN NewTime VARCHAR(20),
    IN NewState INT

)
BEGIN
UPDATE orders SET orders.DriverID=NewDriverID, orders.CarID=NewCarID, orders.Date=NewDate, orders.Time=NewTime, orders.State=NewState
WHERE orders.id = NewOrderID;
END //
DELIMITER