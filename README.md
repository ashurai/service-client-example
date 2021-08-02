# Form3 Client application

Init Example:
```
NewClientService := LoadAccountsService(nil)
```

## Instructions
Run whole application command
```
docker-compose up
```

Sample result of Intigration test
```
client_app_1  | Hit:1 http://deb.debian.org/debian buster InRelease
client_app_1  | Hit:2 http://security.debian.org/debian-security buster/updates InRelease
client_app_1  | Hit:3 http://deb.debian.org/debian buster-updates InRelease
client_app_1  | Reading package lists...
client_app_1  | Reading package lists...
client_app_1  | Building dependency tree...
client_app_1  | Reading state information...
client_app_1  | make is already the newest version (4.2.1-1.2).
client_app_1  | 0 upgraded, 0 newly installed, 0 to remove and 29 not upgraded.
client_app_1  | ======================= intigration-test =======================
client_app_1  | === RUN   TestCreateAccountService
client_app_1  | --- PASS: TestCreateAccountService (0.06s)
client_app_1  | === RUN   TestGetAccountByID
client_app_1  | --- PASS: TestGetAccountByID (0.01s)
client_app_1  | === RUN   TestDeleteByID
client_app_1  | --- PASS: TestDeleteByID (0.01s)
client_app_1  | PASS
client_app_1  | ok      github.com/ashurai/form3        (cached)
```

While running successfully your application can also test your client intigration-test by using 
```
make intigration-test
```
