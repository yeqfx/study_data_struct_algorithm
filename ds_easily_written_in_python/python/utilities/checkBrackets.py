import os
import sys
sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))

from dataStruct import ArrayStack

def checkBrackets(statement):
    stack = ArrayStack.ArrayStack()
    for ch in statement:
        if ch in ('{[('):
            stack.push(ch)
        elif ch in ('}])'):
            if stack.isEmpty():
                return False
            else:
                left = stack.pop()
                if (ch == "}" and left != "{") or \
                    (ch == "]" and left != "[") or \
                    (ch == ")" and left != "("):
                    return False
    return stack.isEmpty()

def checkBracketsV2(lines):
    stack = ArrayStack.ArrayStack()
    for line in lines:
        for ch in line:
            if ch in ('{[('):
                stack.push(ch)
            elif ch in ('}])'):
                if stack.isEmpty():
                    return False
                else:
                    left = stack.pop()
                    if (ch == "}" and left != "{") or \
                        (ch == "]" and left != "[") or \
                        (ch == ")" and left != "("):
                        return False
    return stack.isEmpty()
 

