CREATE TABLE `sessions` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键",
    token_id VARCHAR(64) NOT NULL COMMENT "Token的唯一标识uuid",
    user_id INT UNSIGNED NOT NULL COMMENT "外键,用户表ID",
    refresh_token VARCHAR(2048) NOT NULL COMMENT "用户的refresh token",
    user_agent VARCHAR(1024) NOT NULL COMMENT "登录设备的信息",
    client_ip VARCHAR(20) NOT NULL COMMENT "登录时IP",
    expires_at TIMESTAMP NOT NULL COMMENT "session或refresh token 过期时间",
    created_at TIMESTAMP NOT NULL DEFAULT NOW() COMMENT "创建时间",
    login_type SMALLINT NOT NULL DEFAULT 0 COMMENT "登录类型(0:密码登录, 1:短信验证码)",
    UNIQUE unq_token_id(token_id),
    INDEX idx_userId(user_id)
) COMMENT = "会话表";