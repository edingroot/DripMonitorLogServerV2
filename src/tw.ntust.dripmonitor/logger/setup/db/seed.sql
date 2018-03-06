-- phpMyAdmin SQL Dump
-- version 4.6.6
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Mar 06, 2018 at 05:13 PM
-- Server version: 5.5.56-MariaDB
-- PHP Version: 5.6.32

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `dripmonitor_logging`
--

-- --------------------------------------------------------

--
-- Table structure for table `event_log`
--

CREATE TABLE `event_log` (
  `sn` int(11) NOT NULL,
  `event_code` tinyint(4) NOT NULL DEFAULT '0',
  `message` text COLLATE utf8_unicode_ci,
  `mac_adapter` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mac_drip` varchar(16) COLLATE utf8_unicode_ci DEFAULT NULL,
  `src_ip` varchar(39) COLLATE utf8_unicode_ci DEFAULT NULL,
  `src_port` int(11) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tcp_log_stream`
--

CREATE TABLE `tcp_log_stream` (
  `sn` int(10) UNSIGNED NOT NULL,
  `message` text COLLATE utf8_unicode_ci,
  `src_ip` varchar(39) COLLATE utf8_unicode_ci DEFAULT NULL,
  `src_port` int(10) UNSIGNED DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `event_log`
--
ALTER TABLE `event_log`
  ADD PRIMARY KEY (`sn`),
  ADD KEY `event_code` (`event_code`),
  ADD KEY `created_at` (`created_at`);

--
-- Indexes for table `tcp_log_stream`
--
ALTER TABLE `tcp_log_stream`
  ADD PRIMARY KEY (`sn`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `event_log`
--
ALTER TABLE `event_log`
  MODIFY `sn` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34656;
--
-- AUTO_INCREMENT for table `tcp_log_stream`
--
ALTER TABLE `tcp_log_stream`
  MODIFY `sn` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11149;