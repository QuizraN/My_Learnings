def Transpose(inputMatrix):
    print("Input Matrix:")
    [print(i,"\n") for i in inputMatrix]

    result=[]
    [result.append(inputMatrix[j][i]) for i in range(len(inputMatrix)) for j in range(len(inputMatrix[0]))]

    rowCount=0
    outputMatrix=[]
    outputMatrixRow=[]
    for i in result:
        rowCount+=1
        outputMatrixRow.append(i)
        if(rowCount%(len(inputMatrix))==0 and rowCount!=0):
            outputMatrix.append(outputMatrixRow)
            outputMatrixRow=[]
            rowCount=0
    
    print("Transpose of Input Matrix:")
    [print(i,"\n") for i in outputMatrix]

if __name__=='__main__':
    inputMatrix=[[1,2,3],[4,5,6],[7,8,9]]
    Transpose(inputMatrix)


# for i in range(len(inputMatrix)):
#     for j in range(len(inputMatrix[0])):
#         if(i<j):
#             inputMatrix[i][j],inputMatrix[j][i]=inputMatrix[j][i],inputMatrix[i][j]
