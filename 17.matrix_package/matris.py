from Matris.matrisOperations import MatrisOperations,RandomMatris,Length

class Matris():
    def __init__(self, data=RandomMatris((3,3))) -> None:
        self.data = data

    def __add__(self, otherMatris):
        return Matris(MatrisOperations.add(self.data, otherMatris.data))

    def __sub__(self, otherMatris):
        return Matris(MatrisOperations.extract2M(self.data, otherMatris.data))

    def __mul__(self, otherMatris):
        if type(otherMatris) != Matris:
            return Matris(MatrisOperations.scalerMulti(self.data, otherMatris))
        elif type(otherMatris) == Matris:
            return Matris(MatrisOperations.multi2M(self.data, otherMatris.data))

    def __truediv__(self, otherMatris):
        if type(otherMatris) == int or type(otherMatris) == float :
            
            return Matris(MatrisOperations.scalerMulti(self.data, 1/otherMatris))
        else:
            return "You cannot do that"

    def __repr__(self) -> str:
        return str(self.data)

    def Length(self):
        return Length(self.data)

    @staticmethod
    def Unit(length:int):
        if type(length) != int: return "Length should be integer"
        return Matris([[(1 if x == y else 0)for x in range(length)] for y in range(length)] )



    # def __call__(self):
    #     return self.data


'''
+	object.__add__(self, other)
-	object.__sub__(self, other)
*	object.__mul__(self, other)
//	object.__floordiv__(self, other)
/	object.__truediv__(self, other)
%	object.__mod__(self, other)
**	object.__pow__(self, other[, modulo])
<<	object.__lshift__(self, other)
>>	object.__rshift__(self, other)
&	object.__and__(self, other)
^	object.__xor__(self, other)
|	object.__or__(self, other)
'''