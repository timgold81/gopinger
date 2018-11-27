# gopinger
Track host using ping and display time when it went up or down
## Usage
<pre>
Usage of pinger.exe:
  -d int
        How many pings get lost when the host declared as down. (default 3)
  -o int
        Interval between pings when host is down. (default 1)
  -s string
        IP host to track. (default "8.8.8.8")
  -u int
        Interval between pings when host is up. (default 1)
  -v    Verbosity
</pre>
## Example
<pre>
go\pinger>pinger.exe -s 192.168.0.167
Starting pinger. Monitoring host 192.168.0.167
Host 192.168.0.167 is down at Tue, 27 Nov 2018 14:14:09 IST
Host 192.168.0.167 is up at Tue, 27 Nov 2018 14:14:39 IST
</pre>