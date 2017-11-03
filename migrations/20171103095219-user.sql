
-- +migrate Up
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `picture` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime not NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
-- +migrate Down
