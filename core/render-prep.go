package core

func (_ *NgCore) onPrep() {
	for cid := 0; cid < len(Core.Render.Canvases); cid++ {
		if Core.Render.Canvases[cid].renderThisFrame() {
			Core.Render.Canvases[cid].onPrep()
		}
	}
}

func (me *RenderCanvas) onPrep() {
	for view := 0; view < len(me.Views); view++ {
		me.Views[view].onPrep()
	}
}

func (me *RenderView) onPrep() {
	if me.Enabled {
		me.Technique.onPrep()
	}
}

func (me *RenderTechniqueScene) onPrep() {
	if scene := me.Camera.Scene(); scene != nil {
		me.Camera.thrPrep.matCamProj.SetFromMult4(&me.Camera.thrPrep.matProj, &me.Camera.Controller.thrPrep.mat)
		// println(strf("F %v", me.Camera.thrPrep.frustum))
		if !scene.thrPrep.done {
			scene.thrPrep.done = true
			scene.onPrep()
		}
		bc := 0
		me.Camera.onPrep(scene.allNodes, 0, &bc)
		me.numDrawCalls = bc
		if me.Batch.Enabled {
			me.prepBatch(scene, me.numDrawCalls)
		}
	}
}

func (me *Camera) onPrep(all SceneNodeLib, nodeID int, batchCounter *int) {
	var mesh *Mesh
	var mat *FxMaterial
	camNodeRender := all[nodeID].Render.Enabled && (all[nodeID].parentID < 1 || me.thrPrep.nodeRender[all[nodeID].parentID])
	if camNodeRender {
		if mesh, mat = all[nodeID].meshMat(); mesh == nil || mat == nil {
			camNodeRender = false
		}
	}
	if me.thrPrep.nodeRender[nodeID] = camNodeRender; camNodeRender {
		if me.Perspective.Enabled {
			if all[nodeID].Render.skyMode {
				me.thrPrep.nodeProjMats[nodeID].SetFromMult4(&me.thrPrep.matCamProj, &me.thrPrep.matPos)
			} else {
				me.thrPrep.nodeProjMats[nodeID].SetFromMult4(&me.thrPrep.matCamProj, &all[nodeID].Transform.thrPrep.matModelView)
			}
			if me.Cull.Frustum && all[nodeID].Render.Cull.Frustum {
				if !me.frustumHasSphere(&all[nodeID].Transform.Pos, all[nodeID].thrPrep.bounding.sphere) {
					camNodeRender, me.thrPrep.nodeRender[nodeID] = false, false
				}
			}
		} else {
			me.thrPrep.nodeProjMats[nodeID] = all[nodeID].Transform.thrPrep.matModelView
		}
		if camNodeRender {
			if mat.HasFaceEffects() {
				*batchCounter = *batchCounter + len(mesh.raw.faces)
			} else {
				*batchCounter = *batchCounter + 1
			}
		}
	}
	for i := 0; i < len(all[nodeID].childNodeIDs); i++ {
		if all.IsOk(all[nodeID].childNodeIDs[i]) {
			me.onPrep(all, all[nodeID].childNodeIDs[i], batchCounter)
		}
	}
}

func (me *Scene) onPrep() {
	me.thrPrep.copyDone, me.thrRend.copyDone = false, false
}
