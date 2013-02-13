package core

type RenderTechnique interface {
	dispose()
	name() string
	onRenderNode()
	render()
}

type renderTechniqueBase struct {
	DefaultEffect FxEffect
	tname         string
}

func (me *renderTechniqueBase) dispose() {
}

func (me *renderTechniqueBase) init(name string) {
	me.tname = name
	me.DefaultEffect.init()
	me.DefaultEffect.Ops.EnableTex2D(-1)
	me.DefaultEffect.UpdateRoutine()
}

func (me *renderTechniqueBase) name() string {
	return me.tname
}

func (me *renderTechniqueBase) onRenderNode() {
}

//	Used only in Core.Rendering.Techniques.
type RenderTechniques map[string]RenderTechnique

func (me *RenderTechniques) dispose() {
	for _, o := range *me {
		o.dispose()
	}
	*me = RenderTechniques{}
}
