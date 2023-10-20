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

INSERT INTO hotels (title, code, address) 
VALUES("蓝色幻想酒店", "LSMHJD", "广州祈福缤纷汇鸿福路77号"),
      ("风景如画旅馆", "FJRHLG", "杭州市西湖区龙井路99号"),
      ("科幻之城宾馆", "KHZCBG", "深圳市盐田区盐梅路88号");

INSERT INTO rooms (
    hotel_id, 
    room_number, 
    room_image,
    room_price,
    booking_status_id,
    room_type_id,
    room_capacity
)
VALUES(1, "201", "http://", 19800, 1, 1, 2),
      (1, "202", "http://", 19800, 1, 1, 2),
      (1, "203", "http://", 29900, 1, 2, 4),
      (1, "204", "http://", 29900, 1, 2, 4),
      (1, "205", "http://", 59900, 1, 3, 4),
      (1, "301", "http://", 19800, 1, 1, 2),
      (1, "302", "http://", 29900, 1, 2, 4),
      (1, "303", "http://", 59900, 1, 3, 4),
      (1, "304", "http://", 19800, 1, 1, 2),
      (1, "305", "http://", 29900, 1, 2, 4),
      (2, "A10", "http://", 12800, 1, 1, 2),
      (2, "A21", "http://", 12800, 1, 1, 2),
      (2, "A31", "http://", 12800, 1, 1, 2),
      (2, "A33", "http://", 25600, 1, 2, 4),
      (2, "B90", "http://", 25600, 1, 2, 4),
      (2, "B91", "http://", 25600, 1, 2, 4),
      (2, "B92", "http://", 77700, 1, 3, 4),
      (2, "B93", "http://", 77700, 1, 3, 4),
      (2, "B94", "http://", 77700, 1, 3, 4),
      (2, "B95", "http://", 77700, 1, 3, 4),
      (3, "801", "http://", 66660, 1, 2, 4),
      (3, "802", "http://", 66660, 1, 2, 4),
      (3, "803", "http://", 66660, 1, 2, 4),
      (3, "804", "http://", 88880, 1, 3, 4),
      (3, "805", "http://", 88880, 1, 3, 4),
      (3, "806", "http://", 88880, 1, 3, 4),
      (3, "807", "http://", 88880, 1, 3, 4),
      (3, "809", "http://", 88880, 1, 3, 4),
      (3, "810", "http://", 33330, 1, 1, 2),
      (3, "811", "http://", 33330, 1, 1, 2),
      (3, "812", "http://", 33330, 1, 1, 2);

      
