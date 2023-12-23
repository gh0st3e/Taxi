DELIMITER //
CREATE PROCEDURE GetDriverOrders(
    IN NewID INT,
    IN NewDate VARCHAR(60)
)
BEGIN
SELECT orders.id,orders.FromCity,orders.ToCity,orders.Date,orders.Time,orders.Tickets,
       user.id,user.name,user.phone,
       orders.UserID
FROM orders INNER JOIN user ON orders.UserID = user.id
WHERE orders.DriverID=NewID AND orders.Date=NewDate;
END //
DELIMITER