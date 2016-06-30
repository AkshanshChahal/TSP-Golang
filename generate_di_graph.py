#!/usr/bin/python

import random
import argparse
import sys

# Function to parse command-line args
def get_args(): 
    parser = argparse.ArgumentParser()
    parser.add_argument('-n', '--nodes', type=str, help="Number of nodes", required=True)
    parser.add_argument('-o', '--ofile', type=str, help="O/p File name", required=True)
    args = parser.parse_args()
    numnodes = args.nodes
    ofile = args.ofile
    return numnodes, ofile
    

# function to create and fill-up the file 
def create_file():
    numnodes, ofile = get_args()
    print 'Numer of nodes: ', numnodes
    print 'Outfile file: ', ofile
    try: 
        #open file
        file = open(ofile, 'w+')
        
        #write into file
        file.write(str(numnodes) +  ' ')
        for i in range(int(numnodes)):
            for j in range(int(numnodes)):
                if i == j:
                    file.write("0" + ' ')
                else:
                    file.write(str(random.randint(1,100)) + ' ')

        #close file
        file.close()
    except IOError as e: 
        print "Error error({0}): {1}", format(e.errno, e.strerror)
    except: 
        print "Unexpected error", sys.exc_info()[0]
        raise

    return ofile

ofile = create_file()

with open(ofile, 'r') as fin:
    print fin.read()
