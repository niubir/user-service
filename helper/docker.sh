docker build -t niubir/user-service:latest .

docker run -p 10011:10011 -p 10012:10012 -v /root/user-service/config:/config -v /root/user-service/data/:/data -v /root/user-service/logs/:/logs --name=user-service-1 -d niubir/user-service:latest

docker push niubir/user-service:latest