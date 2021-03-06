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

CREATE TABLE IF NOT EXISTS `dxp_address` (
  `id` int(11) NOT NULL,
  `name` varchar(300) NOT NULL,
  `addressdid` varchar(300) NOT NULL,
  `addressdx` varchar(300) NOT NULL,
  `pubkey` varchar(300) NOT NULL,
  `mnemonic` varchar(400) NOT NULL,
  `raw` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cos_address`
--
ALTER TABLE `dxp_address`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `dxp_address`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for dumped tables
--
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
