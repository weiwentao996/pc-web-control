package main

import (
	"embed"
	"fmt"
	"github.com/go-vgo/robotgo"
	"io/fs"
	"log"
	"net/http"
	"strconv"
)

//go:embed static/*
var content embed.FS

func main() {
	http.HandleFunc("/click", mouseClick)
	http.HandleFunc("/key", keyPress)
	http.HandleFunc("/move", mouseMove)

	staticFiles, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(staticFiles))) // 提供静态文件服务

	fmt.Println("Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func mouseClick(w http.ResponseWriter, r *http.Request) {
	dir := r.URL.Query().Get("dir")
	if dir == "left" {
		robotgo.Click("left")
	} else if dir == "right" {
		robotgo.Click("right")
	}

	w.Write([]byte("Mouse Clicked"))
}

func keyPress(w http.ResponseWriter, r *http.Request) {
	dir := r.URL.Query().Get("dir")
	if dir == "left" {
		robotgo.KeyTap("left")
	} else if dir == "right" {
		robotgo.KeyTap("right")
	} else if dir == "up" {
		robotgo.KeyTap("up")
	} else if dir == "down" {
		robotgo.KeyTap("down")
	}
	w.Write([]byte("Key Pressed: " + dir))
}

func mouseMove(w http.ResponseWriter, r *http.Request) {
	// 获取光标当前坐标
	x, y := robotgo.Location()

	// 获取请求中的偏移量
	xOffsetStr := r.URL.Query().Get("x")
	yOffsetStr := r.URL.Query().Get("y")

	xOffset, _ := strconv.Atoi(xOffsetStr)
	yOffset, _ := strconv.Atoi(yOffsetStr)

	// 获取屏幕的宽度和高度
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// 计算新的光标坐标
	newX := x + xOffset
	newY := y + yOffset

	// 限制光标在屏幕范围内
	if newX < 0 {
		newX = 0
	} else if newX > screenWidth-1 {
		newX = screenWidth - 1
	}

	if newY < 0 {
		newY = 0
	} else if newY > screenHeight-1 {
		newY = screenHeight - 1
	}

	// 移动光标
	robotgo.Move(newX, newY)

	w.Write([]byte(fmt.Sprintf("Mouse moved by X: %d, Y: %d", xOffset, yOffset)))
}
