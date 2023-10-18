CREATE TABLE `rooms` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    hotel_id INT UNSIGNED NOT NULL COMMENT "外键,酒店表ID",
    room_number varchar(10) NOT NULL COMMENT "客房号",
    room_image varchar(255) NOT NULL COMMENT "客房图片",
    room_price INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "客房价格(单位分)",
    booking_status_id INT UNSIGNED NOT NULL COMMENT "外键,预定类型表ID",
    room_type_id INT UNSIGNED NOT NULL COMMENT "外键,客房类型表ID",
    room_capacity INT UNSIGNED NOT NULL DEFAULT 1 COMMENT "容纳人数",
    room_description varchar(255) NOT NULL DEFAULT "" COMMENT "客房描述",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    is_deleted tinyint(1) NOT NULL DEFAULT 0 COMMENT "是否删除, 0:正常, 1:删除",
    INDEX idx_hotelId(hotel_id),
    INDEX idx_bookingStatusId(booking_status_id),
    INDEX idx_roomTypeId(room_type_id),
    INDEX idx_isDeleted(is_deleted)
) COMMENT="客房表";

-- 每个酒店下，客房号是唯一
ALTER TABLE rooms 
ADD UNIQUE INDEX unq_roomNumber_hotelId(room_number, hotel_id);