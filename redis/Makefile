benchmark:
	redis-benchmark -n 100000 -h 192.168.137.1 -p 9009 -t set,get
def:
	redis-benchmark -n 100000 -h 192.168.137.1 -p 6379 -t set,get
local:
	redis-benchmark -n 1000 -h 127.0.0.1 -p 9009 -t set,get