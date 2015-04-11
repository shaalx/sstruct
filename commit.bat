git add -A
git status
::echo commit for change : 
@echo commit:
@set /p commit_for_change=
git commit -m "%commit_for_change%"
git push oschina master:test
git push github master:test
::net stop server
ping /n 8 127.1 >nul
::net start server