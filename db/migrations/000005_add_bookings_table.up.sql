CREATE TABLE `bookings` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键", 
    user_id INT UNSIGNED NOT NULL COMMENT "外键,用户表ID",
    booking_from DATE NOT NULL COMMENT "入住日期",
    booking_to DATE NOT NULL COMMENT "退房日期",
    room_id INT UNSIGNED NOT NULL COMMENT "外键,客房表ID",
    members INT UNSIGNED NOT NULL COMMENT "入住人数",
    total_amount INT UNSIGNED NOT NULL COMMENT "总金额(单位分)",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    INDEX idx_userId(user_id),
    INDEX idx_roomId(room_id)
) COMMENT="预订表";
