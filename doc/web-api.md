- Web

  - error return value 
    
    ```
    {"state": 1000, "msg": ""} #state code greater than zero
    ```

  - addUser
    
      ```
      curl.exe -X POST http://localhost:8080/user/add -d '{"user_name":"clay", "password":"123"}'

      {"state": 0, "msg": ""}
      ```

      

  - addTwitterUser

      ```
      curl.exe -X POST http://localhost:8080/twitteruser/add -d '{"user_name":"clay", "name":"twitter name"}'

      {"state": 0, "msg": ""}
      ```

  - listTwitterUser

      ```
      curl.exe -X POST http://localhost:8080/twitteruser/list -d ''

      {"state": 0, "data": [{"name":"", "id": ""}]}
      ```
  - listTwitterFollwerByTwitterUser

    ```
    curl.exe -X POST http://localhost:8080/twitteruser/follower -d '{"name": ""}'

    {"state": 0, "data": [{"name":"", "id": ""}]}
    ```

  - addEmail

    ```
    curl.exe -X POST http://localhost:8080/email/add -d '{"user_name": "", "email_address": ""}'

    {"state": 0, "msg": ""}
    ```