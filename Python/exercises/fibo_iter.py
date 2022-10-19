class FibonacciSeries:

    def __init__(self,count=0,index=0):
        self.prev=0
        self.current=1
        self.count=count
        self.index=index
        self.firstValue=True
        self.indexCount=0

    def __iter__(self):
        return self

    def __next__(self):
        
        if self.count<1:
            raise StopIteration

        if(self.index!=self.indexCount):
            self.indexCount+=1
            self.count+=1
            self.__next__() 

            #it is getting called for index number of times
            # and hence those values are getting missed as we are not capturing it
        # import ipdb;ipdb.set_trace()

        if self.firstValue:
            self.firstValue=False
            self.count-=1
            return self.current
        
        value=self.current+self.prev
        self.prev=self.current
        self.current=value
        self.count-=1
        return value

import ipdb;ipdb.set_trace()

for i in FibonacciSeries(5):
    print(i)

print("*********")

for i in FibonacciSeries(5,3):
    print(i)



        