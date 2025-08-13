from Matris.auxiliaryFuncs import *

'''
total = [[(1 if x == y else 0)for x in range(row)] for y in range(col)] 
yukarıdaki ile kare matrislerin standart taban vektörlerini boyuta göre oto alıyoruz
'''

class MatrisOperations:
    @staticmethod
    def multi2M(matrisA, matrisB):
        if ValidMatris(matrisA) == True and ValidMatris(matrisB) == True:
            if canBeMultiply(matrisA, matrisB)[0] == True:
                newRow, newCol = canBeMultiply(matrisA, matrisB)[1]
                newMatris = [[0 for i in range(newCol)] for j in range(newRow)]
                
                a,b = Length(matrisB)
                for i in range(0,newRow,1):
                    for j in range(0,b,1):
                        for k in range(0,a,1):
                            newMatris[i][j] += matrisA[i][k] * matrisB[k][j]

                return newMatris
                

    @staticmethod
    def add(*matrisList):
        if validMatrisList(matrisList) == True:
            if canBeAdded(matrisList) == True:
                row, col = Length(matrisList[0])
                total = [[0 for x in range(col)] for y in range(row)] 

                for i, row in enumerate(matrisList):
                    matris = matrisList[i]
                    for j,row in enumerate(matris):
                        for t,value in enumerate(row):
                            total[j][t] += matris[j][t]

                return total
        return "Enter Valid Matris"

    @staticmethod
    def extract2M(matrisA, matrisB):
        if Length(matrisA) == Length(matrisB):
            row, col = Length(matrisA)
            total = [[0 for x in range(row)] for y in range(col)] 
            
            for i,row in enumerate(matrisA):
                for j,value in enumerate(row):
                    total[i][j] += matrisA[i][j] - matrisB[i][j]

            return total

    @staticmethod
    def scalerMulti(matris, scaler):
        valid = ValidMatris(matris)
        newMatris = matris
        
        if valid == True:

            for i,row in enumerate(newMatris):
                for j, value in enumerate(row):
                    newMatris[i][j] *= scaler

            return newMatris
        else:
            return "Please Enter a Valid Matris"

