#!/bin/bash

# C call cpp
#g++ -Wall -c person.cpp wrapper.cpp
#gcc -Wall -c hello.c
#g++ -o test *.o
#rm *.o

# Cpp to lib
#g++ -Wall -c person.cpp wrapper.cpp
#ar -rv libhello.a *.o
#rm *.o
#gcc -Wall -c hello.c
#g++ -o test *.o -L. -lhello
#rm *.o

# All to lib
g++ -Wall -c console.cpp fan.cpp handtiles.cpp pack.cpp print.cpp tile.cpp wrapper1.cpp
gcc -Wall -c mahjong.c
#ar -crv libhello.a *.o
ar -crsv libmahjong.a *.o
rm *.o
