# NSQ Ecosystem 

Boot up:

`docker-compose up -d`

See running container details 

`docker-compose ps`

Results in

```
       Name                      Command               State                                            Ports                                           
-------------------------------------------------------------------------------------------------------------------------------------------------------
broker_nsqadmin_1     /nsqadmin --lookupd-http-a ...   Up      4150/tcp, 4151/tcp, 4160/tcp, 4161/tcp, 4170/tcp, 0.0.0.0:32768->4171/tcp                
broker_nsqd_1         /nsqd --lookupd-tcp-addres ...   Up      0.0.0.0:32770->4150/tcp, 0.0.0.0:32769->4151/tcp, 4160/tcp, 4161/tcp, 4170/tcp, 4171/tcp 
broker_nsqlookupd_1   /nsqlookupd                      Up      4150/tcp, 4151/tcp, 0.0.0.0:32772->4160/tcp, 0.0.0.0:32771->4161/tcp, 4170/tcp, 4171/tcp 
```

Note port mapped in host for nsqlookupd 4161, above is 32771

Test:

`curl http://127.0.0.1:32771/ping`

Should receive `OK`

