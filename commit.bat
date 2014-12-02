git status
git add -A
::echo commit for change : 
set /p commit_for_change=
git commit -m "%commit_for_change%"
git push osc master
git push origin master
::net stop server
ping /n 8 127.1 >nul
::net start server