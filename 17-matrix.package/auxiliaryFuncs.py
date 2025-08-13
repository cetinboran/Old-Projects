import random as r

def ValidMatris(matris):
    colLengthList = []
    valid = False

    for index,r in enumerate(matris):
        colLengthList.append(len(matris[index]))

    for value in colLengthList:
        if colLengthList[0] == value:
            valid = True
        else: 
            valid = False
            break
    
    if valid == True:
        return True

    return False

def validMatrisList(matrisList):
    checkMatrices = [True for matris in matrisList if ValidMatris(matris) == True]
    valid = True if len(checkMatrices) > 0 else False
    return valid

def canBeAdded(matrisList):
    canBeAdded = False    

    FirstLength = Length(matrisList[0])
    for matris in matrisList:
        if FirstLength == Length(matris):
            canBeAdded = True
        else:
            canBeAdded = False
            break
    
    return canBeAdded

def canBeMultiply(matrisA, matrisB):
    r1, c1 = Length(matrisA)
    r2, c2 = Length(matrisB)
    if c1 == r2 :
        return True,(r1,c2)
    else:
        return False,False

def SquareMatris(matris):
    return True if Length(matris)[0] == Length(matris)[1] else False
    
def Length(matris):
    valid = ValidMatris(matris)

    if valid == True:
        row = len(matris)
        col = len(matris[0])

        return row,col
    
    return "Please Enter Valid Matris"

def RandomMatris(length:tuple):
    row, col = length
    matris = [[r.randint(0,9) for i in range(col)] for j in range(row)]
    return matris