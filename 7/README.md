# 7

1. Go to Kibana -> Visualise -> Create visualization

2. Select "Pie" chart and next choose your source

3. Add a bucket with "split slice" settings and aggregation "terms"

* 3.1 add your field 
* 3.2 customize your chart
* 3.3 save

![ssh-hd-insight](../7/img/1.png)


4.  Crate new visualization as in step 1 
 
5. Select "Line" chart and next choose your source

6. Add a bucket with "X-axis" settings and aggregation "terms"

* 6.1 add your field 
* 6.2 customize your chart
* 6.3 save

![ssh-hd-insight](../7/img/2.png)


7. Crate new visualization as in step 1  

8. Select the "Controls" chart and next choose your source

9. Select the option list and choose your source in the "Index pattern" dropdown

* 9.1 add your field 
* 9.2 customize selection
* 9.3 save

![ssh-hd-insight](../7/img/3.png)


10. Go to Kibana -> Dashboards -> Create dashboard 

11. Add created charts

![ssh-hd-insight](../7/img/4.png)

12. Create "Get" request to score 5 results with keyword "department"

![ssh-hd-insight](../7/img/5.png)

13. Create "Get" request to score 1 "bfy" result wich has "2020" inside 

![ssh-hd-insight](../7/img/6.png)
