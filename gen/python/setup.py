from distutils.core import setup
import re
import os

def find_packages(path='.'):
    ret = []
    for root, dirs, files in os.walk(path):
        if files and '__pycache__' not in root:
            ret.append(root)
    return ret

print(list(find_packages('syncano')))
setup(
    name='syncanoapis',
    version='0.1dev',
    packages=list(find_packages('syncano')),
)
