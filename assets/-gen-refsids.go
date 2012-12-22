package assets

func (me *LightAttenuation) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Constant, nil, me.Constant.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.Linear, nil, me.Linear.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.Quadratic, nil, me.Quadratic.Sid); val != nil { return }
	return
}

func (me *GeometryDef) resolveSidPath(path []string) (val interface{}) {
	if me.Brep != nil {
	if val = me.Brep.resolveSidPath(path); val != nil { return }
	}
	if me.Mesh != nil {
	if val = me.Mesh.resolveSidPath(path); val != nil { return }
	}
	if me.Spline != nil {
	if val = me.Spline.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *ControllerDef) resolveSidPath(path []string) (val interface{}) {
	if me.Morph != nil {
	if val = me.Morph.resolveSidPath(path); val != nil { return }
	}
	if me.Skin != nil {
	if val = me.Skin.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *PxMaterialDef) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.TC.DynamicFriction, nil, me.TC.DynamicFriction.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.TC.Restitution, nil, me.TC.Restitution.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.TC.StaticFriction, nil, me.TC.StaticFriction.Sid); val != nil { return }
	return
}

func (me *KxModelDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.TC.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.TC.Links { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.TC.Formulas { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *AnimationClipDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Animations { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Formulas { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *Document) resolveSidPath(path []string) (val interface{}) {
	if me.Scene != nil {
	if val = me.Scene.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *GeometryBrepNurbs) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *FxProfileCommon) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Technique, &me.Technique, me.Technique.Sid); val != nil { return }
	return
}

func (me *AnimationDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	for _, v := range me.AnimationDefs { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *CameraOptics) resolveSidPath(path []string) (val interface{}) {
	if me.TC.AspectRatio != nil {
	if val = sidResolveCore(path, me.TC.AspectRatio, nil, me.TC.AspectRatio.Sid); val != nil { return }
	}
	if val = sidResolveCore(path, &me.TC.Zfar, nil, me.TC.Zfar.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.TC.Znear, nil, me.TC.Znear.Sid); val != nil { return }
	if me.TC.Orthographic != nil {
	if val = me.TC.Orthographic.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Perspective != nil {
	if val = me.TC.Perspective.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *ChildNode) resolveSidPath(path []string) (val interface{}) {
	if me.Inst != nil {
	if val = sidResolveCore(path, me.Inst, nil, me.Inst.Sid); val != nil { return }
	}
	return
}

func (me *KxFrameTcp) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxJoint) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Axis, nil, me.Axis.Sid); val != nil { return }
	if me.Limits != nil {
	if val = me.Limits.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *KxLink) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Attachments { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *PxRigidConstraintAttachment) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *GeometryMesh) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *FxEffectDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Profiles { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *GeometryBrepSurfaces) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.All { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *FxColorOrTexture) resolveSidPath(path []string) (val interface{}) {
	if me.Color != nil {
	if val = sidResolveCore(path, me.Color, nil, me.Color.Sid); val != nil { return }
	}
	return
}

func (me *FxSampler) resolveSidPath(path []string) (val interface{}) {
	if me.Image != nil {
	if val = sidResolveCore(path, me.Image, nil, me.Image.Sid); val != nil { return }
	}
	return
}

func (me *PxRigidConstraintSpring) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Stiffness, nil, me.Stiffness.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.Damping, nil, me.Damping.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.TargetValue, nil, me.TargetValue.Sid); val != nil { return }
	return
}

func (me *PxRigidBodyCommon) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Dynamic, nil, me.Dynamic.Sid); val != nil { return }
	if me.Mass != nil {
	if val = sidResolveCore(path, me.Mass, nil, me.Mass.Sid); val != nil { return }
	}
	for _, v := range me.MassFrame { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	if me.Inertia != nil {
	if val = sidResolveCore(path, me.Inertia, nil, me.Inertia.Sid); val != nil { return }
	}
	if val = me.Material.resolveSidPath(path); val != nil { return }
	for _, v := range me.Shapes { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *GeometryBrepCurves) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.All { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *KxKinematicsSystem) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Models { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.TC.AxisInfos { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	if val = me.TC.Frame.Origin.resolveSidPath(path); val != nil { return }
	if val = me.TC.Frame.Tip.resolveSidPath(path); val != nil { return }
	if me.TC.Frame.Tcp != nil {
	if val = me.TC.Frame.Tcp.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Frame.Object != nil {
	if val = me.TC.Frame.Object.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *FxPass) resolveSidPath(path []string) (val interface{}) {
	if me.Evaluate != nil {
	if val = me.Evaluate.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *GeometryBrepCurve) resolveSidPath(path []string) (val interface{}) {
	if me.Element.Nurbs != nil {
	if val = me.Element.Nurbs.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *Scene) resolveSidPath(path []string) (val interface{}) {
	if me.Visual != nil {
	if val = sidResolveCore(path, me.Visual, nil, me.Visual.Sid); val != nil { return }
	}
	if me.Kinematics != nil {
	if val = sidResolveCore(path, me.Kinematics, me.Kinematics, me.Kinematics.Sid); val != nil { return }
	}
	for _, v := range me.Physics { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *GeometryBrepNurbsSurface) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *PxModelInst) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.ForceFields { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.RigidBodies { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.RigidConstraints { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxSceneDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Models { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.ArticulatedSystems { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *KxJointDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.All { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *KxFrameOrigin) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *GeometryBrepSurface) resolveSidPath(path []string) (val interface{}) {
	if me.Element.NurbsSurface != nil {
	if val = me.Element.NurbsSurface.resolveSidPath(path); val != nil { return }
	}
	if me.Element.SweptSurface != nil {
	if val = me.Element.SweptSurface.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *FxTechniqueCommonPhong) resolveSidPath(path []string) (val interface{}) {
	if me.Emission != nil {
	if val = me.Emission.resolveSidPath(path); val != nil { return }
	}
	if me.Reflective != nil {
	if val = me.Reflective.resolveSidPath(path); val != nil { return }
	}
	if me.Reflectivity != nil {
	if val = me.Reflectivity.resolveSidPath(path); val != nil { return }
	}
	if me.Transparent != nil {
	if val = me.Transparent.resolveSidPath(path); val != nil { return }
	}
	if me.Transparency != nil {
	if val = me.Transparency.resolveSidPath(path); val != nil { return }
	}
	if me.IndexOfRefraction != nil {
	if val = me.IndexOfRefraction.resolveSidPath(path); val != nil { return }
	}
	if me.Ambient != nil {
	if val = me.Ambient.resolveSidPath(path); val != nil { return }
	}
	if me.Diffuse != nil {
	if val = me.Diffuse.resolveSidPath(path); val != nil { return }
	}
	if me.Specular != nil {
	if val = me.Specular.resolveSidPath(path); val != nil { return }
	}
	if me.Shininess != nil {
	if val = me.Shininess.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *FxPassEvaluationTarget) resolveSidPath(path []string) (val interface{}) {
	if me.Image != nil {
	if val = sidResolveCore(path, me.Image, nil, me.Image.Sid); val != nil { return }
	}
	return
}

func (me *KxJointLimits) resolveSidPath(path []string) (val interface{}) {
	if me.Min != nil {
	if val = sidResolveCore(path, me.Min, nil, me.Min.Sid); val != nil { return }
	}
	if me.Max != nil {
	if val = sidResolveCore(path, me.Max, nil, me.Max.Sid); val != nil { return }
	}
	return
}

func (me *KxEffector) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *NodeDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Insts.Camera { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Insts.Controller { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.Insts.Geometry { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.Insts.Light { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Nodes { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *FxTechniqueCommonConstant) resolveSidPath(path []string) (val interface{}) {
	if me.Emission != nil {
	if val = me.Emission.resolveSidPath(path); val != nil { return }
	}
	if me.Reflective != nil {
	if val = me.Reflective.resolveSidPath(path); val != nil { return }
	}
	if me.Reflectivity != nil {
	if val = me.Reflectivity.resolveSidPath(path); val != nil { return }
	}
	if me.Transparent != nil {
	if val = me.Transparent.resolveSidPath(path); val != nil { return }
	}
	if me.Transparency != nil {
	if val = me.Transparency.resolveSidPath(path); val != nil { return }
	}
	if me.IndexOfRefraction != nil {
	if val = me.IndexOfRefraction.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *ParamOrSidFloat) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.F, nil, me.F.Sid); val != nil { return }
	return
}

func (me *PxModelDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.RigidBodies { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.RigidConstraints { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.Insts { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *MaterialBinding) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Params { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.TC.Materials { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxAttachment) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	if me.Link != nil {
	if val = sidResolveCore(path, me.Link, me.Link, me.Link.Sid); val != nil { return }
	}
	return
}

func (me *GeometrySpline) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *PxShape) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Hollow, nil, me.Hollow.Sid); val != nil { return }
	if me.Mass != nil {
	if val = sidResolveCore(path, me.Mass, nil, me.Mass.Sid); val != nil { return }
	}
	if me.Density != nil {
	if val = sidResolveCore(path, me.Density, nil, me.Density.Sid); val != nil { return }
	}
	if val = me.Material.resolveSidPath(path); val != nil { return }
	if me.Geometry.Inst != nil {
	if val = sidResolveCore(path, me.Geometry.Inst, me.Geometry.Inst, me.Geometry.Inst.Sid); val != nil { return }
	}
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *CameraDef) resolveSidPath(path []string) (val interface{}) {
	if val = me.Optics.resolveSidPath(path); val != nil { return }
	return
}

func (me *Formula) resolveSidPath(path []string) (val interface{}) {
	if me.Inst != nil {
	if val = sidResolveCore(path, me.Inst, nil, me.Inst.Sid); val != nil { return }
	}
	return
}

func (me *PxRigidBodyInst) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.TC.Dynamic, nil, me.TC.Dynamic.Sid); val != nil { return }
	if me.TC.Mass != nil {
	if val = sidResolveCore(path, me.TC.Mass, nil, me.TC.Mass.Sid); val != nil { return }
	}
	for _, v := range me.TC.MassFrame { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	if me.TC.Inertia != nil {
	if val = sidResolveCore(path, me.TC.Inertia, nil, me.TC.Inertia.Sid); val != nil { return }
	}
	if val = me.TC.Material.resolveSidPath(path); val != nil { return }
	for _, v := range me.TC.Shapes { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *KxMotionAxis) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *GeometryBrepSweptSurface) resolveSidPath(path []string) (val interface{}) {
	if me.Curve != nil {
	if val = sidResolveCore(path, me.Curve, me.Curve, me.Curve.Sid); val != nil { return }
	}
	return
}

func (me *ControllerMorph) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *KxKinematicsAxis) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Formulas { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *KxMotionSystem) resolveSidPath(path []string) (val interface{}) {
	if me.ArticulatedSystem != nil {
	if val = sidResolveCore(path, me.ArticulatedSystem, me.ArticulatedSystem, me.ArticulatedSystem.Sid); val != nil { return }
	}
	for _, v := range me.TC.AxisInfos { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	if me.TC.EffectorInfo != nil {
	if val = sidResolveCore(path, me.TC.EffectorInfo, me.TC.EffectorInfo, me.TC.EffectorInfo.Sid); val != nil { return }
	}
	return
}

func (me *VisualSceneDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Nodes { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	for _, v := range me.Evaluations { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *LightPoint) resolveSidPath(path []string) (val interface{}) {
	if val = me.Attenuation.resolveSidPath(path); val != nil { return }
	return
}

func (me *FxTechniqueCommonLambert) resolveSidPath(path []string) (val interface{}) {
	if me.Emission != nil {
	if val = me.Emission.resolveSidPath(path); val != nil { return }
	}
	if me.Reflective != nil {
	if val = me.Reflective.resolveSidPath(path); val != nil { return }
	}
	if me.Reflectivity != nil {
	if val = me.Reflectivity.resolveSidPath(path); val != nil { return }
	}
	if me.Transparent != nil {
	if val = me.Transparent.resolveSidPath(path); val != nil { return }
	}
	if me.Transparency != nil {
	if val = me.Transparency.resolveSidPath(path); val != nil { return }
	}
	if me.IndexOfRefraction != nil {
	if val = me.IndexOfRefraction.resolveSidPath(path); val != nil { return }
	}
	if me.Ambient != nil {
	if val = me.Ambient.resolveSidPath(path); val != nil { return }
	}
	if me.Diffuse != nil {
	if val = me.Diffuse.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *PxRigidBodyDef) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.TC.Dynamic, nil, me.TC.Dynamic.Sid); val != nil { return }
	if me.TC.Mass != nil {
	if val = sidResolveCore(path, me.TC.Mass, nil, me.TC.Mass.Sid); val != nil { return }
	}
	for _, v := range me.TC.MassFrame { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	if me.TC.Inertia != nil {
	if val = sidResolveCore(path, me.TC.Inertia, nil, me.TC.Inertia.Sid); val != nil { return }
	}
	if val = me.TC.Material.resolveSidPath(path); val != nil { return }
	for _, v := range me.TC.Shapes { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *GeometryInst) resolveSidPath(path []string) (val interface{}) {
	if me.MaterialBinding != nil {
	if val = me.MaterialBinding.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *LightSpot) resolveSidPath(path []string) (val interface{}) {
	if val = me.Attenuation.resolveSidPath(path); val != nil { return }
	if val = sidResolveCore(path, &me.Falloff.Angle, nil, me.Falloff.Angle.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.Falloff.Exponent, nil, me.Falloff.Exponent.Sid); val != nil { return }
	return
}

func (me *KxModelInst) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxArticulatedSystemDef) resolveSidPath(path []string) (val interface{}) {
	if me.Kinematics != nil {
	if val = me.Kinematics.resolveSidPath(path); val != nil { return }
	}
	if me.Motion != nil {
	if val = me.Motion.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *KxFrame) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *GeometryBrep) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	if me.Curves != nil {
	if val = me.Curves.resolveSidPath(path); val != nil { return }
	}
	if me.SurfaceCurves != nil {
	if val = me.SurfaceCurves.resolveSidPath(path); val != nil { return }
	}
	if me.Surfaces != nil {
	if val = me.Surfaces.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *CameraOrthographic) resolveSidPath(path []string) (val interface{}) {
	if me.MagX != nil {
	if val = sidResolveCore(path, me.MagX, nil, me.MagX.Sid); val != nil { return }
	}
	if me.MagY != nil {
	if val = sidResolveCore(path, me.MagY, nil, me.MagY.Sid); val != nil { return }
	}
	return
}

func (me *FxMaterialDef) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Effect, nil, me.Effect.Sid); val != nil { return }
	return
}

func (me *FxTechniqueGlsl) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Passes { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *FxProfileGlsl) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.CodesIncludes { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Techniques { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *KxSceneInst) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxArticulatedSystemInst) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *KxFrameObject) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *CameraPerspective) resolveSidPath(path []string) (val interface{}) {
	if me.FovX != nil {
	if val = sidResolveCore(path, me.FovX, nil, me.FovX.Sid); val != nil { return }
	}
	if me.FovY != nil {
	if val = sidResolveCore(path, me.FovY, nil, me.FovY.Sid); val != nil { return }
	}
	return
}

func (me *FxTechniqueCommon) resolveSidPath(path []string) (val interface{}) {
	if me.Blinn != nil {
	if val = me.Blinn.resolveSidPath(path); val != nil { return }
	}
	if me.Constant != nil {
	if val = me.Constant.resolveSidPath(path); val != nil { return }
	}
	if me.Lambert != nil {
	if val = me.Lambert.resolveSidPath(path); val != nil { return }
	}
	if me.Phong != nil {
	if val = me.Phong.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *ControllerInst) resolveSidPath(path []string) (val interface{}) {
	if me.BindMaterial != nil {
	if val = me.BindMaterial.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *ControllerSkin) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Sources { if val = v.resolveSidPath(path); val != nil { return } }
	return
}

func (me *GeometryBrepSurfaceCurves) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.All { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	return
}

func (me *PxRigidConstraintDef) resolveSidPath(path []string) (val interface{}) {
	if val = me.RefAttachment.resolveSidPath(path); val != nil { return }
	if val = me.Attachment.resolveSidPath(path); val != nil { return }
	if val = sidResolveCore(path, &me.TC.Enabled, nil, me.TC.Enabled.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.TC.Interpenetrate, nil, me.TC.Interpenetrate.Sid); val != nil { return }
	if me.TC.Limits.Angular != nil {
	if val = me.TC.Limits.Angular.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Limits.Linear != nil {
	if val = me.TC.Limits.Linear.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Spring.Angular != nil {
	if val = me.TC.Spring.Angular.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Spring.Linear != nil {
	if val = me.TC.Spring.Linear.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *PxRigidConstraintLimit) resolveSidPath(path []string) (val interface{}) {
	if val = sidResolveCore(path, &me.Min, nil, me.Min.Sid); val != nil { return }
	if val = sidResolveCore(path, &me.Max, nil, me.Max.Sid); val != nil { return }
	return
}

func (me *KxFrameTip) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Transforms { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *FormulaDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *VisualSceneEvaluation) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.RenderPasses { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

func (me *FxTechniqueCommonBlinn) resolveSidPath(path []string) (val interface{}) {
	if me.Emission != nil {
	if val = me.Emission.resolveSidPath(path); val != nil { return }
	}
	if me.Reflective != nil {
	if val = me.Reflective.resolveSidPath(path); val != nil { return }
	}
	if me.Reflectivity != nil {
	if val = me.Reflectivity.resolveSidPath(path); val != nil { return }
	}
	if me.Transparent != nil {
	if val = me.Transparent.resolveSidPath(path); val != nil { return }
	}
	if me.Transparency != nil {
	if val = me.Transparency.resolveSidPath(path); val != nil { return }
	}
	if me.IndexOfRefraction != nil {
	if val = me.IndexOfRefraction.resolveSidPath(path); val != nil { return }
	}
	if me.Ambient != nil {
	if val = me.Ambient.resolveSidPath(path); val != nil { return }
	}
	if me.Diffuse != nil {
	if val = me.Diffuse.resolveSidPath(path); val != nil { return }
	}
	if me.Specular != nil {
	if val = me.Specular.resolveSidPath(path); val != nil { return }
	}
	if me.Shininess != nil {
	if val = me.Shininess.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *FxProfile) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.NewParams { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	if me.Common != nil {
	if val = me.Common.resolveSidPath(path); val != nil { return }
	}
	if me.Glsl != nil {
	if val = me.Glsl.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *Source) resolveSidPath(path []string) (val interface{}) {
	if me.TC.Accessor != nil {
	if val = me.TC.Accessor.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *FxPassEvaluation) resolveSidPath(path []string) (val interface{}) {
	if me.Color.Target != nil {
	if val = me.Color.Target.resolveSidPath(path); val != nil { return }
	}
	if me.Depth.Target != nil {
	if val = me.Depth.Target.resolveSidPath(path); val != nil { return }
	}
	if me.Stencil.Target != nil {
	if val = me.Stencil.Target.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *LightDef) resolveSidPath(path []string) (val interface{}) {
	if me.TC.Point != nil {
	if val = me.TC.Point.resolveSidPath(path); val != nil { return }
	}
	if me.TC.Spot != nil {
	if val = me.TC.Spot.resolveSidPath(path); val != nil { return }
	}
	return
}

func (me *PxMaterial) resolveSidPath(path []string) (val interface{}) {
	if me.Inst != nil {
	if val = sidResolveCore(path, me.Inst, nil, me.Inst.Sid); val != nil { return }
	}
	return
}

func (me *PxSceneDef) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.ForceFields { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	for _, v := range me.Models { if val = sidResolveCore(path, v, v, v.Sid); val != nil { return } }
	if me.TC.Gravity != nil {
	if val = sidResolveCore(path, me.TC.Gravity, nil, me.TC.Gravity.Sid); val != nil { return }
	}
	if me.TC.TimeStep != nil {
	if val = sidResolveCore(path, me.TC.TimeStep, nil, me.TC.TimeStep.Sid); val != nil { return }
	}
	return
}

func (me *SourceAccessor) resolveSidPath(path []string) (val interface{}) {
	for _, v := range me.Params { if val = sidResolveCore(path, v, nil, v.Sid); val != nil { return } }
	return
}

