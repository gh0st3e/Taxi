DELIMITER //
CREATE PROCEDURE GetAllOrdersDateAsc(IN NewState INT)
BEGIN
SELECT orders.id, orders.UserID, orders.DriverID, orders.CarID, orders.FromCity, orders.ToCity, orders.Date, orders.Time, orders.Tickets, orders.State
FROM orders
WHERE orders.State = NewState
ORDER BY orders.Date ASC;
END //
DELIMITER