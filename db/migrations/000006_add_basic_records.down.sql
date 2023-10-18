DELETE FROM users WHERE phone_number in ("13876543210", "15876543210");
DELETE FROM room_types WHERE room_type in ("Single","StandardTwin","DeluxeTwin");
DELETE FROM booking_status WHERE status_type in ("Free","Reserved","","Occupied");
