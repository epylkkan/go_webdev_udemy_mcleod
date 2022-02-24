
GCP, see the lab or https://acloudguru.com/hands-on-labs/global-load-balancing-with-google-compute-engine

=======

Follow the steps in the exercise 031. GCP:

create your binary

GOOS=linux GOARCH=amd64 go build -o amigos

Create VM amigos-1

SSH from the consolo

Upload binary and templates  (SSH -> star)

chmod 777 amigos 

// RUN ./amigos 
//check it in a browser at [public-IP]

Persisting your application


Create a configuration file - sudo nano /etc/systemd/system/amigos.service

[Unit]

Description=Go Server


[Service]

ExecStart=/home/epylkkan/amigos

WorkingDirectory=/home/epylkkan

User=root

Group=root

Restart=always


[Install]

WantedBy=multi-user.target


Add the service to systemd. - sudo systemctl enable amigos.service

Activate the service. - sudo systemctl start amigos.service

Check if systemd started it. - sudo systemctl status amigos.service

Stop systemd if so desired. - sudo systemctl stop amigos.service

=========

1. Create fw rule for the health check


Click VPC network > Firewall. Notice the existing ICMP, internal, RDP, and SSH firewall rules.

Each Google Cloud project starts with the default network and these firewall rules.

Click Create Firewall Rule.

Specify the following, and leave the remaining settings as their defaults:

Property

Value (type value or select option as specified)

Name

fw-allow-health-checks

Network

default

Targets

Specified target tags

Target tags

allow-health-checks

Source filter

IP Ranges

Source IP ranges

130.211.0.0/22, 35.191.0.0/16

Protocols and ports

Specified protocols and ports

Make sure to include the /22 and /16 in the Source IP ranges.

Select tcp and specify port 80.



2. Create NAT using cloud router (no public IP's for VM's)

In the Cloud Console, on the Navigation menu (), click Network services > Cloud NAT.

Click Get started.

Specify the following, and leave the remaining settings as their defaults:

Property

Value (type value or select option as specified)

Gateway name

nat-config

Region

europe-north

Click Cloud Router, and select Create new router.

For Name, type nat-router-us-central1.

Click Create.

In Create a NAT gateway, click Create.


3. Create custom image incl. application 

Select Keep Disk

Internal IP only

Netwwork tags: "allow-health-checks"

Set application to start at boot
    sudo systemctl enable amigos.service 
    sudo systemctl start amigos.service etc.

     

5. Create a new Image from the Disk (you can delete the disk)


6. Create Instance Template from the Image

In the Cloud Console, on the Navigation menu (), click Compute Engine > Instance templates.

Click Create instance template.

For Name, type mywebserver-template.

For Series, select N1.

For Machine type, select f1-micro (1 vCPU).

For Boot disk, click Change.

Click Custom images.

For Image, Select mywebserver.

Click Select.

Click Management, security, disks, networking, sole tenancy.

 Click Networking.

    ◦ For Network tags, type allow-health-checks.

   ◦ Under External IP dropdown, select None.

• Click Create.


7. Create Managed Instance Group from the Instance Template + Health Check

On the Navigation menu, click Compute Engine > Instance groups.

Click Create Instance group.

Specify the following, and leave the remaining settings as their defaults:

Property

Value (type value or select option as specified)

Name

amigos-mig

Location

Multiple zones

Region
europe-north1

Instance template

amigos-template

Under Autoscaling metrics, click on the edit pencil icon.

Under Metric type, select HTTP load balancing utilization.

Enter Target HTTP load balancing utilization to 80.

Click Done.

Set Cool down period to 60 seconds.

Enter Minimum number of instances 1 and Maximum number of instances 2.

Managed instance groups offer autoscaling capabilities that allow you to automatically add or remove instances from a managed instance group based on increases or decreases in load. Autoscaling helps your applications gracefully handle increases in traffic and reduces cost when the need for resources is lower. You just define the autoscaling policy, and the autoscaler performs automatic scaling based on the measured load.

For Health check, select Create a health check.

Specify the following, and leave the remaining settings 
as their defaults:

Property
Value (select option as specified)

Name
http-health-check

Protocol
TCP

Port
80

Managed instance group health checks proactively signal to delete and recreate instances that become unhealthy.

Click Save and continue.

For Initial delay, type 60. This is how long the Instance Group waits after initializing the boot-up of a VM before it tries a health check. You don't want to wait 5 minutes for this during the lab, so you set it to 1 minute.

Click Create.
    
NOTE: If a warning window will appear stating that There is no backend service attached to the instance group. Ignore this; you will configure the load balancer with a backend service in the next section of the lab.


8. Configure HTTP(S) Load Balancer 

On the Navigation menu, click Network Services > Load balancing.

Click Create load balancer.

Under HTTP(S) Load Balancing, click Start configuration.

Select From Internet to my VMs, then click Continue.

For Name, type http-lb.

Configure the backend

Backend services direct incoming traffic to one or more attached backends. Each backend is composed of an instance group and additional serving capacity metadata.

Click Backend configuration.
For Backend services & backend buckets, click Create or select backend services & backend buckets > Backend services > Create a backend service.

Specify the following, and leave the remaining settings as their defaults:

Property
Value (select option as specified)

Name
http-backend

Backend type

Instance group
amigos-mig

Port numbers
80

Balancing mode
Rate

Maximum RPS
50

Capacity
100

This configuration means that the load balancer attempts to keep each instance at or below 50 requests per second (RPS).

Click Done.

Click Add backend.

For Health Check, select http-health-check (TCP).

Click check for the Enable logging checkbox.

Specify Sample rate as 1.

Click Create.

Configure the frontend

Click Frontend configuration.

Specify the following, and leave the remaining settings as their defaults:

Property
Value (type value or select option as specified)

Protocol
HTTP

IP version
IPv4

IP address
Ephemeral

Port
80

Click Done.

Click Add Frontend IP and port.

Specify the following, and leave the remaining settings as their defaults:

Property
Value (type value or select option as specified)

Protocol
HTTP

IP version
IPv6

IP address
Ephemeral

Port
80

Click Done.

HTTP(S) load balancing supports both IPv4 and IPv6 addresses for client traffic. Client IPv6 requests are terminated at the global load balancing layer and then proxied over IPv4 to your backends


Stress test the HTTP load balancer


============

AWS:

Create auto scaling
Auto Scaling helps you maintain application availability and allows you to scale your Amazon EC2 capacity up or down automatically according to conditions you define. You can use Auto Scaling to help ensure that you are running your desired number of Amazon EC2 instances. Auto Scaling can also automatically increase the number of Amazon EC2 instances during demand spikes to maintain performance and decrease capacity during lulls to reduce costs. Auto Scaling is well suited both to applications that have stable demand patterns or that experience hourly, daily, or weekly variability in usage.

Configure auto scaling
EC2 / autoscaling / launch configuration
create auto scaling group / create launch configuration
My AMIs / choose your AMI
next / next
configure details
name: auto-scale-config-2019-10-31
next / next
configure security group
select an existing security group / select the "web-tier" security group
next / next / create launch configuration
choose an existing key pair / create launch configuration
Create auto scaling group
Configure auto scaling group
name: auto-scale-group-2019-10-31
group size: this is the minimum number of instances we'll always be running
network: default vpc
subnet: choose the availability zones (AZs) into which you've launched instances
advanced details
load balancing: check "receive traffic from elastic load balancer"
select your load balancer
health check: ELB (this is what we set up)
configure scaling policies
keep group at initial size
configure tags
value: web-server-auto-scaled
create auto scaling group
Scaling policies
this is where we'd add policies to say when we scale up / scale down
