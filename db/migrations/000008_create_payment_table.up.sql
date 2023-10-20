CREATE TABLE `payments` (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT "主键", 
    booking_id INT UNSIGNED NOT NULL COMMENT "外键,预定表ID",
    payment_type INT NOT NULL DEFAULT 0 COMMENT "支付类型(-1:取消支付,0:未支付,1:现金,2:微信,3:支付宝,4:银联,5:其他)",
    payment_amount INT UNSIGNED NOT NULL COMMENT "付款金额(单位分)",
    is_deleted tinyint(1) NOT NULL DEFAULT 0 COMMENT "是否删除, 0:正常, 1:删除",
    INDEX idx_bookingId(booking_id),
    INDEX idx_paymentType(payment_type)
) COMMENT="付款表";

