class Person():
    def __init__(self, firstName, lastName, idNumber):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNumber

    def printPerson(self):
        print "Name:", self.lastName + ",", self.firstName
        print "ID:", self.idNumber


from collections import OrderedDict
GRADING = OrderedDict(
    [
        (90,  "O"),
        (80, "E"),
        (70, "A"),
        (55, "P"),
        (40, "D"),
        (0, "T"),
    ],
)

class Student(Person):

    def __init__(self, firstName, lastName, idNumber, scores):
        Person.__init__(self, firstName, lastName, idNumber)
        self.scores = scores

    def calculate(self):
        s = sum(self.scores) / len(self.scores)
        for k, v in GRADING.iteritems():
            if s >= k:
                return v


line = raw_input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(raw_input()) # not needed for Python
scores = map(int, raw_input().split())
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print "Grade:", s.calculate()
