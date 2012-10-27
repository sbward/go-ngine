package main

import (
	"math"

	ngine "github.com/go3d/go-ngine/core"
	ngine_samples "github.com/go3d/go-ngine/_sampleprogs/_sharedcode"
)

var (
	tri, quad *ngine.TNode
)

func main () {
	ngine_samples.MaxKeyHint = 3
	ngine_samples.SamplesMainFunc(LoadSampleScene_01_TriQuad)
}

func onLoop () {
	ngine_samples.CheckToggleKeys()
	tri.Transform.Rot.X -= 0.005
	tri.Transform.Rot.Y -= 0.005
	tri.Transform.Pos.Set(-3.75, 1 * math.Sin(ngine.Loop.TickNow), 1)
	tri.Transform.OnPosRotChanged()
	quad.Transform.Rot.Y += 0.0001
	quad.Transform.Rot.Z += 0.0001
	quad.Transform.Pos.Set(-4.125, 1 * math.Cos(ngine.Loop.TickNow), 0)
	quad.Transform.OnPosRotChanged()
}

func LoadSampleScene_01_TriQuad () {
	ngine.Core.Options.SetGlBackfaceCulling(false)

	ngine.Core.Textures["texcat"] = ngine.Core.Textures.Load(ngine.TextureProviders.LocalFile, "misc/cat.png")
	ngine.Core.Textures["texdog"] = ngine.Core.Textures.Load(ngine.TextureProviders.LocalFile, "misc/dog.png")

	ngine.Core.Materials["matcat"] = ngine.Core.Materials.New("texcat")
	ngine.Core.Materials["matdog"] = ngine.Core.Materials.New("texdog")

	ngine.Core.Meshes["face3"], _ = ngine.Core.Meshes.Load(ngine.MeshProviders.PrefabTri)
	ngine.Core.Meshes["face4"], _ = ngine.Core.Meshes.Load(ngine.MeshProviders.PrefabQuad)

	ngine.Core.MeshBuffers.Add("buf_tri", 3, 3, false)
	ngine.Core.MeshBuffers.Add("buf_quad", 6, 6, false)

	var scene = ngine.NewScene()
	ngine.Core.Scenes[""] = scene

	scene.RootNode.AddSubNodesNamed(map[string]string { "tri":  "face3", "quad": "face4" })
	tri, quad = scene.RootNode.SubNodes["tri"], scene.RootNode.SubNodes["quad"]
	tri.SetMatKey("matcat")
	quad.SetMatKey("matdog")
	ngine.Core.SyncUpdates()
	ngine.Loop.AddHandler(onLoop)
}
