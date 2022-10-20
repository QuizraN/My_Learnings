
def FibonacciSeries(count=0,index=0):

    indexCount=0
    firstValue=True
    prev=0
    current=1

    if count<1:
        raise StopIteration

    while(index!=indexCount):
        indexCount+=1
        value=current+prev
        prev=current
        current=value

    if firstValue:
        firstValue=False
        count-=1
        yield current

    while(count!=0):
        value=current+prev
        prev=current
        current=value
        count-=1
        yield value

# import ipdb;ipdb.set_trace()
if __name__=='__main__':
    [print(index) for index in FibonacciSeries(5)]
    print("*********")
    [print(index) for index in FibonacciSeries(5,3)]
