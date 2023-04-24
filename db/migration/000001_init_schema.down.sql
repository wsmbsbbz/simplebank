-- 由于entries和transfers表中引用了accounts表的id字段，所以需要先删除这两个表
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;