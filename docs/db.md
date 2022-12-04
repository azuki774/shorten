## DB Scheme

```
MariaDB [shorten]> show columns from URL_table;
+------------+--------------+------+-----+---------------------+-------------------------------+
| Field      | Type         | Null | Key | Default             | Extra                         |
+------------+--------------+------+-----+---------------------+-------------------------------+
| id         | int(11)      | NO   | PRI | NULL                | auto_increment                |
| short_key  | varchar(255) | NO   | UNI | NULL                |                               |
| target_url | varchar(255) | NO   |     | NULL                |                               |
| expired_at | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
| created_at | timestamp    | NO   |     | current_timestamp() |                               |
| updated_at | timestamp    | NO   |     | current_timestamp() | on update current_timestamp() |
+------------+--------------+------+-----+---------------------+-------------------------------+
```
