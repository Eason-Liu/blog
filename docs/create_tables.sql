CREATE TABLE `blog` (
                        `id` int NOT NULL AUTO_INCREMENT COMMENT '博客Id',
                        `title_img` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题的图片',
                        `title_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题的名称',
                        `sub_title` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '子标题',
                        `content` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章的内容',
                        `author` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章作者',
                        `create_at` int NOT NULL COMMENT '创建时间, 时间戳',
                        `update_at` int NOT NULL COMMENT '更新时间',
                        `publish_at` int NOT NULL COMMENT '发布时间',
                        `status` tinyint(1) NOT NULL COMMENT '0: 草稿, 1: 发布',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `tag` (
                       `id` int NOT NULL AUTO_INCREMENT COMMENT '标签Id',
                       `key` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签的名称',
                       `value` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签的值',
                       `blog_id` int NOT NULL COMMENT '关联的文章',
                       `create_at` int NOT NULL COMMENT '创建时间',
                       `color` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签的颜色',
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `idx_id` (`key`,`value`) USING BTREE COMMENT 'key和value构成tag的唯一标识'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;