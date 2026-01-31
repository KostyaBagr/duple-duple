
<h3>
CLI Utility for backup and restore data in DBMS. Written on GO. <hr>


Duple-duple is an infterface tool between your DBMS and external storage services such as S3, GoogleDrive, etc. Moreover, it provides notifications with all statistics when the dump is ready!

</h3>


<h2>Currently supported features:</h2>

<h3><b>DBMS</b>
    
    - postgres 

    - mysql (no yet. Will be the next)

<b>Extrnal storage services </b>

    - S3: you can connect any S3 provider you want 

    - Google drive (not yet. Will be the next)

<b>Notifications</b>

    - SMTP

</h3>

<h2>Installation</h2>

```git clone https://github.com/KostyaBagr/duple-duple.git```



<h2>Configuration</h2>

Create new file <i>config.toml</i> (see config.toml.example)

    
    <!-- required -->
        [dbms]
        [dbms.postgres]
        host="DB HOST"
        user="DB USER"
        password="DB USER PASSWORD"
        db="PUT DB NAME HERE OR * TO DUMP A CLUSTER"
        port="DB PORT"

        <!-- required -->
        [storage]
        <!-- required -->
        [storage.s3]
        url="S3 PROVIER URL"
        backetName="S3 BACKET NAME"
        accessKey="S3 ACCESS KEY"
        secretAccessKey="S3 SECRET ACCESS KEY"
        region="S3 REGION"
        pathInBucket="just a dir in your s3. optional"

        <!-- optional -->
        [storage.local]
        path="PATH. STARTS AND ENDS WITH /"

        <!-- optional -->
        [notifications]
        [notifications.email]
        smtpServer="SMTP SERVER"
        sender="SMTP SENDER (sender@mysmtp.com)"
        port=INT
        password="SMTP SENDER PASSWORD"
        receiver="YOUR EMAIL TO GET DATA"
</code>


<h3>

- <p>Duple-duple uses pg_dump and pg_dumpall for postgres, so you can define * (to dump a cluster) or <i>db_name</i> in db row configuration </p>
<p>All files are compressed<p>

- <p> For now there are 2 types of storage locations: local and S3. You can define both or only one but not none. During dump duple-duple creates temporary file in directory <i>/tmp/duple-duple/backup/</i> for all dump files and delete them later but if you define [storage.local] all dumps will be saved in your location without delete.

- <p>If you full out notifications config after each successful dump you will receive an email with statistics. See example below<p>

<p> 
<br>
DBMS: postgres; <br>
Start time: 2026-02-01 00:41:16; <br>
End time: 2026-02-01 00:41:30; <br>
File path: /mnt/myPath/2026-02-01T00:41:16+05:00.tar.gz; <br>
File size: 59.18434 mb; <br>
Storagies: [local S3]; <br>

<p>

</h3>


<h2>Build your tool<h2>

<h3>Sooo it almost done! First of all, you need to be sure that you have installed golang (1.25+) and your DBMS (for example Postgres)</h3>

<h3> 
<p>Steps: </p> <br> 


<p> 1) Build </p>

<code>go build .</code>

<p> As a result you will get a binary file </p>


<p> 2) Run <p>

<code>./duple-duple backup --storage "local,S3" --dbms postgres </code>

OR


<code>./duple-duple backup --storage "S3" --dbms postgres </code>

</h3>

