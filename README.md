 ## One reply trivial DNS server
 This is an simple example for my students about creating network services
 at Go.
 Also I would like to deliver my expressions about Go, compare it with
 Erlang.

 It can reply with speed about 112922 qps on my old notebook.

 # Compilation:

 Just type
 $make 

 # Tests:

 Run it as root and test resolving by "nslookup 1000.dip 127.0.0.1" or "kdig 1000.dip @127.0.0.1" 
 and then do benchmarks by dnsperf or something else

