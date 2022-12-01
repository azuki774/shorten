-- +migrate Up
CREATE TABLE `URL_table` (
  `id` int NOT NULL AUTO_INCREMENT,
  `source` varchar(255) NOT NULL,
  `target` varchar(255) NOT NULL,
  `expired_at` timestamp NOT NULL,
  `created_at` timestamp NOT NULL default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp,
  PRIMARY KEY (`id`),
  INDEX `idx_m1` (`source`, `expired_at`)
);

-- +migrate Down
DROP TABLE `URL_table`;
