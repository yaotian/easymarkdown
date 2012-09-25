Oozie HOWTO
==============

This is a note of how to setup oozie working with hadoop.

To simplify the setup, we will use just a single node hadoop. It is just straight forward to use a distributed hadoop cluster.

Software
--------------

	- hadoop 0.20.205.0
	- oozie  3.2.0
	- extjs  2.2
	- maven  3.0.4

Download:

	- wget http://extjs.com/deploy/ext-2.2.zip
	- wget http://www.globalish.com/am/hadoop/common/hadoop-0.20.205.0/hadoop-0.20.205.0.tar.gz 
	- wget http://mirrors.gigenet.com/apache/incubator/oozie/stable/oozie-3.2.0-incubating.tar.gz
	- wget http://linux-files.com/maven/binaries/apache-maven-3.0.4-bin.tar.gz
	
Hadoop Setup
--------------
Add a new user 'hadoop'.

```
$ sudo useradd -m -d /data/homes/hadoop -s /bin/bash -g hadoop hadoop
```

Configure ssh.

```
$ sudo su - hadoop     
$ cd ~     
$ mkdir .ssh
$ chmod go-rwx .ssh
$ cd .ssh
$ ssh-keygen -t rsa
$ cat id_rsa.pub >> authorized_keys
```

Unzip hadoop package to {HADOOP_HOME} and chown.

```
$ sudo chown -R :hadoop {HADOOP_HOME}
$ sudo chmod -R g+rw {HADOOP_HOME}
```

Edit hadoop config files:

$HADOOP_HOME/conf/mapred-site.xml

```
<property>
  <name>mapred.job.tracker</name>
  <value>localhost:9001</value>
</property>
<property>
  <name>mapred.job.tracker.http.address</name>
  <value>0.0.0.0:9003</value>
</property>
```

$HADOOP_HOME/conf/hdfs-site.xml

```
<property>
  <name>fs.default.name</name>
  <value>hdfs://localhost:9000</value>
</property>
<property>
  <name>dfs.http.address</name>
  <value>0.0.0.0:9002</value>
</property>
<property>
  <name>hadoop.tmp.dir</name>
  <value>/tmp/hadoop-${user.name}</value>
</property>
```

$HADOOP_HOME/conf/hadoop-env.sh, add:

```
export JAVA_HOME=/usr/lib/jvm/java-6-sun/
```

Format namenode (first time only).

```
hadoop@HADOOP_HOME$ bin/hadoop namenode -format
```

Start hadoop.

```
hadoop@HADOOP_HOME$ bin/start-all.sh
```

Check status.

	- open http://{HOST}:9002 - HDFS status
	- open http://{HOST}:9003 - JobTracker status
	
Oozie build
---------------
Unzip oozie to {OOZIE_SRC}.

Modify {OOZIE_SRC}/docs/pom.xml, change the <version> of org.apache.maven.doxia and org.apache.maven.doxia to 1.0-alpha-9

Set proxy for maven (if behind a firewall). Edit ~/.m2/settings.xml

```
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0
    http://maven.apache.org/xsd/settings-1.0.0.xsd">

    <proxies>
        <proxy>
            <id>default</id>
            <active>true</active>
            <host>{PROXY_HOST}</host>
            <port>{PROXY_PORT}</port>
        </proxy>
    </proxies>

</settings>
```

Set proxy for maven-ant. Edit {OOZIE_SRC}/bin/mkdistro.sh , add "-Dhttp.proxyHost=proxy -Dhttp.proxyPort=8080" to MVN_OPTS

Execute maven to build.

```
$ {MAVEN_HOME}/bin/mvn package
```

You should get {OOZIE_SRC}/distro/target. Set it as {OOZIE_HOME}.

Oozie Setup
-----------------
Fix hadoop security. Edit {HADOOP_HOME}/conf/core-site.xml, add:

```
<property>
  <name>hadoop.proxyuser.hadoop.hosts</name>
  <value>*</value>
</property>
<property>
  <name>hadoop.proxyuser.hadoop.groups</name>
  <value>*</value>
</property>
```

Restart hadoop.

Edit {OOZIE_HOME}/conf/oozie-site.xml, add:

```
<property>
<name>oozie.services.ext</name>
<value>
org.apache.oozie.service.HadoopAccessorService
</value>
<description>
To add/replace services defined in 'oozie.services' with custom implementations.Class names must be separated by commas.
</description>
</property>
```

Edit {OOZIE_HOME}/conf/oozie-env.sh

```
export OOZIE_HTTP_HOSTNAME=107.20.246.96
export OOZIE_HTTP_PORT=9010
```

Setup oozie.

```
$ {OOZIE_HOME}/bin/oozie-setup.sh -hadoop 0.20.200 {HADOOP_HOME} -extjs {/PATH/TO}/ext-2.2.zip
```

Start oozie.

```
$ {OOZIE_HOME}/bin/start-oozie.sh
```

Check status.
	
	- open http://{HOST}:9010 - oozie web console
	
Run example job
--------------------
Unzip example

```
$ tar xvfz {OOZIE_HOME}/oozie-examples.tar.gz
```

Edit examples/apps/map-reduce/job.properties:

```
nameNode=hdfs://localhost:9000
jobTracker=localhost:9001
```

Put job to HDFS:

```
hadoop fs -mkdir user/hadoop
hadoop fs -put examples hdfs://localhost:9000/user/hadoop/examples
hadoop fs -ls  hdfs://localhost:9000/user/hadoop/examples
```

Run job:

```
bin/oozie job -oozie http://localhost:9010/oozie -config examples/apps/map-reduce/job.properties -run
```

Run self-coded job
------------------
Download the source codes here and use maven to package it.

Deploy it to HDFS:

```
hadoop fs -mkdir hdfs://localhost:9000/user/hadoop/app1/
hadoop fs -mkdir  hdfs://localhost:9000/user/hadoop/app1/lib
hadoop fs -mkdir  hdfs://localhost:9000/user/hadoop/app1/input-data

hadoop fs -put data.txt hdfs://localhost:9000/user/hadoop/app1/input-data
hadoop fs -put target/wordcount-0.1.jar hdfs://localhost:9000/user/hadoop/app1/lib

hadoop fs -put src/workflow.xml hdfs://localhost:9000/user/hadoop/app1
```

Run the job:

```
{OOZIE_HOME}/bin/oozie job -oozie http://localhost:9010/oozie -config src/job.properties -run
```

Check the result:

```
hadoop fs -cat hdfs://localhost:9000/user/hadoop/app1/output-data/word-count/part-00000
```

Please notice that job.properties is not needed to be put to HDFS.

More workflow example please refere to https://github.com/yahoo/oozie/wiki/Oozie-WF-use-cases

Oozie REST API
------------------
It is very simple and easy to use oozie's REST API.

Reference: http://incubator.apache.org/oozie/docs/3.1.3/docs/WebServicesAPI.html

For example:

```
curl http://localhost:9010/oozie/versions
curl http://localhost:9010/oozie/v1/admin/status
curl http://localhost:9010/oozie/v1/job/0000009-120820193550619-oozie-hado-W?show=definition
curl http://localhost:9010/oozie/v1/job/0000009-120820193550619-oozie-hado-W?show=info
```

ENJOY!
