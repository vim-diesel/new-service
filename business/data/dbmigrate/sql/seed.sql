INSERT INTO users (user_id, name, email, roles, password_hash, department, enabled, date_created, date_updated) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'Admin Gopher', 'admin@example.com', '{ADMIN,USER}', '$2a$10$GAoWilm7dJBKx8JAV0JASOTo0P/EmQVDV3Nhnc7WWXbUF4v1g5byi', NULL, true, '2019-03-24 00:00:00', '2019-03-24 00:00:00'),
	('45b5fbd3-755f-4379-8f07-a58d4a30fa2f', 'User Gopher', 'user@example.com', '{USER}', '2a$10$f03pZLikjsmVEuw.D..kL.l2H/1z6RpjMOu/ySCp5.p1IfG3tE1aS', NULL, true, '2019-03-24 00:00:00', '2019-03-24 00:00:00')
ON CONFLICT DO NOTHING;