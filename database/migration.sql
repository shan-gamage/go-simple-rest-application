CREATE TABLE exchange_rate (
   id int NOT NULL AUTO_INCREMENT,
   assest_type varchar(20) NOT NULL,
   exchange_type varchar(20) NOT NULL,
   rate DECIMAL(20,12),
   created_at TIMESTAMP NOT NULL DEFAULT (UTC_TIMESTAMP),
   PRIMARY KEY (id)
);

INSERT INTO `exchange_rate` (`assest_type`, `exchange_type`, `rate`) VALUES 
('BTC','USD', '20664.391743189553'),
('BTC','USD', '20665.491743189554'),
('BTC','USD', '20666.591743189555'),
('BTC','USD', '20667.691743189556'),
('BTC','USD', '20668.791743189557'),
('BTC','USD', '20669.891743189558')