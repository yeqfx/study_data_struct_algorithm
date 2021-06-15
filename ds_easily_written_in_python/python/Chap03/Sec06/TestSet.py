import os
import sys
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))

from dataStruct import Sets

setA = Sets.Set()
setA.insert('휴대폰')
setA.insert('지갑')
setA.insert('손수건')
setA.display('Set A :')

setB = Sets.Set()
setB.insert('빗')
setB.insert('파이썬 자료구조')
setB.insert('야구공')
setB.insert('지갑')
setB.display('Set B :')

setB.insert('빗')
setA.delete('손수건')
setA.delete('발수건')
setA.display('Set A :')
setB.display('Set B :')

setA.union(setB).display('A U B :')
setA.intersect(setB).display('A ^ B :')
setA.difference(setB).display('A - B :')
