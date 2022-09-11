●●●Tổng hợp những lệnh cần khi sử dụng git●●●

Chương 1 : tổng hợp các lệnh và giải thích 
//@Khoi Vuong

Chương 2: Guideline Merge to local main 
//@Nguyên Trần

Chương 3: Guideline để update main mới nhất lên github repository 
//@Lee Duy
①Save lại nội dung đã chỉnh sửa trên file.

②Commit file mới sau khi chỉnh sửa nội dung: git commit -m "comment nội dung chỉnh sửa"
*Chú ý: Nếu sau khi thực hiện lệnh commit mà xuất hiện thông báo như bên dưới thì dùng lệnh: [git add .] xong sau đó mới dùng lệnh commit.
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   readme.md

no changes added to commit (use "git add" and/or "git commit -a")
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝

③Kiểm tra tình trạng file trước khi push lên repository: git pull
*Chú ý: Nếu file không trong trạng thái mới nhất sẽ có thông báo như bên dưới. Trong trường hợp này thì cần dùng lệnh [git ]
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝
CONFLICT (content): Merge conflict in readme.md
＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝

④Update file đã chỉnh sửa lên github repository: git push