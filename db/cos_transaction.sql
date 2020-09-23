-- phpMyAdmin SQL Dump
-- version 4.4.15.10
-- https://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jun 11, 2020 at 12:26 AM
-- Server version: 5.6.48-log
-- PHP Version: 5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `testblockchain`
--

-- --------------------------------------------------------

--
-- Table structure for table `cos_address`
--

CREATE TABLE IF NOT EXISTS `block_transaction` (
  `id` int(11) NOT NULL,
  `fid` int(11) NOT NULL,
  `coin` varchar(300) NOT NULL,
  `from` varchar(300) NOT NULL,
  `to` varchar(300) NOT NULL,
  `amount` varchar(300) NOT NULL,
  `hash` varchar(300) NOT NULL,
  `block_height` int(11) NOT NULL,
  `time` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cos_address`
--
ALTER TABLE `block_transaction`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `block_transaction`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


--
-- AUTO_INCREMENT for dumped tables
--
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
