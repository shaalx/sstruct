git add -A
git status
::echo commit for change : 
@echo commit:
@set /p commit_for_change=
git commit -m "%commit_for_change%"
git push oschina master:mast
git push github master:mast
::net stop server
ping /n 8 127.1 >nul
::net start server