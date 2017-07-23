class Difference:
    def __init__(self, a):
        self.__elements = a

    maximumDifference = 0

    def computeDifference(self):
        for k, v in enumerate(self.__elements):
            for l, w in enumerate(self.__elements):
                if k == l:
                    continue
                if abs(v - w) > self.maximumDifference:
                    self.maximumDifference = abs(v - w)


_ = raw_input()
a = [int(e) for e in raw_input().split(' ')]

d = Difference(a)
d.computeDifference()

print d.maximumDifference
