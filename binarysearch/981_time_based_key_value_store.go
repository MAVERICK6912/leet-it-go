package binarysearch

type timestampValue struct {
	value     string
	timestamp int
}

type TimeMap struct {
	store map[string][]timestampValue
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]timestampValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], timestampValue{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res:=""
    if vals,ok:=this.store[key];ok{
        
        left,right:=0,len(vals)-1       
        for left<=right{
            mid:=left+(right-left)/2            
            if vals[mid].timestamp<=timestamp{
                res=vals[mid].value
                left=mid+1                                
            }else{
                right=mid-1
            }
        }        
    }
    return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
