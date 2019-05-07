# bookery

![alt text](https://github.com/hassannmoussaa/bookery/raw/master/static/src/img/logo.png)

### Admin Credentials

Email : hassannmoussaa@gmail.com
<br /> 
Password : 123456

### Binary

```
GOOS=linux GOARCH=amd64 go build -v -o bookery github.com/hassannmoussaa/bookery
GOOS=darwin GOARCH=amd64 go build -v -o bookery github.com/hassannmoussaa/bookery
GOOS=windows GOARCH=amd64 go build -v -o bookery.exe github.com/hassannmoussaa/bookery

tar --exclude ./.git --exclude ./static --exclude ./uploads --exclude ./pkg --exclude ./main.go -czf ~/Desktop/bookery.tar .

```

### HTTPS proxy
```
httpsify --backends="http://<IP>:80" --domains="bookery,www.bookery,api.bookery" --cert="./cert.pem" --key="./key.pem" --address="<IP>"
```