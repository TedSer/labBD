

# 6
1. Go Computer engine -> VM instance  -> Press the button "create instance" (img11)
![ssh-hd-insight](../img/6/11.png)
2. Change "boot disk" operating system into the "Ubuntu" (Note! for this lab I have used version 16.04)
3. Allow HTTP/HTTPS traffic 
4. Press the button "Create"

5. Find Network interfaces chapter and click on "View details" (img13)
![ssh-hd-insight](../img/6/13.png)

6. On your left menu find the "Firewall" chapter. Create a new firewall rule 

* 6.a elasticsearch
* 6.1 turn on "logs"
* 6.2 "Targets" needs to be "All instances in the network"
* 6.3 "Source IP ranges " - 0.0.0.0/0
* 6.4 At Protocols and ports match "tcp" and write 9200
* 6.5 Press create 
* 6.b kibana
* 6.6 turn on "logs"
* 6.7 "Targets" needs to be "All instances in the network"
* 6.8 "Source IP ranges " - 0.0.0.0/0
* 6.9 At Protocols and ports match "tcp" and write 5601
* 6.10 Press create 
 
7. Go Computer engine -> Your instance -> press "SSH" to open a console (NOTE! for further actions I have used "ttps://logz.io/blog/elk-stack-google-cloud/")

8. Use those command 

8.a lasticsearch
```
1  sudo apt-get install default-jre
2  sudo apt update
3  sudo apt install apt-transport-https
4  wget -qO - https://artifacts.elastic.co/GPG-KEY-elasticsearch | sudo apt-key add -
5  sudo sh -c 'echo "deb https://artifacts.elastic.co/packages/7.x/apt stable main" > /etc/apt/sources.list.d/elastic-7.x.list'
6  sudo apt update
7  sudo apt install elasticsearch
8  sudo service elasticsearch status
9  sudo systemctl enable elasticsearch.service
10  sudo systemctl start elasticsearch.service
11  curl -X GET "localhost:9200/"
12  sudo nano /etc/elasticsearch/elasticsearch.yml   
(Change networkhost into "0.0.0.0")
(Change discovery seed hosts into "[]")
13  sudo service elasticsearch restart
````

8.1 Go Your VM instance -> Network interfaces -> Find "External IP" and copy it. In a new tab paste your external IP + :9200. 
You should have smth like this. (img14)
![ssh-hd-insight](../img/14.png)


8.b kibana

```
14  sudo apt-get install apt-transport-https
15  echo "deb https://artifacts.elastic.co/packages/5.x/apt stable main" | sudo tee -a /etc/apt/sources.list.d/elastic-5.x.list
16  sudo apt-get update
17  sudo apt-get install logstash
18  sudo service logstash start
19  echo "deb http://packages.elastic.co/kibana/7.0/debian stable main" | sudo tee -a /etc/apt/sources.list.d/kibana-7.0.x.list
20  sudo wget --directory-prefix=/opt/ https://artifacts.elastic.co/downloads/kibana/kibana-7.6.1-amd64.deb
21  sudo dpkg -i /opt/kibana*.deb
22  sudo apt-get update
23  sudo apt-get install kibana
24  sudo nano /etc/kibana/kibana.yml 
(Change server.host into "0.0.0.0")
(Change server.port into "5601")
25  sudo service kibana start
26  sudo service kibana status
```
8.1 Go Your VM instance -> Network interfaces -> Find "External IP" and copy it. In a new tab paste your external IP + :5601. 
   You should have smth like this. (img15)
   ![ssh-hd-insight](../img/15.png)

9. Prepare AZURE

* 9.1 Create new or use res group from 5 lab(NOTE! You can review how to create a new res grop on top of the guidepage for lab 5)
* 9.2 Search "logic apps" and press "+" button 
* 9.3 "In Development Tools" chapter find logic app designer
* 9.4 Go to the Templates and choose a "Blank logic app"
* 9.5 Search for "Event hubs" and make configs like at the screenshot (img16)
![ssh-hd-insight](../img/16.png)
* 9.6 Search for "HTTP" and make configs like at the screenshot (img17)(NOTE! Find your own IP range)
![ssh-hd-insight](../img/17.png)

