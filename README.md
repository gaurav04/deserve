# deserve
<!---
Later on we can make it part of CI.
-->

Build:
```sudo docker build . -t gaurav04/deserve:latest```

Push:
```sudo docker push gaurav04/deserve:latest```

Run:
```sudo docker run -it --rm -d -p 8081:8081 --name web gaurav04/deserve:latest```

Result:
```
curl localhost:8081
{"message":"Hello Deserve","branches":"US, Pune, Banglore","container_name":"9767e5baee53"}
```
