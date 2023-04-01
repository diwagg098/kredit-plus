how to run this project :

first step
 - run apache on your local computer
 - run mysql on your local computer
 - run sql script to create new database kredit-plus "CREATE DATABASE kredit-plus";
 
create env
 - create .env file copy from .env.example
 
how to run docker image
 - run command "docker compose up --build" for build docker compose image it will be migration columns database
 - run command "docker compose up -d"
 
how to create minio bucket
 - open minio console in browser http://127.0.0.1:9000
 - login minio :
      - username = userminio
      - password = kiasu123

 - create access key in menu access key
      - access key = AKIAIOSFODNN7EXAMPLE
      - secret key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY 
 - look like this image
<img width="947" alt="image" src="https://user-images.githubusercontent.com/61501287/229271238-3d526f9d-23cd-4668-b100-4ed2f33fa688.png">

- how to create bucket image :
 - create bucket with name asset
<img width="958" alt="image" src="https://user-images.githubusercontent.com/61501287/229271416-82982f74-8418-4d38-975e-968965f13793.png">
<img width="960" alt="image" src="https://user-images.githubusercontent.com/61501287/229271431-1d037a0f-ff2b-43dc-b873-4e1673aaf2b0.png">

how to run unit test cmd :
 - go test -v ./pkg/unit_testing
 
