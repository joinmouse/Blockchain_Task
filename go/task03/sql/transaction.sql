-- 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
-- （包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）
-- 要求 ：
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
-- 如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

-- =============================================
-- 创建数据库表结构
-- =============================================

-- 创建账户表：用于存储用户账户信息
CREATE TABLE accounts (
    id INT PRIMARY KEY AUTO_INCREMENT,    -- 账户ID，主键，自动递增
    balance DECIMAL(10,2) NOT NULL        -- 账户余额，精确到小数点后两位，不允许为空
);

-- 创建交易记录表：用于记录所有转账交易历史
CREATE TABLE transactions (
    id INT PRIMARY KEY AUTO_INCREMENT,    -- 交易ID，主键，自动递增
    from_account_id INT NOT NULL,         -- 转出账户ID，不允许为空
    to_account_id INT NOT NULL,           -- 转入账户ID，不允许为空
    amount DECIMAL(10,2) NOT NULL,        -- 转账金额，精确到小数点后两位，不允许为空
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 交易创建时间，默认为当前时间
    FOREIGN KEY (from_account_id) REFERENCES accounts(id),  -- 外键约束：确保转出账户存在
    FOREIGN KEY (to_account_id) REFERENCES accounts(id)     -- 外键约束：确保转入账户存在
);

-- =============================================
-- 插入测试数据
-- =============================================

-- 创建两个测试账户
INSERT INTO accounts (balance) VALUES (500.00);  -- 账户A：初始余额500元
INSERT INTO accounts (balance) VALUES (200.00);  -- 账户B：初始余额200元

-- =============================================
-- 创建转账存储过程
-- =============================================

DELIMITER //
CREATE PROCEDURE transfer_money(
    IN from_account INT,                  -- 输入参数：转出账户ID
    IN to_account INT,                    -- 输入参数：转入账户ID
    IN transfer_amount DECIMAL(10,2)      -- 输入参数：转账金额
)
BEGIN
    -- 声明变量：用于存储当前账户余额
    DECLARE current_balance DECIMAL(10,2);
    
    -- 声明异常处理器：捕获任何SQL错误
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;  -- 发生错误时回滚事务
        SELECT '交易失败：系统发生错误' AS message;
    END;

    -- 开始事务
    START TRANSACTION;
    
    -- 检查转出账户的余额并锁定该记录
    SELECT balance INTO current_balance 
    FROM accounts 
    WHERE id = from_account 
    FOR UPDATE;  -- 锁定记录，防止并发操作
    
    -- 检查余额是否充足
    IF current_balance < transfer_amount THEN
        -- 余额不足，回滚事务
        ROLLBACK;
        SELECT '交易失败：余额不足' AS message;
    ELSE
        -- 余额充足，执行转账操作
        
        -- 步骤1：从转出账户扣除金额
        UPDATE accounts 
        SET balance = balance - transfer_amount 
        WHERE id = from_account;
        
        -- 步骤2：向转入账户增加金额
        UPDATE accounts 
        SET balance = balance + transfer_amount 
        WHERE id = to_account;
        
        -- 步骤3：记录交易信息
        INSERT INTO transactions (from_account_id, to_account_id, amount)
        VALUES (from_account, to_account, transfer_amount);
        
        -- 提交事务
        COMMIT;
        SELECT '交易成功' AS message;
    END IF;
END //
DELIMITER ;

-- =============================================
-- 测试转账功能
-- =============================================

-- 执行转账：从账户1向账户2转账100元
CALL transfer_money(1, 2, 100.00);

-- 查看转账后的账户余额
SELECT * FROM accounts;

-- 查看交易记录
SELECT * FROM transactions; 
