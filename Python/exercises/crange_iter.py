import time

class custom_range:
    def __init__(self,start,stop,step=1):
        self.start=start
        self.stop=stop
        self.step=step
        self.count=start

    def __iter__(self):
        return self

    def __next__(self):
        if(self.count<self.stop):
            value=self.count
            self.count += self.step
            return value
        else:
            raise StopIteration()


def custom_range2(start=0,stop=0,step=1):
    count=start
    result=[]
    while count!=stop:
        result.append(count)
        count+=1
    
    iterable=iter(result)
    return iterable

if __name__=='__main__':

    start_time = time.time()
    for index in range(2,5):
        print(index)
    print("range function took", time.time() - start_time, "to run")

    start_time = time.time()
    for index in custom_range(2,5):
        print(index)
    print("custom_range function took", time.time() - start_time, "to run")
   
    start_time = time.time()
    for index in custom_range2(2,5):
        print(index)
    print("custom_range function took", time.time() - start_time, "to run")




