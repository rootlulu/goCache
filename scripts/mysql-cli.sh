sudo docker run -it --rm --name mysql-cli --link mysql:mysql mysql:5.7 mysql -h mysql -u root -p --default-character-set=utf8mb4


# create table:
# CREATE TABLE user (
#     id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
#     name VARCHAR(255) NOT NULL,
#     sex BOOLEAN NOT NULL DEFAULT TRUE,
# 	orgId INT UNSIGNED NOT NULL,
# 	age INT UNSIGNED NOT NULL DEFAULT 18
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# insert data:
# INSERT INTO `user` (name, sex, orgId, age) VALUES
# ('lulu1', 1, 3, 25),
# ('xiaomi2', 0, 2, 32),
# ('James', 1, 5, 19),
# ('Russel', 1, 1, 28),
# ('Brrrk', 0, 4, 45),
# ('Baiden', 0, 3, 22),
# ('Tramp', 1, 2, 36),
# ('Harden', 0, 5, 29),
# ('Green', 1, 1, 51),
# ('John', 0, 4, 27);
