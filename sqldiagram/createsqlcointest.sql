-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema sqlcointest
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema sqlcointest
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `sqlcointest` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `sqlcointest` ;

-- -----------------------------------------------------
-- Table `sqlcointest`.`blocks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcointest`.`blocks` (
  `height` INT(11) NOT NULL AUTO_INCREMENT,
  `hash` VARCHAR(64) NULL,
  `prevHash` VARCHAR(64) NULL,
  `merkle` VARCHAR(64) NULL,
  `timestamp` BIGINT(20) NULL DEFAULT NULL,
  PRIMARY KEY (`height`),
  UNIQUE INDEX `height_UNIQUE` (`height` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 0
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcointest`.`inputs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcointest`.`inputs` (
  `inID` INT(11) NOT NULL AUTO_INCREMENT,
  `txHash` VARCHAR(64) NULL,
  `prevTxHash` VARCHAR(64) NULL,
  `prevTxIndex` INT NULL,
  PRIMARY KEY (`inID`),
  UNIQUE INDEX `inID_UNIQUE` (`inID` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 0
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcointest`.`outputs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcointest`.`outputs` (
  `outID` INT(11) NOT NULL AUTO_INCREMENT,
  `txHash` VARCHAR(64) NULL,
  `txIndex` INT(11) NULL,
  `amount` BIGINT(20) NULL DEFAULT NULL,
  `address` VARCHAR(59) NULL,
  `used` BIGINT(20) NULL DEFAULT NULL,
  PRIMARY KEY (`outID`),
  UNIQUE INDEX `outID_UNIQUE` (`outID` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 0
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


-- -----------------------------------------------------
-- Table `sqlcointest`.`txs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `sqlcointest`.`txs` (
  `txID` INT(11) NOT NULL AUTO_INCREMENT,
  `hash` VARCHAR(64) NULL,
  `blockHeight` INT(11) NULL,
  PRIMARY KEY (`txID`),
  UNIQUE INDEX `txID_UNIQUE` (`txID` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 0
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
