CREATE TABLE `hotels` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    title varchar(200) NOT NULL COMMENT "酒店名称",
    code varchar(8) NOT NULL COMMENT "酒店编码",
    address varchar(255) NOT NULL COMMENT "酒店地址",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "更新时间",
    is_deleted tinyint(1) NOT NULL DEFAULT 0 COMMENT "是否删除, 0:正常, 1:删除",
    INDEX idx_hotelCode(title),
    INDEX idx_isDeleted(is_deleted),
    UNIQUE unq_hotelCode(code)
) COMMENT="酒店表";
