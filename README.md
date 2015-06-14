# goghost
Go bindings for the Ghostscript interpreter API

==== Build instructions ====


Download the Ghostscript 9.15 source from:
http://www.ghostscript.com/download/gsdnld.html

Copy the iapi.h and ierrors.h files from ghostscript-9.15\psi to the go-ghost directory.

Install or build Ghostscript 9.15. Copy the ghostscript dll file gsdll64.dll to the go-ghost directory.

For the binary install, the dll path is:
C:\Program Files\gs\gs9.15\bin\gsdll64.dll

gsdll64.dll must also be copied to the location of your compiled binary that uses goghost.

Install a mingw compiler:
http://sourceforge.net/projects/mingw-w64
Make sure gcc is on your path
