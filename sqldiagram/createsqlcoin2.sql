-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema sqlcoin
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema sqlcoin
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sqlcoin` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `sqlcoin` ;

-- -----------------------------------------------------
-- Table `sqlcoin`.`blocks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcoin`.`blocks` (
  `height` INT(11) NOT NULL AUTO_INCREMENT,
  `hash` VARCHAR(64) NOT NULL,
  `prevHash` VARCHAR(64) NOT NULL,
  `merkle` VARCHAR(64) NOT NULL,
  `timestamp` BIGINT(20) NULL DEFAULT NULL,
  PRIMARY KEY (`height`),
  UNIQUE INDEX `height_UNIQUE` (`height` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcoin`.`inputs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcoin`.`inputs` (
  `inID` INT(11) NOT NULL AUTO_INCREMENT,
  `txID` INT(11) NOT NULL,
  `outID` INT(11) NOT NULL,
  PRIMARY KEY (`inID`),
  UNIQUE INDEX `inID_UNIQUE` (`inID` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcoin`.`outputs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcoin`.`outputs` (
  `outID` INT(11) NOT NULL AUTO_INCREMENT,
  `txID` INT(11) NOT NULL,
  `index` INT(11) NOT NULL,
  `amount` BIGINT(20) NULL DEFAULT NULL,
  `address` VARCHAR(59) NOT NULL,
  `used` BIGINT(20) NULL DEFAULT NULL,
  PRIMARY KEY (`outID`),
  UNIQUE INDEX `outID_UNIQUE` (`outID` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcoin`.`txs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcoin`.`txs` (
  `txID` INT(11) NOT NULL AUTO_INCREMENT,
  `hash` VARCHAR(64) NOT NULL,
  `blockHeight` INT(11) NOT NULL,
  PRIMARY KEY (`txID`),
  UNIQUE INDEX `txID_UNIQUE` (`txID` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
