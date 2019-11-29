DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` mediumint(9) NOT NULL AUTO_INCREMENT,
    `body` varchar(2000) NOT NULL,
    `post_id` mediumint(9) NOT NULL,
    PRIMARY KEY(`id`)
);