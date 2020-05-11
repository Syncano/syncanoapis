from distutils.core import setup
import os


def find_packages(path='.'):
    ret = []
    for root, dirs, files in os.walk(path):
        if files and '__pycache__' not in root:
            ret.append(root)
    return ret


setup(
    name='syncanoapis',
    version='0.1',
    packages=list(find_packages('syncano')),
)
