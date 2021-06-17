import os
import sys

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))
from dataStruct import ArrayStack

def evalPostfix(expr):
    s = ArrayStack.ArrayStack()
    for token in expr:
        if token in "+-*/":
            val2 = s.pop()
            val1 = s.pop()
            if (token == '+'): s.push(val1 + val2)
            elif (token == '-'): s.push(val1 - val2)
            elif (token == '*'): s.push(val1 * val2)
            elif (token == '/'): s.push(val1 / val2)
        else: s.push(float(token))
    return s.pop()

def Infix2Postfix(expr):
    s = ArrayStack.ArrayStack()
    output = []
    for term in expr:
        if term in '(':
            s.push('(')
        elif term in ')':
            while not s.isEmpty():
                op = s.pop()
                if op=='(' : break
                else: output.append(op)
        elif term in "+-*/":
            while not s.isEmpty():
                op = s.peek()
                if (precedence(term) <= precedence(op)):
                    output.append(op)
                    s.pop()
                else: break
            s.push(term)
        else:
            output.append(term)

    while not s.isEmpty():
        output.append(s.pop())

    return output

def precedence(op):
    if op=='(' or op==')': return 0
    elif op=='+' or op=='-': return 1
    elif op=='*' or op=='/': return 2
    else: return -1

if __name__ == "__main__":
    expr1 = [ '8', '2', '/', '3', '-', '3', '2', '*', '+' ]
    expr2 = [ '1', '2', '/', '4', '*', '1', '4', '/', '*' ]
    print(expr1, ' --> ', evalPostfix(expr1))
    print(expr2, ' --> ', evalPostfix(expr2))
    print()
    infix1 = [ '8', '/', '2', '-', '3', '+', '(', '3', '*', '2', ')' ]
    infix2 = [ '1', '/', '2', '*', '4', '*', '(', '1', '/', '4', ')' ]
    postfix1 = Infix2Postfix(infix1)
    postfix2 = Infix2Postfix(infix2)
    result1 = evalPostfix(postfix1)
    result2 = evalPostfix(postfix2)
    print(' 중위표기 : ', infix1)
    print(' 후위표기 : ', postfix1)
    print(' 계산결과 : ', result1)
    print()
    print(' 중위표기 : ', infix2)
    print(' 후위표기 : ', postfix2)
    print(' 계산결과 : ', result2)