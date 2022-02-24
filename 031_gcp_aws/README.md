Deploying our session example

1)
change your port number from 8080 to 80

2)
create your binary
GOOS=linux GOARCH=amd64 go build -o expiresession


3)
GCP

Create VM expire
SSH from the consolo
mkdir expire -> cd expire -> mkdir templates
Upload binary and templates  (SSH -> star)
chmod 777 expiresession
RUN ./expiresession

check it in a browser at [public-IP]

Persisting your application
To run our application after the terminal session has ended, we must do the following:

Create a configuration file - sudo nano /etc/systemd/system/expire.service

[Unit]
Description=Go Server

[Service]
ExecStart=/home/epylkkan/expiresession
WorkingDirectory=/home/epylkkan
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target

Add the service to systemd. - sudo systemctl enable expire.service
Activate the service. - sudo systemctl start expire.service
Check if systemd started it. - sudo systemctl status expire.service
Stop systemd if so desired. - sudo systemctl stop expire.service


4)
AWS

SSH into your server
ssh -i /path/to/[your].pem ubuntu@[public-DNS]:
create directories to hold your code
for example, "wildwest" & "wildwest/templates"
copy binary to the server

copy your "templates" to the server

scp -i /path/to/[your].pem templates/* ubuntu@[public-DNS]:/home/ubuntu/templates
chmod permissions on your binary

Run your code

sudo ./[some-name]
check it in a browser at [public-IP]
Persisting your application
To run our application after the terminal session has ended, we must do the following:

Create a configuration file - sudo nano /etc/systemd/system/<filename>.service
[Unit]
Description=Go Server

[Service]
ExecStart=/home/<username>/<path-to-exe>/<exe>
WorkingDirectory=/home/<username>/<exe-working-dir>
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
Add the service to systemd. - sudo systemctl enable <filename>.service
Activate the service. - sudo systemctl start <filename>.service
Check if systemd started it. - sudo systemctl status <filename>.service
Stop systemd if so desired. - sudo systemctl stop <filename>.service
FOR EXAMPLE
[Unit]
Description=Go Server

[Service]
ExecStart=/home/ubuntu/cowboy
WorkingDirectory=/home/ubuntu
User=root
Group=root
Restart=always

[Install]
WantedBy=multi-user.target
