hashchk
=======

_hashchk_ is a command line tool for showing md5 or sha1 checksum of files in a directory.  This is a [Go] port of a similar tool I wrote in [Python] mostly as a exercise.


My use case is when I'm writing a script that has file output, I sometimes want to be able to verify that my code changes did not change the program's output.  I rename the output files from the intial run or move them to a different directory and then re-run with my code changes.  Often used together a companion program [version],

Usage
-----
    >hashchk
    No file specifications given as arguments.
    Usage:  hashchk [-h=md5|sha1] files...
    
    This utility will hash all the files given on the command line using the given hash function.
    It will print the results two ways, sorted by hash digest and by file name.
      -h="md5": hash function to use: md5|sha1
      files...: 1 or more file name specifications

Example
-------
    >hashchk *.txt
    by file name
    78e6221f6393d1356681db398f14ce6d -> output.1.txt
    78e6221f6393d1356681db398f14ce6d -> output.2.txt
    e98d2f001da5678b39482efbdf5770dc -> report.1.txt
    30707d3a1d72dd5f19f0cd8d7efa0370 -> report.2.txt


    by file digest
    30707d3a1d72dd5f19f0cd8d7efa0370 -> report.2.txt
    78e6221f6393d1356681db398f14ce6d -> output.2.txt
    78e6221f6393d1356681db398f14ce6d -> output.1.txt
    e98d2f001da5678b39482efbdf5770dc -> report.1.txt


Using the example above you can see that the 2nd run created an identical 'output' file as the 1st run. However, something messed up in the 'report' function because the file digest between run 1 and run 2 is different.  Usually a quick scan by eye is all that is needed.  

Version
----
1.0

License
----
MIT

[go]:http://golang.org/
[python]:http://python.org
[version]:https://github.com/jaywgraves/version