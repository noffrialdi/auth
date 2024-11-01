CREATE TABLE `user` (
  `id` uuid NOT NULL PRIMARY KEY,
  `first_name` varchar(100) NOT NULL,
  `last_name` varchar(100),
  `phone_number` varchar(50),
  `username` varchar(50),
  `address` varchar(255),
  `password` varchar(100) NOT NULL,
  `user_id_created` bigint(20) NOT NULL,
  `user_id_updated` bigint(20) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0,
  `created_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `updated_time` timestamp(6) NOT NULL DEFAULT current_timestamp(6),
  `deleted_time` timestamp(6) NULL DEFAULT NULL
);
