/*
 * ╔════════════════════════════════════════════════════════════════╗
 * ║ Intro_raylib_1                                                 ║
 * ║ Plik / File: drawFigures.go                                    ║
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


package main

import (
    "math"

    rl "github.com/gen2brain/raylib-go/raylib"
)

func drawCone(position rl.Vector3, radius float32, height float32, segments int32, color rl.Color) {
	// Rysowanie podstawy
	for i := 0; i < int(segments); i++ {
		angle1 := float32(i) * (2 * math.Pi / float32(segments))
		angle2 := float32(i+1) * (2 * math.Pi / float32(segments))

		p1 := rl.NewVector3(position.X+radius*float32(math.Cos(float64(angle1))), position.Y, position.Z+radius*float32(math.Sin(float64(angle1))))
		p2 := rl.NewVector3(position.X+radius*float32(math.Cos(float64(angle2))), position.Y, position.Z+radius*float32(math.Sin(float64(angle2))))

		rl.DrawLine3D(p1, p2, color) // Podstawa
	}

	// Rysowanie boków
	for i := 0; i < int(segments); i++ {
		angle := float32(i) * (2 * math.Pi / float32(segments))

		p1 := rl.NewVector3(position.X+radius*float32(math.Cos(float64(angle))), position.Y, position.Z+radius*float32(math.Sin(float64(angle))))
		p2 := rl.NewVector3(position.X, position.Y+height, position.Z)

		rl.DrawLine3D(p1, p2, color) // Bok
	}
}

func drawPyramid(position rl.Vector3, size float32, color rl.Color) {
	halfSize := size / 2

	// Wierzchołki piramidy
	v1 := rl.NewVector3(position.X, position.Y+size, position.Z)
	v2 := rl.NewVector3(position.X-halfSize, position.Y, position.Z-halfSize)
	v3 := rl.NewVector3(position.X+halfSize, position.Y, position.Z-halfSize)
	v4 := rl.NewVector3(position.X+halfSize, position.Y, position.Z+halfSize)
	v5 := rl.NewVector3(position.X-halfSize, position.Y, position.Z+halfSize)

	// Rysowanie podstawy
	rl.DrawLine3D(v2, v3, color)
	rl.DrawLine3D(v3, v4, color)
	rl.DrawLine3D(v4, v5, color)
	rl.DrawLine3D(v5, v2, color)

	// Rysowanie boków
	rl.DrawLine3D(v1, v2, color)
	rl.DrawLine3D(v1, v3, color)
	rl.DrawLine3D(v1, v4, color)
	rl.DrawLine3D(v1, v5, color)
}

func drawDiamond(position rl.Vector3, size float32, color rl.Color) {
	halfSize := size / 2

	// Wierzchołki diamentu
	top := rl.NewVector3(position.X, position.Y+halfSize, position.Z)
	bottom := rl.NewVector3(position.X, position.Y-halfSize, position.Z)
	front := rl.NewVector3(position.X, position.Y, position.Z+halfSize)
	back := rl.NewVector3(position.X, position.Y, position.Z-halfSize)
	left := rl.NewVector3(position.X-halfSize, position.Y, position.Z)
	right := rl.NewVector3(position.X+halfSize, position.Y, position.Z)

	// Rysowanie górnej części diamentu (górna piramida)
	rl.DrawLine3D(top, front, color)
	rl.DrawLine3D(top, back, color)
	rl.DrawLine3D(top, left, color)
	rl.DrawLine3D(top, right, color)

	// Rysowanie dolnej części diamentu (dolna piramida)
	rl.DrawLine3D(bottom, front, color)
	rl.DrawLine3D(bottom, back, color)
	rl.DrawLine3D(bottom, left, color)
	rl.DrawLine3D(bottom, right, color)

	// Rysowanie połączeń między górną a dolną piramidą
	rl.DrawLine3D(front, back, color)
	rl.DrawLine3D(back, left, color)
	rl.DrawLine3D(left, right, color)
	rl.DrawLine3D(right, front, color)
}

func drawCrystal(position rl.Vector3, size float32, color rl.Color) {
	halfSize := size / 2
	quarterSize := size / 4

	// Wierzchołki brylantu
	top := rl.NewVector3(position.X, position.Y+halfSize, position.Z)
	bottom := rl.NewVector3(position.X, position.Y-halfSize, position.Z)
	middleTopFront := rl.NewVector3(position.X, position.Y+quarterSize, position.Z+quarterSize)
	middleTopBack := rl.NewVector3(position.X, position.Y+quarterSize, position.Z-quarterSize)
	middleTopLeft := rl.NewVector3(position.X-quarterSize, position.Y+quarterSize, position.Z)
	middleTopRight := rl.NewVector3(position.X+quarterSize, position.Y+quarterSize, position.Z)
	middleBottomFront := rl.NewVector3(position.X, position.Y-quarterSize, position.Z+quarterSize)
	middleBottomBack := rl.NewVector3(position.X, position.Y-quarterSize, position.Z-quarterSize)
	middleBottomLeft := rl.NewVector3(position.X-quarterSize, position.Y-quarterSize, position.Z)
	middleBottomRight := rl.NewVector3(position.X+quarterSize, position.Y-quarterSize, position.Z)

	// Rysowanie górnej części brylantu
	rl.DrawLine3D(top, middleTopFront, color)
	rl.DrawLine3D(top, middleTopBack, color)
	rl.DrawLine3D(top, middleTopLeft, color)
	rl.DrawLine3D(top, middleTopRight, color)

	// Rysowanie dolnej części brylantu
	rl.DrawLine3D(bottom, middleBottomFront, color)
	rl.DrawLine3D(bottom, middleBottomBack, color)
	rl.DrawLine3D(bottom, middleBottomLeft, color)
	rl.DrawLine3D(bottom, middleBottomRight, color)

	// Rysowanie połączeń między środkowymi wierzchołkami
	rl.DrawLine3D(middleTopFront, middleTopRight, color)
	rl.DrawLine3D(middleTopRight, middleTopBack, color)
	rl.DrawLine3D(middleTopBack, middleTopLeft, color)
	rl.DrawLine3D(middleTopLeft, middleTopFront, color)

	rl.DrawLine3D(middleBottomFront, middleBottomRight, color)
	rl.DrawLine3D(middleBottomRight, middleBottomBack, color)
	rl.DrawLine3D(middleBottomBack, middleBottomLeft, color)
	rl.DrawLine3D(middleBottomLeft, middleBottomFront, color)

	// Rysowanie połączeń między górnymi a dolnymi wierzchołkami
	rl.DrawLine3D(middleTopFront, middleBottomFront, color)
	rl.DrawLine3D(middleTopBack, middleBottomBack, color)
	rl.DrawLine3D(middleTopLeft, middleBottomLeft, color)
	rl.DrawLine3D(middleTopRight, middleBottomRight, color)
}

func drawBrilliantBasedOnCone(position rl.Vector3, size float32, height float32, sides int, color rl.Color) {
	halfHeight := height / 2
	radius := size / 2

	// Wierzchołki brylantu
	top := rl.NewVector3(position.X, position.Y+halfHeight, position.Z)
	bottom := rl.NewVector3(position.X, position.Y-halfHeight, position.Z)

	// Górne i dolne punkty podstawy
	points := make([]rl.Vector3, sides)
	for i := 0; i < sides; i++ {
		angle := float32(i) * (2 * math.Pi / float32(sides))
		x := radius * float32(math.Cos(float64(angle)))
		z := radius * float32(math.Sin(float64(angle)))
		points[i] = rl.NewVector3(position.X+x, position.Y, position.Z+z)
	}

	// Rysowanie górnej części brylantu (korona)
	for i := 0; i < sides; i++ {
		nextIndex := (i + 1) % sides
		rl.DrawLine3D(top, points[i], color)
		rl.DrawLine3D(points[i], points[nextIndex], color)
	}

	// Rysowanie dolnej części brylantu (pavilion)
	for i := 0; i < sides; i++ {
		nextIndex := (i + 1) % sides
		rl.DrawLine3D(bottom, points[i], color)
		rl.DrawLine3D(bottom, points[nextIndex], color)
	}

	// Rysowanie podstawy brylantu
	for i := 0; i < sides; i++ {
		nextIndex := (i + 1) % sides
		midPoint := rl.NewVector3((points[i].X+points[nextIndex].X)/2, (points[i].Y+points[nextIndex].Y)/2, (points[i].Z+points[nextIndex].Z)/2)
		rl.DrawLine3D(midPoint, bottom, color)
	}
}

func drawBrilliant(position rl.Vector3, radius float32, height float32, segments int, color rl.Color) {
	// Ustalanie parametrów
	topRadius := radius * 0.5           // Promień górnej spłaszczonej części
	bottomRadius := radius              // Promień dolnej części
	topHeight := height * 0.3           // Wysokość górnej spłaszczonej części
	bottomHeight := height * 0.7        // Wysokość dolnej części
	centerY := position.Y               // Współrzędna Y dla środka brylantu
	topY := centerY + topHeight/2       // Współrzędna Y dla górnej spłaszczonej części
	bottomY := centerY - bottomHeight/2 // Współrzędna Y dla dolnej części

	// Wierzchołki podstaw
	topVertices := make([]rl.Vector3, segments)
	bottomVertices := make([]rl.Vector3, segments)
	for i := 0; i < segments; i++ {
		angle := float32(i) * (2 * math.Pi / float32(segments))
		x := float32(math.Cos(float64(angle)))
		z := float32(math.Sin(float64(angle)))
		topVertices[i] = rl.NewVector3(position.X+x*topRadius, topY, position.Z+z*topRadius)
		bottomVertices[i] = rl.NewVector3(position.X+x*bottomRadius, bottomY, position.Z+z*bottomRadius)
	}

	// Rysowanie boków brylantu
	for i := 0; i < segments; i++ {
		nextIndex := (i + 1) % segments
		// Bok między górną i dolną częścią
		rl.DrawLine3D(topVertices[i], bottomVertices[i], color)
		rl.DrawLine3D(bottomVertices[i], bottomVertices[nextIndex], color)
		rl.DrawLine3D(topVertices[i], topVertices[nextIndex], color)
	}

	// Rysowanie górnej podstawy (table)
	for i := 0; i < segments; i++ {
		nextIndex := (i + 1) % segments
		rl.DrawLine3D(topVertices[i], topVertices[nextIndex], color)
	}

	// Rysowanie dolnej podstawy (rozszerzającej się)
	for i := 0; i < segments; i++ {
		nextIndex := (i + 1) % segments
		rl.DrawLine3D(bottomVertices[i], bottomVertices[nextIndex], color)
	}
}

func drawBrilliantWithCone(position rl.Vector3, radius float32, height float32, segments int, color rl.Color) {
	// Ustalanie parametrów
	topRadius := radius * 0.4         // Promień górnej spłaszczonej części
	bottomRadius := radius            // Promień dolnej części
	topHeight := height * 0.3         // Wysokość górnej spłaszczonej części
	bottomHeight := height * 0.7      // Wysokość dolnej części
	centerY := position.Y             // Współrzędna Y dla środka brylantu
	topY := centerY + topHeight/2     // Współrzędna Y dla górnej spłaszczonej części
	bottomY := centerY - bottomHeight // Współrzędna Y dla dolnej części

	// Wierzchołki podstaw
	topVertices := make([]rl.Vector3, segments)
	for i := 0; i < segments; i++ {
		angle := float32(i) * (2 * math.Pi / float32(segments))
		x := float32(math.Cos(float64(angle)))
		z := float32(math.Sin(float64(angle)))
		topVertices[i] = rl.NewVector3(position.X+x*topRadius, topY, position.Z+z*topRadius)
	}

	// Rysowanie boków górnej części
	for i := 0; i < segments; i++ {
		nextIndex := (i + 1) % segments
		rl.DrawLine3D(topVertices[i], topVertices[nextIndex], color)                         // Połączenie wierzchołków górnej podstawy
		rl.DrawLine3D(topVertices[i], rl.NewVector3(position.X, centerY, position.Z), color) // Połączenie wierzchołków z wierzchołkiem brylantu
	}

	// Rysowanie dolnego stożka
	for i := 0; i < segments; i++ {
		angle := float32(i) * (2 * math.Pi / float32(segments))
		x := float32(math.Cos(float64(angle)))
		z := float32(math.Sin(float64(angle)))
		bottomVertex := rl.NewVector3(position.X+x*bottomRadius, bottomY, position.Z+z*bottomRadius)

		// Rysowanie połączenia z górnymi wierzchołkami
		rl.DrawLine3D(topVertices[i], bottomVertex, color)

		// Rysowanie boków stożka
		rl.DrawLine3D(bottomVertex, rl.NewVector3(position.X, centerY, position.Z), color)
	}
}