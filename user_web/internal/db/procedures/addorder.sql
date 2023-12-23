DELIMITER //
CREATE PROCEDURE AddOrder(
    IN NewID INT,
    IN NewFrom VARCHAR(50),
    IN NewTo VARCHAR(50),
    IN NewDate VARCHAR(60),
    IN NewTime VARCHAR(20),
    IN NewTickets INT
)
BEGIN
INSERT INTO orders(UserID,FromCity,ToCity,Date,Time,Tickets) VALUES(NewID,NewFrom,NewTo,NewDate,NewTime,NewTickets);
END //
DELIMITER