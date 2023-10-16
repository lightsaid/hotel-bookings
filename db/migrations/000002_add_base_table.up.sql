CREATE TABLE `room_types` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键", 
    room_type varchar(50) NOT NULL COMMENT "客房类型",
    room_label varchar(50) NOT NULL COMMENT "客房类型标签",
    UNIQUE unq_roomType(room_type)
) COMMENT="客房类型表";

CREATE TABLE `booking_status` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键", 
    status_type varchar(50) NOT NULL COMMENT "客房类型",
    status_label varchar(50) NOT NULL COMMENT "客房类型标签",
    UNIQUE unq_statusType(status_type)
) COMMENT="预定类型表";

