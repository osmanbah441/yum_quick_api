-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema yum_quick_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema yum_quick_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `yum_quick_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `yum_quick_db` ;

-- -----------------------------------------------------
-- Table `yum_quick_db`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password_hash` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`carts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`carts` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `userId` INT NOT NULL,
  `deliveryCost` DOUBLE NOT NULL,
  `quantity` INT NOT NULL,
  `subtotal` DOUBLE NOT NULL,
  `total` DOUBLE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `userId` (`userId` ASC) VISIBLE,
  CONSTRAINT `carts_ibfk_1`
    FOREIGN KEY (`userId`)
    REFERENCES `yum_quick_db`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `price` DOUBLE NOT NULL,
  `description` TEXT NOT NULL,
  `imageUrl` VARCHAR(255) NOT NULL DEFAULT 'https://via.placeholder.com/150',
  `category` ENUM('shawarma', 'pizza', 'burger', 'yogurt') NOT NULL,
  `averageRating` DOUBLE NOT NULL DEFAULT '0',
  `inventory` INT NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`cart_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`cart_items` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `cartId` INT NOT NULL,
  `productId` INT NOT NULL,
  `quantity` INT NOT NULL,
  `totalPrice` DOUBLE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `cartId` (`cartId` ASC) VISIBLE,
  INDEX `productId` (`productId` ASC) VISIBLE,
  CONSTRAINT `cart_items_ibfk_1`
    FOREIGN KEY (`cartId`)
    REFERENCES `yum_quick_db`.`carts` (`id`),
  CONSTRAINT `cart_items_ibfk_2`
    FOREIGN KEY (`productId`)
    REFERENCES `yum_quick_db`.`products` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`favorites`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`favorites` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `userId` INT NOT NULL,
  `productId` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `userId` (`userId` ASC) VISIBLE,
  INDEX `productId` (`productId` ASC) VISIBLE,
  CONSTRAINT `favorites_ibfk_1`
    FOREIGN KEY (`userId`)
    REFERENCES `yum_quick_db`.`users` (`id`),
  CONSTRAINT `favorites_ibfk_2`
    FOREIGN KEY (`productId`)
    REFERENCES `yum_quick_db`.`products` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`orders` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `userId` INT NOT NULL,
  `deliveryDate` DATETIME NOT NULL,
  `totalPrice` DOUBLE NOT NULL,
  `quantity` INT NOT NULL,
  `status` ENUM('pending', 'ongoing', 'completed', 'cancelled') NOT NULL DEFAULT 'pending',
  PRIMARY KEY (`id`),
  INDEX `userId` (`userId` ASC) VISIBLE,
  CONSTRAINT `orders_ibfk_1`
    FOREIGN KEY (`userId`)
    REFERENCES `yum_quick_db`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `yum_quick_db`.`order_items`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `yum_quick_db`.`order_items` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `orderId` INT NOT NULL,
  `productId` INT NOT NULL,
  `quantity` INT NOT NULL,
  `totalPrice` DOUBLE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `orderId` (`orderId` ASC) VISIBLE,
  INDEX `productId` (`productId` ASC) VISIBLE,
  CONSTRAINT `order_items_ibfk_1`
    FOREIGN KEY (`orderId`)
    REFERENCES `yum_quick_db`.`orders` (`id`),
  CONSTRAINT `order_items_ibfk_2`
    FOREIGN KEY (`productId`)
    REFERENCES `yum_quick_db`.`products` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;

 
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
