DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
    `id` mediumint(9) NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    PRIMARY KEY(`id`)
);