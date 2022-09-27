●●●Workflow về cách làm việc trên git●●●
1. Dùng lệnh [git fetch] để load file origin/main mới nhất về
2. Dùng lệnh [git checkout origin/main] chuyển qua branch origin/main
3. Trên branch origin/main tạo ra branch mới với tên branch là tên cộng với ngày làm việc
    Ví dụ: Người làm việc là Nguyên, và ngày là 27/09 thì tạo branch là: Nguyen2709
4. Chuyển qua branch vừa mới tạo và bắt đầu phiên làm việc.
5. Sau khi hoàn thành cập nhật thì dùng lệnh [git push --set-upstream] update lên hệ thống
6. Dùng lệnh [git fetch] để load file origin/main mới nhất về
7. Dùng lệnh [git checkout origin/main] chuyển qua branch origin/main
8. Dùng lệnh merge để merge branch mới tạo vào branch origin/main
Chú ý: Bước 6,7,8 thì chỉ Admin mới thực thi


●●●Tổng hợp những lệnh cần khi sử dụng git●●●

Chương 1 : tổng hợp các lệnh và giải thích
1- git init: khoi tao repository
2- git clone copy các repo từ serve
3- git pull cap nhat du lieu mới từ kho ở remote về kho local
4- git add thêm file mới
5- git commit luu vào csdl nội dung mơi câp nhật
6- git push đưa nội dung mới lên repo
7- git log xem lại lịch sử commit
8- git branch tạo một branch mới song song

Chương 2: Guideline Merge to local main 
//@Nguyên Trần
1. On the main branch: git merge 'branch name'
2. Git add. (if there's a new content in a branch)
3. Git commit
4. git checkout vào  branch cần được cập nhât
5, dùng lệnh merge để chọn branch
Chương 3: Guideline để update main mới nhất lên github repository 
//@Lee Duy
①Save lại nội dung đã chỉnh sửa trên file.

②Commit file mới sau khi chỉnh sửa nội dung: git commit -m "comment nội dung chỉnh sửa"
*Chú ý: Nếu sau khi thực hiện lệnh commit mà xuất hiện thông báo như bên dưới thì dùng lệnh: [git add .] sau đó mới dùng lệnh commit.
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   readme.md

no changes added to commit (use "git add" and/or "git commit -a"
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝

③Kiểm tra tình trạng file trước khi push lên repository: git pull
*Chú ý: Nếu file không trong trạng thái mới nhất sẽ có thông báo như bên dưới. Để giải quyết thì phải đưa file về tình trạng mới nhất, xong sau đó sẽ thực hiện lại từ bước ② để update file lên github repository.
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝
CONFLICT (content): Merge conflict in readme.md
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝

④Update file đã chỉnh sửa lên github repository: git push