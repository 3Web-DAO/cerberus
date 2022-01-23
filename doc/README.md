- database

    ```
    ./xorm reverse mysql root:*******@(127.0.0.1:3306)/blockchaindata?charset=utf8 templates/goxorm

    ```
    - user

        ```
        CREATE TABLE `user` (
          `user_id` varchar(9) NOT NULL,
          `password` varchar(20) NOT NULL,
          `user_name` varchar(64) NOT NULL,
          `create_time` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
          PRIMARY KEY (`user_id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
        ```

    - tweet_user

        ```
        CREATE TABLE `twitter_user` (
          `user_id` varchar(12) NOT NULL,
          `id` varchar(12) NOT NULL,
          `name` varchar(255) NOT NULL,
          `last_tweet_id` varchar(32) NOT NULL DEFAULT '0',
          PRIMARY KEY (`user_id`) USING BTREE
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
        ```

    - tweet_follower

        ```
        CREATE TABLE `twitter_follower` (
          `tweet_user_id` varchar(12) NOT NULL,
          `name` varchar(255) NOT NULL,
          `id` varchar(12) NOT NULL,
          PRIMARY KEY (`tweet_user_id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

        ```

    - email

        ```
        CREATE TABLE `email` (
          `user_id` varchar(9) NOT NULL,
          `email_address` varchar(64) NOT NULL,
          PRIMARY KEY (`user_id`)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
        ```


- tweet API

    - GetUserIdByUserName
    - GetLastTweetByUserName
    
