# File Generator

[![Build Status](https://cloud.drone.io/api/badges/khmarbaise/filegenerator/status.svg)](https://cloud.drone.io/khmarbaise/filegenerator)


# Status

The current state of development is `Prototype`.

# Text Command

The current `text` command (`fgen text`) generates by default
a file `generated.txt` which looks like this:
```
0 2Dj5rCebweHP9ZBYihU3 bfpvWolDG8sVQFNh2ST9 8wPKZ2uB2xO3asYRTEim fvknaOZMRiwrTnXHO19S e4HrIkXCP6wTLWT8hJm3 r5C2wDPOhmmx5cCCvFPO yWdVDUvottUeuhRxXX5K oEK67e1xbXvpJxz4NEXz xEgT75h7weAjRHdDEVcS tRbjYfbHMjXDzHNv1Oxt 
1 HjZIm1HE2B0ASZ24dBhA qOBXLFd0qDoNkFqcQokI 39l2nc04MQOhz0QFxlHr MCsWWpduewmqE6lHqYJ8 m96ow4z1FBVCzM0PfLnX 9mHC3MCOZWqsBZv1puCY NKnOSANW8zHpDlX05xCz bzO0lMpM3t3K6Tr5zlTJ rjj58Fg6atrmRksJxobf 6xoJCVd9fwg61kZyY1R8 
...
9999 uS0CZxL9WCxi8iRhBJlv VWGAkWSI0fgMx1ZjlpKz nuXXTgCWpkGnxsmZwB4k kT1aVi96YLD2B4ljyLXa TW4p6vf2UVlB33PLjjMM ho74Ua8JaOy0pdhEVU9s MTByNMVK4cDFNJOZthxq vCyJsL4hnSYM1EhJCvDZ dClCHPEk51Pyu2MZOg7W MKOD2bcZItfTYyUdWydt 
```
The number at the beginning is the line number (starting from 0). 
Each block is a string containing characters (`a-zA-Z0-9`) generated based
on a random generator. The length of the string can be changed by using the command
line option `--size Number` where as the number of entries of strings per line
can be changed by using `--entries Number` and finally the number of lines
which are generated can be changed by using `--lines Nr.`.
