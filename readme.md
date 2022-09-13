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