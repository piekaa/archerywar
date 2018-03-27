package main

import (
	"encoding/binary"
	"log"

	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
	"github.com/piekaa/archerywar/src/utils/shader"
	"github.com/piekaa/archerywar/src/utils/textures/png"
)

var (
	program gl.Program
	buf     gl.Buffer
)

func main() {

	app.Main(func(a app.App) {
		var glctx gl.Context
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(glctx)
					glctx = nil
				}
			case size.Event:
			case paint.Event:
				if glctx == nil || e.External {
					// As we are actively painting as fast as
					// we can (usually 60 FPS), skip any paint
					// events sent by the system.
					continue
				}
				onPaint(glctx)
				a.Publish()
				// Drive the animation by preparing to paint the next frame
				// after this one is shown.
				a.Send(paint.Event{})
			case touch.Event:
			}
		}
	})
}

func onStart(g gl.Context) {
	fmt.Println("On start")
	var err error


	g.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	g.Enable(gl.BLEND)

	vs, fs, err := shader.Read("simple")

	if err != nil {
		//todo change that
		panic(err)
	}

	fmt.Println(vs)

	program, err = glutil.CreateProgram(g, vs, fs)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}
	g.UseProgram(program)


	u_color := g.GetUniformLocation(program, "u_Color")
	fmt.Println(u_color)

	for g.GetError() !=0 {
	}
	g.Uniform4f(u_color, 0.0, 0.0, 0.1, 1.0  )
	glError := g.GetError()
	if glError != 0 {
		fmt.Println("Errors: ")
		for( glError != 0) {
			fmt.Println(glError)
			glError = g.GetError()
		}
		fmt.Println("End errors")
	}



	positionsAndTextCords := f32.Bytes(binary.LittleEndian,
		-0.5,-0.5,0,0,  0.5,-0.5,1,0, 0.5, 0.5,1,1, -0.5,0.5,0,1)



	indices := []byte{
		0, 1, 2,
		2, 3, 0}


	vao := g.CreateVertexArray()
	g.BindVertexArray(vao)

	buffer := g.CreateBuffer()
	g.BindBuffer(gl.ARRAY_BUFFER, buffer)
	g.BufferData(gl.ARRAY_BUFFER, positionsAndTextCords, gl.STATIC_DRAW)





	positionAttrib := g.GetAttribLocation(program, "position")
	g.EnableVertexAttribArray(positionAttrib)
	g.VertexAttribPointer(positionAttrib, 2, gl.FLOAT, false, 16, 0)


	texCoordAttrib := g.GetAttribLocation(program, "texCoord")
	fmt.Println("Tex")
	fmt.Println(texCoordAttrib)

	//g.DisableVertexAttribArray(positionAttrib)


	g.EnableVertexAttribArray(texCoordAttrib)
	g.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 16, 8)


	indexBuffer := g.CreateBuffer()
	g.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer)
	g.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

	g.ActiveTexture(gl.TEXTURE0)

	tex := g.CreateTexture()
	g.BindTexture(gl.TEXTURE_2D, tex)
	g.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	g.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	g.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	g.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	texBytes,w,h, err := png.ReadAsBytes("square")

	if err != nil {
		panic(err)
	}

	g.TexImage2D(gl.TEXTURE_2D, 0, w, h, gl.RGBA, gl.UNSIGNED_BYTE, texBytes)





	glError = g.GetError()
	if glError != 0 {
		fmt.Println("Errors: ")
		for( glError != 0) {
			fmt.Println(glError)
			glError = g.GetError()
		}
		fmt.Println("End errors")
	}

	slot := g.GetUniformLocation(program, "u_Texture")
	g.Uniform1i(slot, 0)


}

func onStop(glctx gl.Context) {
	glctx.DeleteProgram(program)
	glctx.DeleteBuffer(buf)
}

func onPaint(g gl.Context) {
	g.ClearColor(0, 0, 0.2, 1)
	g.Clear(gl.COLOR_BUFFER_BIT)

	g.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_BYTE , 0)
}
