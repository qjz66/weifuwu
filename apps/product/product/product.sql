CREATE DATABASE product;
USE product;

CREATE TABLE `product` (
                           `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
                           `name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
                           `detail` text COMMENT '商品详情',
                           `price` decimal(20,2) NOT NULL DEFAULT 0 COMMENT '价格,单位-元保留两位小数',
                           `stock` int(11) NOT NULL DEFAULT 0 COMMENT '库存数量',
                           `status` int(6) NOT NULL DEFAULT 1 COMMENT '商品状态.1-在售 2-下架 3-删除',
                           `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商品表';