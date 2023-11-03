# 无触心率检测Restful API文档

### User

##### 注册用户

- **URL:** http://124.70.18.154/api/user/register

- **Method:** POST

- Request Body:

  ```json
  jsonCopy code{
    "username": "newuser",
    "password": "password123"
  }
  ```

- Response:

  - Success (HTTP Status: 201 Created)

    ```json
    jsonCopy code{
      "message": "注册成功"
    }
    ```

  - Error (HTTP Status: 400 Bad Request)

    ```json
    jsonCopy code{
      "error": "用户名已存在"
    }
    ```

##### 用户登录

- **URL:** http://124.70.18.154/api/user/login

- **Method:** POST

- Request Body:

  ```json
  jsonCopy code{
    "username": "existinguser",
    "password": "password123"
  }
  ```

- Response:

  - Success (HTTP Status: 200 OK)

    ```json
    jsonCopy code{
      "token": "your-jwt-token"
    }
    ```

  - Error (HTTP Status: 401 Unauthorized)

    ```json
    jsonCopy code{
      "error": "用户名或密码错误"
    }
    ```

##### 获取用户信息（需要JWT身份验证）

- **URL:** http://124.70.18.154/api/protected/get_user_info

- **Method:** GET

- **Headers:**

  - Authorization: Bearer your-jwt-token

- **Response:**

  - Success (HTTP Status: 200 OK)

    ```json
    jsonCopy code{
      "id": 1,
      "username": "existinguser"
    }
    ```

  - Error (HTTP Status: 401 Unauthorized)

    ```json
    jsonCopy code{
      "error": "令牌已过期"
    }
    ```
    
  - Error (HTTP Status: 500 Internal Server Error)
  
    ```JSON
    jsonCopy code{
      "error": "无法验证令牌"
    }
    ```
  
    

##### 删除用户（需要JWT身份验证）

- **URL:** http://124.70.18.154/api/protected/delete_user

- **Method:** DELETE

- **Headers:**

  - Authorization: Bearer your-jwt-token

- **Response:**

  - Success (HTTP Status: 200 OK)

    ```json
    jsonCopy code{
      "message": "删除用户成功"
    }
    ```

  - Error (HTTP Status: 401 Unauthorized)

    ```json
    jsonCopy code{
      "error": "未注册"
    }
    ```

  - Error (HTTP Status: 500 Internal Server Error)

    ```JSON
    jsonCopy code{
      "error": "删除用户失败"
    }
    ```


##### 更新用户（需要JWT身份验证）

- **URL:** http://124.70.18.154/api/protected/update_user_info

- **Method:** PUT

- **Headers:**

  - Authorization: Bearer your-jwt-token

- **Response:**

  - Success (HTTP Status: 200 OK)

    ```json
    jsonCopy code{
      "message": "成功更新用户信息"
    }
    ```

  - Error (HTTP Status: 401 Unauthorized)

    ```json
    jsonCopy code{
      "error": "令牌已过期"
    }
    ```

  - Error (HTTP Status: 500 Internal Server Error)

    ```JSON
    jsonCopy code{
      "error": "无法验证令牌"/"无法更新用户信息"
    }
    ```

### Device

##### 获取设备信息

- **URL:**http://124.70.18.154/api/protected/get_device_info`

- **Method:**GET

- **Response:**

  - Success (HTTP Status: 200 OK)
  
  ```json
  jsonCopy code[
    {
      "id": 1,
      "url": "test",
      "position": "Living Room",
      "connected": true
    },
    {
      "id": 2,
      "url": "camera1",
      "position": "Bedroom",
      "connected": false
    }
    // ...
  ]
  ```
  
  - Error (HTTP Status: 500 Internal Server Error)
  
    ```json
    jsonCopy code{
      "error": "读取设备信息失败"
    }
    ```

