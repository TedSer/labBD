
# 8-9
1. Create new or use already existed res group 
2. Add Databriks and Storage account into res group 

![ssh-hd-insight](../8-9/img/1.png)
3. Go Container from Storrage account(Data lake chapter) -> Add container -> Create 
* (NOTE!Hierarchical namespace – Enabled. Replecation – LRS) 
4. Add a directory 
* (NOTE! Final result should be displayed here)

5. Go Active registration -> App registration 

![ssh-hd-insight](../8-9/img/2.png)

6. Certification & secretes -> new client secret
* (NOTE! Copy VALUE field as it would be shown only once and we need it to use in next steps)

7. Go Res group -> Access control -> Role assignment -> Add
* (NOTE! Choose not the user but your instance)

![ssh-hd-insight](../8-9/img/3.png)

8. Go from your Res group to databriks service 

![ssh-hd-insight](../8-9/img/4.png)

9. Pres "Launch workspace"
* (NOTE! If there is no button "Launch workspace" wait till databriks will deploy)

10. Create a new cluster with the following setup 

![ssh-hd-insight](../8-9/img/5.png)

* (NOTE! Use 60 minutes terminates to reduce costs)

11. While the cluster is creating go to the databricks main screen and create 2 Notebooks: 1 for python and 1 for Scala

![ssh-hd-insight](../8-9/img/6.png)

12. When creating is finished go from your cluster -> libraries -> install new 

13. Choose Maven and paste ```com.microsoft.azure:azure-eventhubs-spark_2.12:2.3.18-9 ```
* (NOTE! Try to restart the whole cluster if lib does not deploy)


![ssh-hd-insight](../8-9/img/7.png)

14. We need to give acces to our folder and container we created before. Use Azure Storage Explorer.

*  find your folder and press Manage ACLs. Add your instance 
*  give those permissions 

![ssh-hd-insight](../8-9/img/12.png)

*  same but now with container

![ssh-hd-insight](../8-9/img/13.png)



15. Go to your Python notebook from Workspace and paste: (as it is on screen)

```
configs = {"fs.azure.account.auth.type": "OAuth",
         "fs.azure.account.oauth.provider.type": "org.apache.hadoop.fs.azurebfs.oauth2.ClientCredsTokenProvider",
         "fs.azure.account.oauth2.client.id": "8443c1e9-fe30-4ecb-b3ce-f58f43473a5b",
         "fs.azure.account.oauth2.client.secret": "3nP_vZ37iAD4Ve449d~FFPe_lUUok17~-8",
         "fs.azure.account.oauth2.client.endpoint": "https://login.microsoftonline.com/ea0d9016-05e2-4f0b-9a86-21eba11a11d5/oauth2/token",
         "fs.azure.createRemoteFileSystemDuringInitialization": "true"}

dbutils.fs.mount(
        source = "abfss://labiot89@labiot89.dfs.core.windows.net",
        mount_point = "/mnt/labs",
        extra_configs = configs)
```

```
display(dbutils.fs.ls("/mnt/labs"))
```

![ssh-hd-insight](../8-9/img/8.png)

16. Go to your Python notebook from Workspace and paste:

* (NOTE! Your eventhub instance should not be empty so generate some files)
```
import org.apache.spark.eventhubs.{ ConnectionStringBuilder, EventHubsConf, EventPosition }
import org.apache.spark.sql.types._
import org.apache.spark.sql.functions._


val connectionString = ConnectionStringBuilder("Endpoint=sb://labiot89.servicebus.windows.net/;SharedAccessKeyName=labiot89;SharedAccessKey=TcewBTe6SIVgILPsX5g9EUtR0Nr0z94N+KEk5FkIYMc=;EntityPath=labiot89")
  .setEventHubName("labiot89")
  .build
val eventHubsConf = EventHubsConf(connectionString)
  .setStartingPosition(EventPosition.fromEndOfStream)

var dataset = 
  spark.readStream
    .format("eventhubs")
    .options(eventHubsConf.toMap)
    .load()
      
val filtered = dataset.select(
    from_unixtime(col("enqueuedTime").cast(LongType)).alias("enqueuedTime")
      , get_json_object(col("body").cast(StringType), "$.date").alias("date")
      , get_json_object(col("body").cast(StringType), "$.string").alias("string")
      , get_json_object(col("body").cast(StringType), "$.text").alias("text")
      , get_json_object(col("body").cast(StringType), "$.id").alias("id")
        
  )
  
filtered.writeStream
  .format("com.databricks.spark.csv")
  .outputMode("append")
  .option("checkpointLocation", "/mnt/labs/labiot89")
  .start("/mnt/labs/labiot89")
  ```

  ![ssh-hd-insight](../8-9/img/9.png)


17. In my case I have used those data 

  ![ssh-hd-insight](../8-9/img/10.png)

18. Result a

  ![ssh-hd-insight](../8-9/img/11.png)
