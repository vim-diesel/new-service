-- Version: 1.01
-- Description: Create table users
CREATE TABLE "users" (
  "user_id" text PRIMARY KEY UNIQUE,
  "full_name" text,
  "first_name" text,
  "last_name" text,
  "email" text UNIQUE,
  "enabled" boolean,
  "created_at" timestamp
);

-- Version: 1.02
-- Description: Create table products
CREATE TABLE "retreats" (
  "retreat_id" uuid PRIMARY KEY,
  "title" text,
  "body" text,
  "user_id" text,
  "status" text,
  "cost" decimal(10,2),
  "open_spots" integer,
  "start_date" date,
  "end_date" date,
  "created_at" timestamp,

	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Version: 1.03
-- Description: Add user_summary view.
CREATE OR REPLACE VIEW user_summary AS
SELECT
    u.user_id   AS user_id,
	u.full_name      AS user_name,
    COUNT(r.*)  AS total_count
FROM
    users AS u
JOIN
    retreats AS r ON r.user_id = u.user_id
GROUP BY
    u.user_id