-- 删除并重新创建 school 数据库
DROP DATABASE IF EXISTS school;
CREATE DATABASE school CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
USE school;

-- 创建 class 表
CREATE TABLE IF NOT EXISTS class (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    grade_level INT NOT NULL,
    teacher_name VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL
) CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 创建 user 表  
CREATE TABLE IF NOT EXISTS user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    gender ENUM('男', '女') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    age INT NOT NULL,
    class_id INT NOT NULL,
    FOREIGN KEY (class_id) REFERENCES class(id)
) CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 插入班级数据
INSERT INTO class (name, grade_level, teacher_name) VALUES
('一年级一班', 1, '张老师'),
('一年级二班', 1, '李老师'),
('二年级一班', 2, '王老师'), 
('二年级二班', 2, '陈老师'),
('三年级一班', 3, '刘老师'),
('三年级二班', 3, '赵老师'),
('四年级一班', 4, '孙老师'),
('四年级二班', 4, '周老师'),
('五年级一班', 5, '吴老师'),
('五年级二班', 5, '郑老师');

-- 随机生成500名学生数据
INSERT INTO user (name, gender, age, class_id)
WITH RECURSIVE cte AS (
    SELECT 1 AS n, FLOOR(1 + RAND() * 10) AS class_id
    UNION ALL
    SELECT n + 1, FLOOR(1 + RAND() * 10) 
    FROM cte WHERE n < 500
)
SELECT 
    CONCAT('学生', n) AS name,
    IF(RAND() > 0.5, '男', '女') AS gender,
    FLOOR(6 + RAND() * 6) AS age,
    class_id
FROM cte;
