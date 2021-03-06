package hexmaker

import (
  "github.com/fogleman/gg"
  "github.com/golang/freetype/truetype"
  "golang.org/x/image/font/gofont/goregular"
  "math"
  "strconv"
  "fmt"
  "os"
  "log"
  "errors"
)

func MakeHexagon(PP [][]int, color bool) {
  const S = 200
  dc := gg.NewContext(S, S)
  radius := float64(S/2)

  // Set up font
  font, err := truetype.Parse(goregular.TTF)
  if err != nil {
    panic("")
  }
  face := truetype.NewFace(font, &truetype.Options{
    Size: 12,
  })
  dc.SetFontFace(face)

  // Draw hexagon
  dc.DrawRegularPolygon(6, S/2, S/2, radius, 0) // sides, x, y, radius, rotation
  if color {
    dc.SetHexColor("016d9a")
    dc.Fill()
  } else {
    dc.SetRGB(1, 1, 1)
    dc.Fill()
    dc.SetRGB(0,0,0)
    dc.DrawRegularPolygon(6, S/2, S/2, radius, 0) // sides, x, y, radius, rotation
    dc.Stroke()
  }

  // Draw circles and text
for j := 0; j < 31; j++ {
    dc.Push()
    // Card number
    cardno := strconv.Itoa(j+1)
    if color {
      dc.SetHexColor("a61313")
      dc.DrawCircle(S/2, S/2, S/10)
      dc.Fill()
      dc.SetRGB(1, 1, 1)
      dc.DrawStringAnchored(cardno, S/2, S/2, 0.5, 0.5)
    } else {
      dc.SetRGB(1,1,1)
      dc.DrawCircle(S/2, S/2, S/10)
      dc.Fill()
      dc.SetRGB(0,0,0)
      dc.DrawStringAnchored(cardno, S/2, S/2, 0.5, 0.5)
    }
    dc.Pop()
  for i := 0; i < 6; i++ {
    dc.Push()

    // Connection points
    angle := gg.Radians(float64(i*60 - 30)) // -30 for edge-alignment
    x := S/2 + 0.65*radius*math.Cos(angle) 
    y := S/2 + 0.65*radius*math.Sin(angle)
    dc.RotateAbout(angle, x, y)
    if color {
      dc.SetHexColor("e42828")
      dc.DrawCircle(x, y, S/10)
      dc.Fill()
      dc.SetRGB(1, 1, 1)
      dc.RotateAbout(gg.Radians(270), x, y)
      text := strconv.Itoa(PP[j][i])
      dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
    } else {
      dc.SetRGB(1,1,1)
      dc.DrawCircle(x, y, S/10)
      dc.Fill()
      dc.SetRGB(0,0,0)
      dc.RotateAbout(gg.Radians(270), x, y)
      text := strconv.Itoa(PP[j][i])
      dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
    }
    dc.Pop()
  }
  // Write the current card
  if _, err := os.Stat("deck"); errors.Is(err, os.ErrNotExist) {
    err := os.Mkdir("deck", os.ModePerm)
    if err != nil {
      log.Println(err)
    }
  }
  filename := fmt.Sprintf("deck/hex%d.png", j+1)
  dc.SavePNG(filename)
  }
}

