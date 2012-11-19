package core

import (
	gl "github.com/chsc/gogl/gl42"
	ugl "github.com/go3d/go-glutil"
)

var (
	techs map[string]renderTechnique
	tmpMat *Material
)

type techniqueCtor func (string) renderTechnique

func initTechniques () {
	techs = map[string]renderTechnique {}
	for techName, techMaker := range map[string]techniqueCtor { "rt_unlit_colored": newTechnique_UnlitColored, "rt_unlit_textured": newTechnique_UnlitTextured } {
		techs[techName] = techMaker(techName)
	}
}

type renderTechnique interface {
	initMeshBuffer (meshBuffer *MeshBuffer)
	name () string
	onPreRender ()
	onRenderMesh ()
	onRenderMeshModel ()
	onRenderNode ()
}

type baseTechnique struct {
	prog *ugl.ShaderProgram
}

	func (me *baseTechnique) initMeshBuffer (meshBuffer *MeshBuffer) {
		gl.EnableVertexAttribArray(me.prog.AttrLocs["aPos"])
		gl.VertexAttribPointer(me.prog.AttrLocs["aPos"], 3, gl.FLOAT, gl.FALSE, 8 * 4, gl.Pointer(nil))
	}

	func (me *baseTechnique) name () string {
		return me.prog.Name
	}

	func (me *baseTechnique) onPreRender () {
	}

	func (me *baseTechnique) onRenderMesh () {
	}

	func (me *baseTechnique) onRenderMeshModel () {
	}

	func (me *baseTechnique) onRenderNode () {
	}

	func (me *baseTechnique) setProg (name string, unifs []string, attrs []string) {
		var prog = glShaderMan.AllProgs[name]
		prog.SetUnifLocations("uMatCam", "uMatModelView", "uMatProj")
		if len(unifs) > 0 { prog.SetUnifLocations(unifs...) }
		prog.SetAttrLocations("aPos")
		if len(attrs) > 0 { prog.SetAttrLocations(attrs...) }
		me.prog = prog
	}

type techniqueUnlitColored struct {
	baseTechnique
}

	func newTechnique_UnlitColored (progName string) renderTechnique {
		var tech = &techniqueUnlitColored {}
		tech.baseTechnique.setProg(progName, nil, nil)
		return tech
	}

type techniqueUnlitTextured struct {
	baseTechnique
}

	func newTechnique_UnlitTextured (progName string) renderTechnique {
		var tech = &techniqueUnlitTextured {}
		tech.baseTechnique.setProg(progName, []string { "uTex0" }, []string { "aTexCoords" })
		return tech
	}

	func (me *techniqueUnlitTextured) initMeshBuffer (meshBuffer *MeshBuffer) {
		me.baseTechnique.initMeshBuffer(meshBuffer)
		gl.EnableVertexAttribArray(me.prog.AttrLocs["aTexCoords"])
		gl.VertexAttribPointer(me.prog.AttrLocs["aTexCoords"], 2, gl.FLOAT, gl.FALSE, 8 * 4, gl.Offset(nil, 3 * 4))
	}

	func (me *techniqueUnlitTextured) onRenderNode () {
		if tmpMat = curNode.Material(); tmpMat != curMat {
			if curMat = tmpMat; curMat != nil {
				gl.ActiveTexture(gl.TEXTURE0)
				gl.BindTexture(gl.TEXTURE_2D, Core.Textures[curMat.TexName].glTex)
				gl.Uniform1i(curProg.UnifLocs["uTex0"], 0)
			}
		}
	}
