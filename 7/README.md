# 5
1. Create res group (Note! Recommended to pick WEST EUROPE region)

2. Create Azure cash for Redis(the same region as res group). Wait till status "running"

3. Create an Event hub  
 3.1 Press the button "+ Event hub"(img1). It would create the instance.
 3.2 Find your innctance in "Entetis -> Evevnt hub" 
 3.3 Create new Share access policies (img2). (Note! select "Manage" config while creating access policies to avoid further errors)

4. Paste your data from the Share access policies into the "Service" file according to the screenshot( follow comments on the img3)

5. Go "Redis" -> "Access keys". Paste your data from Access keys into the code according to the screenshot (follow comments on the img4)
(Don't forget to migrate CACHE_HOSTNAME and CACHE_KEY to another service file)

6. This is how it should look like (img5)

7. Start Postman and post your JSON file URL (Note! check the ports 9000 or 10000)
 7.1 Start sending POST request into the "eventHub" (img6)
 7.2 You could notice some "Documents" stream in the terminal (img7)
 7.3 Go to the Event hub -> Process data -> Explore. Wait till refreshing. Here you can see proof that your data has been written (img8)
 7.4 Start sending POST request into the "redis"(img9)
 7.5 You could find that there are different types of requests are coming through your terminal.
 7.6 Go to the "Azure cache for Redis" -> your instance -> "console" button. Write "hgetAll ConsoleLog"(img10)(Note! carefull with console comand naming.)

# 6
1. Go Computer engine -> VM instance  -> Press the button "create instance" (img11)
2. Change "boot disk" operating system into the "Ubuntu" (Note! for this lab I have used version 16.04)
3. Allow HTTP/HTTPS traffic 
4. Press the button "Create"

5. Find Network interfaces chapter and click on "View details" (img13)
6. On your left menu find the "Firewall" chapter. Create a new firewall rule 
  6.a elasticsearch
 6.1 turn on "logs"
 6.2 "Targets" needs to be "All instances in the network"
 6.3 "Source IP ranges " - 0.0.0.0/0
 6.4 At Protocols and ports match "tcp" and write 9200
 6.5 Press create 
  6.b kibana
 6.6 turn on "logs"
 6.7 "Targets" needs to be "All instances in the network"
 6.8 "Source IP ranges " - 0.0.0.0/0
 6.9 At Protocols and ports match "tcp" and write 5601
 6.10 Press create 
 
 7. Go Computer engine -> Your instance -> press "SSH" to open a console (NOTE! for further actions I have used "ttps://logz.io/blog/elk-stack-google-cloud/")

 8. Use those command 

  8.a lasticsearch
 '
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
   '
   8.1 Go Your VM instance -> Network interfaces -> Find "External IP" and copy it. In a new tab paste your external IP + :9200. 
   You should have smth like this. (img14)


   8.b kibana

  '
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
 '
   8.1 Go Your VM instance -> Network interfaces -> Find "External IP" and copy it. In a new tab paste your external IP + :5601. 
   You should have smth like this. (img15)

   9. Prepare AZURE
   9.1 Create new or use res group from 5 lab(NOTE! You can review how to create a new res grop on top of the guidepage for lab 5)
   9.2 Search "logic apps" and press "+" button 
   9.3 "In Development Tools" chapter find logic app designer
   9.4 Go to the Templates and choose a "Blank logic app"
   9.5 Search for "Event hubs" and make configs like at the screenshot (img16)
   9.6 Search for "HTTP" and make configs like at the screenshot (img17)(NOTE! Find your own IP range)

# 7

# 8-9
1. new res group 
2. add Databriks
3. add Storrage account
4. create container from Storrage account
5. Add directory from a container
6. Go to databriks from res group and create a new cluster
7. Go Active registration -> App registration 
8. Certification & secretes -> new client secret
9. Res froup Access control