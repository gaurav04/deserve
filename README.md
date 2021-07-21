# talentica
<!---
Later on we can make it part of CI.
-->

Build:
```sudo docker build . -t gaurav04/go-app:latest```

Push:
```sudo docker push gaurav04/go-app:latest```

Run:
```sudo docker run -it --rm -d -p 8081:8081 --name web gaurav04/go-app:latest```

