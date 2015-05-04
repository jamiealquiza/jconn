### Overview

jconn prints TCP connection information as JSON.

### Install

Assuming Go is installed:

 - `go get github.com/jamiealquiza/jconn`
 - `go build github.com/jamiealquiza/jconn`

### Usage

<pre>
% jconn -h
Usage of jconn:
  -tag-hostname=false: Add @hostname: hostname field-value
  -tag-rto=true: Add RTO field if applicable
  -tag-timestamp=false: Add @timestamp: RFC3339 field-value
</pre>

### Examples

Inbound SSH connections:
<pre>
% jconn -tag-rto=false | grep '"local_port":"22"' | grep ESTAB
{"local_ip":"192.168.100.25","local_port":"22","remote_ip":"192.168.100.1","remote_port":"36880","state":"ESTABLISHED"}
</pre>

All listening with hostname and timestamp fields appended:
<pre>
% jconn -tag-timestamp -tag-hostname | grep LIST
{"local_ip":"0.0.0.0","local_port":"50008","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"1560","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"53049","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"53316","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"32400","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"32401","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"32469","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
{"local_ip":"0.0.0.0","local_port":"22","remote_ip":"0.0.0.0","remote_port":"0","state":"LISTEN","@timestamp":"2015-04-06T10:15:16-06:00","@hostname":"plex"}
</pre>
