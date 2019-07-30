class Solution:
    INT_MAX = 2147483647
    INT_MIN = -2147483648

    def isNum(self, a) -> bool:
        if a >= '0' and a <= '9':
            return True
        else:
            return False

    def myAtoi(self, str: str) -> int:
        result = 0
        temp = ""
        str = str.strip()
        flag = True
        hasSign = False

        for i in str:
            if not temp and not hasSign and (i == '+' or i == '-'):
                hasSign = True
                if i == '-':
                    flag = False
                continue
            else:
                if self.isNum(i):
                    temp = temp + i
                else:
                    break

        if not temp:
            return result
        else:
            result = int(temp)


        if not flag:
            result = -result
        if result > self.INT_MAX:
            result = self.INT_MAX
        if result < self.INT_MIN:
            result = self.INT_MIN

        return result
