### Go’s standard image packages, will Lissajous figures
```
go run main.go > anim.gif
```
- After importing a package whose path has multiple components, like image/color, we refer to the package with a name that comes from the last component. Thus tee variable color.White belongs to the image/color package and gif.GIF belongs to image/gif.

- Exercise 1.5: 
Change the Lissajous program’s color palette to green on black, for added authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xff}, where each pair of hexadecimal dig its represents the intensity of the red, green, or blue component of the pixel.

- Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding more values to palette and then displaying them by changing the third argument of Set- ColorIndex in some interesting way.
