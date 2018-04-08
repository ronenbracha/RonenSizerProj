# RonenSizerProj
Ronen's Sizer task

Resizing Logic: 
The image is scaled down to fill the given width and height while retaining the
original aspect ratio and with all of the original image visible. If the requested
dimensions are bigger than the original image&#39;s, the image doesn’t scale up. If
the proportions of the original image do not match the given width and height,
black padding is added to the image to reach the required size.

This small project is using imaging package (https://github.com/disintegration/imaging)

Installation

set GOPATH env variable to local directory (export GOPATH=`pwd`)
Get RonenSizer:
go get -u github.com/ronenbracha/RonenSizerProj

enter application dir "cd src/github.com/ronenbracha/RonenSizerProj"

run as go app in user root:
sudo go run RonenSizer.go
