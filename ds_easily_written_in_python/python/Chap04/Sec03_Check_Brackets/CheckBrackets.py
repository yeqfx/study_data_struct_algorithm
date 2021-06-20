import os
import sys

# sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))))

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))
# from utilities import checkBrackets
from myutil import checkBrackets

if __name__ == "__main__":
    
    infile = open(__file__, "r")
    lines = infile.readlines()
    infile.close()    

    result = checkBrackets.checkBracketsV2(lines)
    print(os.path.basename(__file__), " ---> ", result)
