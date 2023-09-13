INSERT INTO users (user_id, full_name, first_name, last_name, email, enabled, created_at) VALUES
	('user_2V0HcZBFvnRXplo1ElzNIceGpZU', 'Admin Gopher', 'Admin', 'Gopher', 'admin+clerk_test@example.com', true, '2023-09-13 14:25:00')
ON CONFLICT DO NOTHING;

INSERT INTO retreats (
    retreat_id,
    title,
    body,
    user_id,
    status,
    cost,
    open_spots,
    start_date,
    end_date,
    created_at
) VALUES (
    '6a88a612-5244-11ee-be56-0242ac120002', -- Replace with a UUID for the retreat
    'Dummy Retreat',         -- Replace with the title of the retreat
    'This is a dummy retreat for testing.', -- Replace with the body/description
    'user_2V0HcZBFvnRXplo1ElzNIceGpZU', -- The user_id of the Admin Gopher
    'Pending',               -- Replace with the status of the retreat
    1000.00,                   -- Replace with the cost of the retreat
    10,                       -- Replace with the number of open spots
    '2023-09-20',             -- Replace with the start date of the retreat
    '2023-09-25',             -- Replace with the end date of the retreat
    NOW()                    -- Use the current timestamp for created_at
)
ON CONFLICT DO NOTHING;