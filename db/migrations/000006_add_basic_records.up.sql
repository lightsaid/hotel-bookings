INSERT INTO 
    users(phone_number, hashed_password, username, avatar)
VALUES(
    "13876543210", 
    "$2a$10$g37hj82azyUOVZPyPzkoHOfdMTw1GKLKrXAj/SQYAM2lU5es8K00W", -- 123456
    "lazy",
    "/static/images/default_avatar.png"
),(
    "15876543210",
    "$2a$10$g37hj82azyUOVZPyPzkoHOfdMTw1GKLKrXAj/SQYAM2lU5es8K00W", -- 123456
    "xqq",
    "/static/images/default_avatar.png"
);


INSERT INTO room_types (room_type, room_label) 
VALUES ('Single', '单人客房'),
        ('StandardTwin', '标准双床房'),
        ('DeluxeTwin', '豪华双床房');


INSERT INTO booking_status (status_type, status_label) 
VALUES ('Free','空闲'),
        ('Reserved', '保留'),
        ('Occupied', '占用');
