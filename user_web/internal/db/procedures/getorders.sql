DELIMITER //
CREATE PROCEDURE GetOrders(
    IN NewID INT,
    IN NewState INT
)
BEGIN
SELECT orders.id, orders.FromCity, orders.ToCity, orders.Date, orders.Time, orders.Tickets, orders.State, driver.id, driver.name,driver.phone, car.id,car.model, car.lic_plate
FROM orders
         INNER JOIN driver ON orders.DriverID = driver.id
         INNER JOIN car ON orders.CarID = car.id
WHERE orders.UserID = NewID AND orders.State = NewState;
END //
DELIMITER