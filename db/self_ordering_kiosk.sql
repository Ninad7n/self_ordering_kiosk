-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 27, 2024 at 10:53 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `self_ordering_kiosk`
--

-- --------------------------------------------------------

--
-- Table structure for table `food`
--

CREATE TABLE `food` (
  `ID` int(11) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `IS_VEG` int(255) NOT NULL,
  `FULL_PRICE` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `food`
--

INSERT INTO `food` (`ID`, `NAME`, `IS_VEG`, `FULL_PRICE`) VALUES
(1, 'Veg Momos', 1, '110'),
(2, 'Chicken Momos', 0, '130'),
(4, 'Veg Fried Momos', 1, '120'),
(5, 'Checken Fried Momos', 0, '150');

-- --------------------------------------------------------

--
-- Table structure for table `invoice`
--

CREATE TABLE `invoice` (
  `ID` int(11) NOT NULL,
  `TABLE_NO` int(11) NOT NULL,
  `TAX` varchar(225) NOT NULL,
  `AMOUNT` varchar(255) NOT NULL,
  `STATUS` enum('paid','cancelled','created','') NOT NULL DEFAULT 'created',
  `CREATED_DATE` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `invoice`
--

INSERT INTO `invoice` (`ID`, `TABLE_NO`, `TAX`, `AMOUNT`, `STATUS`, `CREATED_DATE`) VALUES
(16, 1, '49', '539', 'paid', '2024-03-27 07:09:31');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `ID` int(11) NOT NULL,
  `INVOICE_ID` int(11) NOT NULL,
  `TABLE_NO` int(255) NOT NULL,
  `CUST_MOBILE` varchar(225) NOT NULL,
  `FOOD_NAME` varchar(255) NOT NULL,
  `PRICE` varchar(225) NOT NULL,
  `QUANTITY` int(255) NOT NULL,
  `PAYMENT` enum('paid','unpaid','in_cart','placed') NOT NULL DEFAULT 'unpaid',
  `CREATED_DATE` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`ID`, `INVOICE_ID`, `TABLE_NO`, `CUST_MOBILE`, `FOOD_NAME`, `PRICE`, `QUANTITY`, `PAYMENT`, `CREATED_DATE`) VALUES
(5, 16, 1, '678678678', 'Veg Fried Momos', '120', 1, 'paid', '2024-03-27 12:58:59'),
(6, 16, 1, '678678678', 'Veg Fried Momos', '120', 2, 'paid', '2024-03-27 12:58:59'),
(7, 0, 2, '678678678', 'Veg Fried Momos', '120', 1, 'paid', '2024-03-27 12:49:10'),
(8, 16, 1, '678678678', 'Chicken Momos', '130', 1, 'paid', '2024-03-27 12:58:59');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `food`
--
ALTER TABLE `food`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `invoice`
--
ALTER TABLE `invoice`
  ADD PRIMARY KEY (`ID`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `food`
--
ALTER TABLE `food`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `invoice`
--
ALTER TABLE `invoice`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
