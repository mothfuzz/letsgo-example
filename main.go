package main

import (
	"embed"

	"github.com/mothfuzz/dyndraw/framework/actors"
	"github.com/mothfuzz/dyndraw/framework/app"
	"github.com/mothfuzz/dyndraw/framework/input"
	"github.com/mothfuzz/dyndraw/framework/render"
	"github.com/mothfuzz/dyndraw/framework/transform"
	. "github.com/mothfuzz/dyndraw/framework/vecmath"
)

//go:embed resources
var Resources embed.FS

type RayTest struct{}

func (r *RayTest) Init()    {}
func (r *RayTest) Update()  {}
func (r *RayTest) Destroy() {}
func (r *RayTest) Draw() {
	mx, my := input.GetMousePosition()
	t := transform.Origin2D(4, 4)
	t.SetPosition(640/2, 400/2, -1)
	render.DrawSprite("pointg.png", t.Mat4())
	ray := Vec3{float32(mx), float32(my), 0}.Sub(Vec3{640 / 2, 400.0 / 2, 0}).Normalize()
	for _, p := range (RayCast(Vec3{640 / 2, 400 / 2, 0}, CurrentLevel.Planes, ray)) {
		t.SetPosition(p.I.X(), p.I.Y(), -1)
		render.DrawSprite("point.png", t.Mat4())
	}
	if hit, ok := RayCastLen(Vec3{640 / 2, 400 / 2, 0}, CurrentLevel.Planes, ray, 640/2); ok {
		t := transform.Origin2D(4, 4)
		t.SetPosition(hit.I.X(), hit.I.Y(), -2)
		render.DrawSprite("pointg.png", t.Mat4())
	}
}

func main() {
	render.Resources = Resources
	app.Init()
	defer app.Quit()

	app.SetWindowSize(640, 400)

	LoadItemDictionary()

	t := &TileMap{
		TileSet: TileSet{"tileset.png", 4, 4, 16, 16},
		Data: [][]uint8{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 3, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 0, 0},
			{0, 0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0, 0},
			{0, 0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0, 0},
			{0, 1, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 3, 0, 2, 0},
			{1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2},
		},
	}
	CurrentLevel = t
	actors.Spawn(t)
	actors.Spawn(&Player{})
	actors.Spawn(&RayTest{})
	//i := &Item{Name: "Thingy", Description: "hello there", Icon: "bnw.png"}
	//actors.Spawn(i)
	//i.SetPosition2D(640/2+32, 400/2)

	actors.SpawnAt(ItemDictionary("thingy.xml"), transform.Location2D(640/2+16, 480/2, 16, 16))
	actors.SpawnAt(ItemDictionary("otherthingy.json"), transform.Location2D(640/2+32, 480/2, 16, 16))

	for app.PollEvents() {
		app.Update()
		app.Draw()
	}
}
