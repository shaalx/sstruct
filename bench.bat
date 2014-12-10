go test 
::-test.run="analyse_test.go"
::net stop server
ping /n 18 127.1 >nul
::net start server