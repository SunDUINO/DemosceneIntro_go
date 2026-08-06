/*
 * ╔════════════════════════════════════════════════════════════════╗
 * ║ Intro_raylib_1                                                 ║
 * ║ Plik / File: main.go                                           ║
 * ╠════════════════════════════════════════════════════════════════╣
 * ║ Autor / Author:                                                ║
 * ║   SunRiver                                                     ║
 * ║   Lothar TeaM                                                  ║
 * ╠════════════════════════════════════════════════════════════════╣
 * ║ GitHub  : Intro_raylib_1                                       ║
 * ║ WWW     : https://lothar-team.pl                               ║
 * ║ Forum   : https://forum.lothar-team.pl                         ║
 * ║                                                                ║
 * ║ Licencja / License: MIT                                        ║
 * ║ Rok / Year: 2026                                               ║
 * ╚════════════════════════════════════════════════════════════════╝
 */

 // v1.0.0


package main

import (
    "log"
    "math"
    "math/rand"
    "os"
    "runtime/pprof"
    "time"

    "github.com/faiface/beep"
    "github.com/faiface/beep/mp3"
    "github.com/faiface/beep/speaker"
    rl "github.com/gen2brain/raylib-go/raylib"
)

const (
    screenWidth  = 800
    screenHeight = 600
    starCount    = 350
)

type Star struct {
    x, y, z float32
    blink   float32
}

type app struct {
    screenWidth         int32
    screenHeight        int32
    stars               []Star
    timeElapsed         float64
    scrollOffsetBottom  float32
    scrollOffsetTop     float32
    rotationCount       int
    baseFontSize        float32
    scrollingTextBottom string
    scrollingTextTop    string
    line1               string
    line2               string
}

func main() {
    // CPU profiling – START
    cpuFile, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(cpuFile)
    defer pprof.StopCPUProfile()

    // Memory profiling – DEFER na końcu
    memFile, _ := os.Create("mem.prof")
    defer memFile.Close()
    defer pprof.WriteHeapProfile(memFile)

    app := newApp()
    app.run()
}

func newApp() *app {
    a := &app{
        screenWidth:         screenWidth,
        screenHeight:        screenHeight,
        stars:               createStars(),
        baseFontSize:        30,
        scrollingTextBottom: " *** SunRiver przesyla pozdrowienia dla: Nefarious, Xbary, Maras52, L3n1n, Gufim, Foreste i wszystkich, których nie wymienilem !! *** ZAPRASZAMY *** https://forum.lothar-team.pl/ ***",
        scrollingTextTop:    "*** Forum Lothar TEAM --->> ESP8266 * ESP32 * STM32 *** kursy praktyczne jezyków Go, Python, C# i wiele innych <<--- Forum Lothar TEAM ***",
        line1:               "",
        line2:               "",
    }

    a.scrollOffsetBottom = float32(a.screenWidth)
    a.scrollOffsetTop = -float32(rl.MeasureText(a.scrollingTextTop, 20))

    return a
}

func createStars() []Star {
    stars := make([]Star, starCount)
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    for i := 0; i < starCount; i++ {
        stars[i] = Star{
            x:     rng.Float32()*float32(screenWidth) - float32(screenWidth)/2,
            y:     rng.Float32()*float32(screenHeight) - float32(screenHeight)/2,
            z:     rng.Float32() * float32(screenWidth),
            blink: rng.Float32(),
        }
    }

    return stars
}

func (a *app) run() {
    rl.InitWindow(screenWidth, screenHeight, "SunIntro v9 --->>  muzyka nickpanek620 z pixabay <<---")
    defer rl.CloseWindow()

    rl.SetTargetFPS(60)

    musicFile, err := os.Open("chiptune.mp3")
    if err != nil {
        log.Fatal(err)
    }
    defer musicFile.Close()

    streamer, format, err := mp3.Decode(musicFile)
    if err != nil {
        log.Fatal(err)
    }
    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    speaker.Play(beep.Loop(-1, streamer))

    for !rl.WindowShouldClose() {
        a.update()
        a.draw()
        time.Sleep(16 * time.Millisecond)
    }
}

func (a *app) update() {
    a.timeElapsed += 0.05
    a.scrollOffsetBottom -= 2
    a.scrollOffsetTop += 2

    if a.scrollOffsetBottom < -float32(rl.MeasureText(a.scrollingTextBottom, 20)) {
        a.scrollOffsetBottom = float32(a.screenWidth)
    }

    if a.scrollOffsetTop > float32(a.screenWidth) {
        a.scrollOffsetTop = -float32(rl.MeasureText(a.scrollingTextTop, 20))
    }

    for i := 0; i < starCount; i++ {
        a.stars[i].z -= 6

        if a.stars[i].z < 1 {
            a.stars[i].x = rand.Float32()*float32(screenWidth) - float32(screenWidth)/2
            a.stars[i].y = rand.Float32()*float32(screenHeight) - float32(screenHeight)/2
            a.stars[i].z = float32(screenWidth)
            a.stars[i].blink = rand.Float32()
        }

        a.stars[i].blink += rand.Float32() * 0.05
        if a.stars[i].blink > 1 {
            a.stars[i].blink = 0
        }
    }

    a.rotationCount = int(a.timeElapsed / (2 * math.Pi))
}

func (a *app) draw() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)

    a.drawStars()
    a.drawText()
    a.draw3DScene()

    rl.EndDrawing()
}

func (a *app) drawStars() {
    for i := 0; i < starCount; i++ {
        factor := 0.001 * a.stars[i].z
        sx := a.stars[i].x/factor + float32(a.screenWidth)/2
        sy := a.stars[i].y/factor + float32(a.screenHeight)/2
        size := (1 - a.stars[i].z/float32(a.screenWidth)) * 8
        alpha := uint8(255 * a.stars[i].blink)

        rl.DrawCircle(int32(sx), int32(sy), size, rl.NewColor(255, 255, 255, alpha))
    }
}

func (a *app) drawText() {
    fontSize := a.baseFontSize + 10*float32(math.Sin(a.timeElapsed))

    textColor := rl.NewColor(
        uint8(128+127*math.Sin(a.timeElapsed)),
        uint8(128+127*math.Sin(a.timeElapsed+2*math.Pi/3)),
        uint8(128+127*math.Sin(a.timeElapsed+4*math.Pi/3)),
        255,
    )

    line1X := float32(a.screenWidth/2 - rl.MeasureText(a.line1, int32(fontSize))/2)
    line1Y := float32(a.screenHeight/2) - fontSize
    line2X := float32(a.screenWidth/2 - rl.MeasureText(a.line2, int32(fontSize))/2)
    line2Y := float32(a.screenHeight/2) + fontSize

    rl.DrawText(a.line1, int32(line1X), int32(line1Y), int32(fontSize), textColor)
    rl.DrawText(a.line2, int32(line2X), int32(line2Y), int32(fontSize), textColor)

    rl.DrawText(a.scrollingTextBottom, int32(a.scrollOffsetBottom), a.screenHeight-50, 25, textColor)
    rl.DrawText(a.scrollingTextTop, int32(a.scrollOffsetTop), 10, 25, textColor)
}

func (a *app) draw3DScene() {
    camera := rl.Camera{
        Position: rl.NewVector3(0, 0, 300),
        Target:   rl.NewVector3(0, 0, 0),
        Up:       rl.NewVector3(0, 1, 0),
        Fovy:     45.0,
    }

    rl.BeginMode3D(camera)

    rotationSpeed := float32(45)
    rotation := float32(a.timeElapsed) * rotationSpeed

    rl.PushMatrix()
    rl.Translatef(0, 0, 0)
    rl.Rotatef(rotation, 1, 1, 0)

    switch (a.rotationCount / 4) % 9 {
    case 0:
        rl.DrawCubeWires(rl.NewVector3(0, 0, 0), 80, 80, 80, rl.NewColor(255, 0, 0, 255))
    case 1:
        rl.DrawSphereWires(rl.NewVector3(0, 0, 0), 50, 16, 16, rl.NewColor(0, 0, 255, 255))
    case 2:
        drawCone(rl.NewVector3(0, 0, 0), 50, 100, 16, rl.NewColor(0, 255, 0, 255))
    case 3:
        drawPyramid(rl.NewVector3(0, 0, 0), 100, rl.NewColor(255, 255, 0, 255))
    case 4:
        drawDiamond(rl.NewVector3(0, 0, 0), 100, rl.NewColor(255, 0, 255, 255))
    case 5:
        drawCrystal(rl.NewVector3(0, 0, 0), 100, rl.NewColor(255, 255, 255, 255))
    case 6:
        drawBrilliantBasedOnCone(rl.NewVector3(0, 0, 0), 80, 100, 16, rl.NewColor(255, 255, 255, 255))
    case 7:
        drawBrilliant(rl.NewVector3(0, 0, 0), 80, 100, 16, rl.NewColor(255, 255, 255, 255))
    case 8:
        drawBrilliantWithCone(rl.NewVector3(0, 0, 0), 80, 100, 16, rl.NewColor(255, 255, 255, 255))
    }

    rl.PopMatrix()
    rl.EndMode3D()
}