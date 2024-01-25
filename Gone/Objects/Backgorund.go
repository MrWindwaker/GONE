package objects

import rl "github.com/gen2brain/raylib-go/raylib"

type BackgorundStatic struct {
}

type BackgroundAnimted struct {
	Texture     rl.Texture2D
	T_frames    int
	C_frame     int
	Is_Animated bool
}

func (b *BackgroundAnimted) Animate() {
	if b.Is_Animated {
		b.C_frame++

		if b.C_frame >= b.T_frames {
			b.C_frame = 0
		}
	}
}
func (b *BackgroundAnimted) Render() {
	rl.DrawTexturePro(
		b.Texture,
		b.get_source(),
		b.get_dest(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (b *BackgroundAnimted) get_source() rl.Rectangle {
	var x float32
	var w float32
	var h float32

	if b.T_frames == 0 && b.C_frame == 0 {
		x = 0
		w = float32(b.Texture.Width)
		h = float32(b.Texture.Height)
	} else {
		x = float32(b.Texture.Width/(int32(b.T_frames))) * float32(b.C_frame)
		w = float32(b.Texture.Width / (int32(b.T_frames)))
		h = float32(b.Texture.Height)
	}

	return rl.NewRectangle(
		x,
		0,
		w,
		h,
	)
}

func (b *BackgroundAnimted) get_dest() rl.Rectangle {

	var w float32
	var h float32

	if b.T_frames == 0 && b.C_frame == 0 {
		w = float32(rl.GetRenderWidth())
		h = float32(rl.GetRenderHeight())
	} else {
		w = float32(b.Texture.Width/(int32(b.T_frames))) * 4
		h = float32(b.Texture.Height) * 4
	}

	return rl.NewRectangle(
		0, 0,
		w,
		h,
	)
}
