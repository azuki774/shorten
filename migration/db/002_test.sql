-- +migrate Up
INSERT INTO `URL_table` (`id`, `short_key`, `target_url`, `expired_at`) VALUES (101, "github", "https://github.com/", "2030-01-01 00:00:00");
INSERT INTO `URL_table` (`id`, `short_key`, `target_url`, `expired_at`) VALUES (102, "youtube", "https://www.youtube.com/?gl=JP&hl=ja", "2030-01-01 00:00:00");
INSERT INTO `URL_table` (`id`, `short_key`, `target_url`, `expired_at`) VALUES (103, "indigo", "https://compas.arena.ne.jp/?service_code=indigo", "2010-01-01 00:00:00");

-- +migrate Down
DELETE FROM `URL_table` WHERE `id` = 101;
DELETE FROM `URL_table` WHERE `id` = 102;
DELETE FROM `URL_table` WHERE `id` = 103;
