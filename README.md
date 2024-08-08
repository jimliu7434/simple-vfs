# simple-vfs

A simple virtual file system that can be used to store files in memory.

## 目錄結構

```text
simple-vfs
├── README.md
├── main.go         # 程式進入點
├── go.mod
├── go.sum
└── internal
    ├── cmds        # 命令處理函式
    │   ├── file      # 命令 category: file
    │   ├── folder    # 命令 category: folder
    │   └── user      # 命令 category: user
    ├── entity      # 實體物件定義
    │   └── storage   # 儲存體物件定義
    └── logger      # 日誌輸出
```

## 編譯

1. 確保您已安裝 [Go](https://golang.org/doc/install)，版本 `1.20` 以上。
2. git clone

    ```sh
    git clone https://github.com/jimliu7434/simple-vfs.git
    ```

3. 進入專案目錄並安裝 dependencies：

    ```sh
    cd simple-vfs
    go mod tidy
    ```

4. 編譯專案：

    ```sh
    go build -o svfs
    ```

## 命令

### User Management

#### > `register` [username]

註冊一個新用戶  

* Arguments
  * `username` - 用戶名稱 (必填) (英數字組合，長度 3~20 字元)

* Response
  * add user [username] successfully
  * Error: The [username] has already existed
  * Error: The [username] contain invalid chars

* Example

  ```sh
  > register myuser
  add user myuser successfully
  ```

### Folder Management

#### > `create-folder` [username] [foldername] [description?]

在指定的用戶資料中，創建一個新資料夾  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填) (不得包含特殊符號，長度 1 ~ 50 字元)
  * `description` - 資料夾描述 (選填)

* Response
  * create folder [foldername] successfully
  * Error: The [username] doesn't exist
  * Error: The [foldername] contain invalid chars
  * Error: The [foldername] has already existed

* Example

  ```sh
  > create-folder myuser myfolder this is my folder
  create folder myfolder successfully
  ```
  
#### > `delete-folder` [username] [foldername]

在指定的用戶資料中，刪除一個資料夾  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填)

* Response

  * delete folder [foldername] successfully
  * Error: The [username] doesn't exist
  * Error: The [foldername] doesn't exist

* Example

  ```sh
  > delete-folder myuser myfolder
  delete folder myfolder successfully
  ```
  
#### > `rename-folder` [username] [foldername] [newfoldername]

在指定的用戶資料中，重新命名一個資料夾  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填)
  * `newfoldername` - 新資料夾名稱 (必填) (不得包含特殊符號，長度 1 ~ 50 字元)

* Response
  * rename folder [foldername] to [newfoldername] successfully
  * Error: The [username] doesn't exist
  * Error: The [foldername] doesn't exist
  * Error: The [newfoldername] contain invalid chars
  * Error: The [newfoldername] has already existed

* Example

  ```sh
  > rename-folder myuser myfolder mynewfolder
  rename folder myfolder to mynewfolder successfully
  ```
  
#### > `list-folders` [username] [--sort-name asc|desc] [--sort-created asc|desc]

列出指定用戶的所有資料夾  
可選擇依照名稱或建立時間排序，預設為 **名稱 + 升冪排序**  

當 `--sort-name` 與 `--sort-created` 同時存在時，以 `--sort-name` 為主  
當 `--sort-name` 或 `--sort-created` 不存在時，預設為 `--sort-name asc`  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `--sort-name` - 依照名稱排序 (選填) (asc: 升冪, desc: 降冪) (預設: asc)
  * `--sort-created` - 依照建立時間排序 (選填) (asc: 升冪, desc: 降冪) (預設: asc)

* Response
  
  ```bash
  Folder\tDesc\tCreate At\tOwner
  [foldername1]\t[description1]\t[created_at1]\t[username]
  [foldername2]\t[description2]\t[created_at2]\t[username]
  ...
  ```
  
* Example

  ```sh
  > list-folders myuser --sort-name desc
  Folder    Desc                Create At   Owner
  myfolder2 this is my folder2  2021-01-01  myuser
  myfolder1 this is my folder1  2021-01-02  myuser
  
  > list-folders myuser --sort-created desc
  Folder    Desc                Create At   Owner
  myfolder1 this is my folder1  2021-01-02  myuser
  myfolder2 this is my folder2  2021-01-01  myuser
  ```
  
### File Management

#### > `create-file` [username] [foldername] [filename] [description?]

在指定的用戶資料夾創建一個新檔案  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填)
  * `filename` - 檔案名稱 (必填) (不得包含特殊符號，長度 1 ~ 50 字元)
  * `description` - 資料夾描述 (選填)

* Response
  * create file [filename] successfully
  * Error: user [username] doesn't exist
  * Error: folder [foldername] doesn't exist
  * Error: file [filename] contain invalid chars
  * Error: file [filename] has already existed

* Example

  ```sh
  > create-file myuser myfolder myfile1 this is my folder
  create file myfile1 successfully
  ```
  
#### > `delete-file` [username] [foldername] [filename]

在指定的用戶資料夾刪除一個檔案  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填)

* Response

  * delete file [filename] successfully
  * Error: user [username] doesn't exist
  * Error: folder [foldername] doesn't exist
  * Error: file [filename] doesn't exist

* Example

  ```sh
  > delete-file myuser myfolder myfile1
  delete file myfile1 successfully
  ```
  
#### > `list-files` [username] [foldername] [--sort-name asc|desc] [--sort-created asc|desc]

列出指定用戶指定資料夾內的所有檔案  
可選擇依照名稱或建立時間排序，預設為 **名稱 + 升冪排序**  

當 `--sort-name` 與 `--sort-created` 同時存在時，以 `--sort-name` 為主  
當 `--sort-name` 或 `--sort-created` 不存在時，預設為 `--sort-name asc`  

* Arguments
  * `username` - 用戶名稱 (必填)
  * `foldername` - 資料夾名稱 (必填)
  * `--sort-name` - 依照名稱排序 (選填) (asc: 升冪, desc: 降冪) (預設: asc)
  * `--sort-created` - 依照建立時間排序 (選填) (asc: 升冪, desc: 降冪) (預設: asc)

* Response
  
  ```bash
  File\tDesc\tCreate At\tFolder\tOwner
  [filename1]\t[description1]\t[created_at1]\t[foldername]\t[username]
  [filename1]\t[description2]\t[created_at2]\t[foldername]\t[username]
  ...
  ```
  
* Example

  ```sh
  > list-files myuser myfolder --sort-name desc
    File      Desc                Create At   Folder    Owner
    myfile2   this is my file2    2021-01-01  myfolder  myuser
    myfile1   this is my file1    2021-01-02  myfolder  myuser
    
  
  > list-files myuser myfolder --sort-created desc
    File      Desc                Create At   Folder    Owner
    myfile1   this is my file1    2021-01-02  myfolder  myuser
    myfile2   this is my file2    2021-01-01  myfolder  myuser
  ```

### Commands Summary

* All response messages of successful or warning are output to **stdout**
* All error messages are output to **stderr**
* All `[...]` are must arguments
* All `[...?]` are **optional** arguments
* `[username]` can use any **alphanumeric** characters and underscores, and must be **3 to 20** characters long
* `[foldername]` can use any **alphanumeric** characters and underscores, and must be **1 to 50** characters long (and so as `[newfoldername]`)
* `[filename]` can use any **alphanumeric** characters and underscores, and must be **1 to 50** characters long
* press `Ctrl + C` to exit the program (or `Cmd + C` on macOS)

## 測試

1. 確認已安裝 [Go](https://golang.org/doc/install)，版本 `1.20` 以上。
2. 進入專案目錄並執行測試：

    ```sh
    go test -cover ./...
    ```

3. 產出測試覆蓋率報告：

    ```sh
    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out
    ```
