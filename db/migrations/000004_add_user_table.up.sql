CREATE TABLE `users` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    phone_number varchar(11) NOT NULL COMMENT "用户手机号码",
    hashed_password varchar(60) NOT NULL COMMENT "哈希密码",
    username varchar(60) NOT NULL COMMENT "用户名称",
    avatar varchar(255) NOT NULL COMMENT "用户头像",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    is_deleted tinyint(1) NOT NULL DEFAULT 0 COMMENT "是否删除, 0:正常, 1:删除",
    INDEX idx_phoneNumber(phone_number),
    INDEX idx_isDeleted(is_deleted),
    UNIQUE unq_phoneNumber(phone_number)
) COMMENT="用户表";