CREATE TABLE `inventory`.`transaction` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `quantity` INT NOT NULL,
  `movement` INT NOT NULL DEFAULT 1,
  `user_creator` INT NOT NULL,
  `user_confirm` INT NULL,
  `product_id` INT NOT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NULL DEFAULT,
  PRIMARY KEY (`id`));

CREATE TABLE `inventory`.`product` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL,
  PRIMARY KEY (`id`));


CREATE TABLE `inventory`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL,
  PRIMARY KEY (`id`));
