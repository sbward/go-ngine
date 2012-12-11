package collimp

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"

	nga "github.com/go3d/go-ngine/assets"
)

func libs_animation_clips(xn *xmlx.Node) {
	var (
		lib *nga.LibAnimationClipDefs
		def *nga.AnimationClipDef
		id  string
	)
	for _, ln := range xcns(xn, "library_animation_clips") {
		id = xa(ln, "id")
		if lib = nga.AllAnimationClipDefLibs[id]; lib == nil {
			lib = nga.AllAnimationClipDefLibs.AddNew(id)
		}
		for _, def = range objs_AnimationClipDef(ln, "animation_clip") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_animations(xn *xmlx.Node) {
	var (
		lib *nga.LibAnimationDefs
		def *nga.AnimationDef
		id  string
	)
	for _, ln := range xcns(xn, "library_animations") {
		id = xa(ln, "id")
		if lib = nga.AllAnimationDefLibs[id]; lib == nil {
			lib = nga.AllAnimationDefLibs.AddNew(id)
		}
		for _, def = range objs_AnimationDef(ln, "animation") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_cameras(xn *xmlx.Node) {
	var (
		lib *nga.LibCameraDefs
		def *nga.CameraDef
		id  string
	)
	for _, ln := range xcns(xn, "library_cameras") {
		id = xa(ln, "id")
		if lib = nga.AllCameraDefLibs[id]; lib == nil {
			lib = nga.AllCameraDefLibs.AddNew(id)
		}
		for _, def = range objs_CameraDef(ln, "camera") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_controllers(xn *xmlx.Node) {
	var (
		lib *nga.LibControllerDefs
		def *nga.ControllerDef
		id  string
	)
	for _, ln := range xcns(xn, "library_controllers") {
		id = xa(ln, "id")
		if lib = nga.AllControllerDefLibs[id]; lib == nil {
			lib = nga.AllControllerDefLibs.AddNew(id)
		}
		for _, def = range objs_ControllerDef(ln, "controller") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_formulas(xn *xmlx.Node) {
	var (
		lib *nga.LibFormulaDefs
		def *nga.FormulaDef
		id  string
	)
	for _, ln := range xcns(xn, "library_formulas") {
		id = xa(ln, "id")
		if lib = nga.AllFormulaDefLibs[id]; lib == nil {
			lib = nga.AllFormulaDefLibs.AddNew(id)
		}
		for _, def = range objs_FormulaDef(ln, "formula") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_geometries(xn *xmlx.Node) {
	var (
		lib *nga.LibGeometryDefs
		def *nga.GeometryDef
		id  string
	)
	for _, ln := range xcns(xn, "library_geometries") {
		id = xa(ln, "id")
		if lib = nga.AllGeometryDefLibs[id]; lib == nil {
			lib = nga.AllGeometryDefLibs.AddNew(id)
		}
		for _, def = range objs_GeometryDef(ln, "geometry") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_lights(xn *xmlx.Node) {
	var (
		lib *nga.LibLightDefs
		def *nga.LightDef
		id  string
	)
	for _, ln := range xcns(xn, "library_lights") {
		id = xa(ln, "id")
		if lib = nga.AllLightDefLibs[id]; lib == nil {
			lib = nga.AllLightDefLibs.AddNew(id)
		}
		for _, def = range objs_LightDef(ln, "light") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_nodes(xn *xmlx.Node) {
	var (
		lib *nga.LibNodeDefs
		def *nga.NodeDef
		id  string
	)
	for _, ln := range xcns(xn, "library_nodes") {
		id = xa(ln, "id")
		if lib = nga.AllNodeDefLibs[id]; lib == nil {
			lib = nga.AllNodeDefLibs.AddNew(id)
		}
		for _, def = range objs_NodeDef(ln, "node") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_visual_scenes(xn *xmlx.Node) {
	var (
		lib *nga.LibVisualSceneDefs
		def *nga.VisualSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_visual_scenes") {
		id = xa(ln, "id")
		if lib = nga.AllVisualSceneDefLibs[id]; lib == nil {
			lib = nga.AllVisualSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_VisualSceneDef(ln, "visual_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_force_fields(xn *xmlx.Node) {
	var (
		lib *nga.LibPxForceFieldDefs
		def *nga.PxForceFieldDef
		id  string
	)
	for _, ln := range xcns(xn, "library_force_fields") {
		id = xa(ln, "id")
		if lib = nga.AllPxForceFieldDefLibs[id]; lib == nil {
			lib = nga.AllPxForceFieldDefLibs.AddNew(id)
		}
		for _, def = range objs_PxForceFieldDef(ln, "force_field") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_materials(xn *xmlx.Node) {
	var (
		lib *nga.LibPxMaterialDefs
		def *nga.PxMaterialDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_materials") {
		id = xa(ln, "id")
		if lib = nga.AllPxMaterialDefLibs[id]; lib == nil {
			lib = nga.AllPxMaterialDefLibs.AddNew(id)
		}
		for _, def = range objs_PxMaterialDef(ln, "physics_material") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_models(xn *xmlx.Node) {
	var (
		lib *nga.LibPxModelDefs
		def *nga.PxModelDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_models") {
		id = xa(ln, "id")
		if lib = nga.AllPxModelDefLibs[id]; lib == nil {
			lib = nga.AllPxModelDefLibs.AddNew(id)
		}
		for _, def = range objs_PxModelDef(ln, "physics_model") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_physics_scenes(xn *xmlx.Node) {
	var (
		lib *nga.LibPxSceneDefs
		def *nga.PxSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_physics_scenes") {
		id = xa(ln, "id")
		if lib = nga.AllPxSceneDefLibs[id]; lib == nil {
			lib = nga.AllPxSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_PxSceneDef(ln, "physics_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_effects(xn *xmlx.Node) {
	var (
		lib *nga.LibFxEffectDefs
		def *nga.FxEffectDef
		id  string
	)
	for _, ln := range xcns(xn, "library_effects") {
		id = xa(ln, "id")
		if lib = nga.AllFxEffectDefLibs[id]; lib == nil {
			lib = nga.AllFxEffectDefLibs.AddNew(id)
		}
		for _, def = range objs_FxEffectDef(ln, "effect") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_images(xn *xmlx.Node) {
	var (
		lib *nga.LibFxImageDefs
		def *nga.FxImageDef
		id  string
	)
	for _, ln := range xcns(xn, "library_images") {
		id = xa(ln, "id")
		if lib = nga.AllFxImageDefLibs[id]; lib == nil {
			lib = nga.AllFxImageDefLibs.AddNew(id)
		}
		for _, def = range objs_FxImageDef(ln, "image") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_materials(xn *xmlx.Node) {
	var (
		lib *nga.LibFxMaterialDefs
		def *nga.FxMaterialDef
		id  string
	)
	for _, ln := range xcns(xn, "library_materials") {
		id = xa(ln, "id")
		if lib = nga.AllFxMaterialDefLibs[id]; lib == nil {
			lib = nga.AllFxMaterialDefLibs.AddNew(id)
		}
		for _, def = range objs_FxMaterialDef(ln, "material") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_articulated_systems(xn *xmlx.Node) {
	var (
		lib *nga.LibKxArticulatedSystemDefs
		def *nga.KxArticulatedSystemDef
		id  string
	)
	for _, ln := range xcns(xn, "library_articulated_systems") {
		id = xa(ln, "id")
		if lib = nga.AllKxArticulatedSystemDefLibs[id]; lib == nil {
			lib = nga.AllKxArticulatedSystemDefLibs.AddNew(id)
		}
		for _, def = range objs_KxArticulatedSystemDef(ln, "articulated_system") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_joints(xn *xmlx.Node) {
	var (
		lib *nga.LibKxJointDefs
		def *nga.KxJointDef
		id  string
	)
	for _, ln := range xcns(xn, "library_joints") {
		id = xa(ln, "id")
		if lib = nga.AllKxJointDefLibs[id]; lib == nil {
			lib = nga.AllKxJointDefLibs.AddNew(id)
		}
		for _, def = range objs_KxJointDef(ln, "joint") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_kinematics_models(xn *xmlx.Node) {
	var (
		lib *nga.LibKxModelDefs
		def *nga.KxModelDef
		id  string
	)
	for _, ln := range xcns(xn, "library_kinematics_models") {
		id = xa(ln, "id")
		if lib = nga.AllKxModelDefLibs[id]; lib == nil {
			lib = nga.AllKxModelDefLibs.AddNew(id)
		}
		for _, def = range objs_KxModelDef(ln, "kinematics_model") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_kinematics_scenes(xn *xmlx.Node) {
	var (
		lib *nga.LibKxSceneDefs
		def *nga.KxSceneDef
		id  string
	)
	for _, ln := range xcns(xn, "library_kinematics_scenes") {
		id = xa(ln, "id")
		if lib = nga.AllKxSceneDefLibs[id]; lib == nil {
			lib = nga.AllKxSceneDefLibs.AddNew(id)
		}
		for _, def = range objs_KxSceneDef(ln, "kinematics_scene") {
			if def != nil {
				lib.Add(def)
			}
		}
		lib.SetDirty()
	}
}

func libs_All(xn *xmlx.Node) {
	libs_animation_clips(xn)
	libs_animations(xn)
	libs_cameras(xn)
	libs_controllers(xn)
	libs_formulas(xn)
	libs_geometries(xn)
	libs_lights(xn)
	libs_nodes(xn)
	libs_visual_scenes(xn)
	libs_force_fields(xn)
	libs_physics_materials(xn)
	libs_physics_models(xn)
	libs_physics_scenes(xn)
	libs_effects(xn)
	libs_images(xn)
	libs_materials(xn)
	libs_articulated_systems(xn)
	libs_joints(xn)
	libs_kinematics_models(xn)
	libs_kinematics_scenes(xn)
}
