package collimp

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"

	nga "github.com/go3d/go-ngine/assets"
)

func obj_ScopedString(xn *xmlx.Node, n string) (obj *nga.ScopedString) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ScopedString)
	}
	return
}

func objs_ScopedString(xn *xmlx.Node, n string) (objs []*nga.ScopedString) {
	xns := xsns(xn, n)
	objs = make([]*nga.ScopedString, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ScopedString(xn, "")
	}
	return
}

func obj_Sources(xn *xmlx.Node, n string) (obj *nga.Sources) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Sources)
	}
	return
}

func objs_Sources(xn *xmlx.Node, n string) (objs []*nga.Sources) {
	xns := xsns(xn, n)
	objs = make([]*nga.Sources, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Sources(xn, "")
	}
	return
}

func obj_FxPassProgram(xn *xmlx.Node, n string) (obj *nga.FxPassProgram) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassProgram)
	}
	return
}

func objs_FxPassProgram(xn *xmlx.Node, n string) (objs []*nga.FxPassProgram) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassProgram, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgram(xn, "")
	}
	return
}

func obj_LightAttenuation(xn *xmlx.Node, n string) (obj *nga.LightAttenuation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightAttenuation)
	}
	return
}

func objs_LightAttenuation(xn *xmlx.Node, n string) (objs []*nga.LightAttenuation) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightAttenuation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightAttenuation(xn, "")
	}
	return
}

func obj_LibsControllerDef(xn *xmlx.Node, n string) (obj *nga.LibsControllerDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsControllerDef)
	}
	return
}

func objs_LibsControllerDef(xn *xmlx.Node, n string) (objs []*nga.LibsControllerDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsControllerDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsControllerDef(xn, "")
	}
	return
}

func obj_KxModelInst(xn *xmlx.Node, n string) (obj *nga.KxModelInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxModelInst)
	}
	return
}

func objs_KxModelInst(xn *xmlx.Node, n string) (objs []*nga.KxModelInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxModelInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxModelInst(xn, "")
	}
	return
}

func obj_GeometryBrepPlane(xn *xmlx.Node, n string) (obj *nga.GeometryBrepPlane) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepPlane)
	}
	return
}

func objs_GeometryBrepPlane(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepPlane) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepPlane, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepPlane(xn, "")
	}
	return
}

func obj_ParamString(xn *xmlx.Node, n string) (obj *nga.ParamString) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamString)
	}
	return
}

func objs_ParamString(xn *xmlx.Node, n string) (objs []*nga.ParamString) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamString, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamString(xn, "")
	}
	return
}

func obj_KxAttachment(xn *xmlx.Node, n string) (obj *nga.KxAttachment) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxAttachment)
	}
	return
}

func objs_KxAttachment(xn *xmlx.Node, n string) (objs []*nga.KxAttachment) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxAttachment, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxAttachment(xn, "")
	}
	return
}

func obj_VisualSceneInst(xn *xmlx.Node, n string) (obj *nga.VisualSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.VisualSceneInst)
	}
	return
}

func objs_VisualSceneInst(xn *xmlx.Node, n string) (objs []*nga.VisualSceneInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.VisualSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneInst(xn, "")
	}
	return
}

func obj_FxCreateMips(xn *xmlx.Node, n string) (obj *nga.FxCreateMips) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateMips)
	}
	return
}

func objs_FxCreateMips(xn *xmlx.Node, n string) (objs []*nga.FxCreateMips) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateMips, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateMips(xn, "")
	}
	return
}

func obj_VisualSceneRendering(xn *xmlx.Node, n string) (obj *nga.VisualSceneRendering) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewVisualSceneRendering()
	}
	return
}

func objs_VisualSceneRendering(xn *xmlx.Node, n string) (objs []*nga.VisualSceneRendering) {
	xns := xsns(xn, n)
	objs = make([]*nga.VisualSceneRendering, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneRendering(xn, "")
	}
	return
}

func obj_FxCreateFormat(xn *xmlx.Node, n string) (obj *nga.FxCreateFormat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateFormat)
	}
	return
}

func objs_FxCreateFormat(xn *xmlx.Node, n string) (objs []*nga.FxCreateFormat) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateFormat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateFormat(xn, "")
	}
	return
}

func obj_SourceAccessor(xn *xmlx.Node, n string) (obj *nga.SourceAccessor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.SourceAccessor)
	}
	return
}

func objs_SourceAccessor(xn *xmlx.Node, n string) (objs []*nga.SourceAccessor) {
	xns := xsns(xn, n)
	objs = make([]*nga.SourceAccessor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SourceAccessor(xn, "")
	}
	return
}

func obj_PxForceFieldDef(xn *xmlx.Node, n string) (obj *nga.PxForceFieldDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxForceFieldDef)
	}
	return
}

func objs_PxForceFieldDef(xn *xmlx.Node, n string) (objs []*nga.PxForceFieldDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxForceFieldDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxForceFieldDef(xn, "")
	}
	return
}

func obj_PxModelInst(xn *xmlx.Node, n string) (obj *nga.PxModelInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxModelInst)
	}
	return
}

func objs_PxModelInst(xn *xmlx.Node, n string) (objs []*nga.PxModelInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxModelInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxModelInst(xn, "")
	}
	return
}

func obj_FxTechniqueCommonConstant(xn *xmlx.Node, n string) (obj *nga.FxTechniqueCommonConstant) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueCommonConstant)
	}
	return
}

func objs_FxTechniqueCommonConstant(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueCommonConstant) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueCommonConstant, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommonConstant(xn, "")
	}
	return
}

func obj_LibFxMaterialDefs(xn *xmlx.Node, n string) (obj *nga.LibFxMaterialDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibFxMaterialDefs)
	}
	return
}

func objs_LibFxMaterialDefs(xn *xmlx.Node, n string) (objs []*nga.LibFxMaterialDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibFxMaterialDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibFxMaterialDefs(xn, "")
	}
	return
}

func obj_KxArticulatedSystemAxisLimits(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemAxisLimits) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemAxisLimits)
	}
	return
}

func objs_KxArticulatedSystemAxisLimits(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemAxisLimits) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemAxisLimits, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemAxisLimits(xn, "")
	}
	return
}

func obj_Technique(xn *xmlx.Node, n string) (obj *nga.Technique) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Technique)
	}
	return
}

func objs_Technique(xn *xmlx.Node, n string) (objs []*nga.Technique) {
	xns := xsns(xn, n)
	objs = make([]*nga.Technique, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Technique(xn, "")
	}
	return
}

func obj_KxSceneInstJointAxis(xn *xmlx.Node, n string) (obj *nga.KxSceneInstJointAxis) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxSceneInstJointAxis)
	}
	return
}

func objs_KxSceneInstJointAxis(xn *xmlx.Node, n string) (objs []*nga.KxSceneInstJointAxis) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxSceneInstJointAxis, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneInstJointAxis(xn, "")
	}
	return
}

func obj_GeometryInst(xn *xmlx.Node, n string) (obj *nga.GeometryInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryInst)
	}
	return
}

func objs_GeometryInst(xn *xmlx.Node, n string) (objs []*nga.GeometryInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryInst(xn, "")
	}
	return
}

func obj_PxRigidConstraintInst(xn *xmlx.Node, n string) (obj *nga.PxRigidConstraintInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidConstraintInst)
	}
	return
}

func objs_PxRigidConstraintInst(xn *xmlx.Node, n string) (objs []*nga.PxRigidConstraintInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidConstraintInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintInst(xn, "")
	}
	return
}

func obj_KxArticulatedSystemKinematics(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemKinematics) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemKinematics)
	}
	return
}

func objs_KxArticulatedSystemKinematics(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemKinematics) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemKinematics, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemKinematics(xn, "")
	}
	return
}

func obj_ControllerInst(xn *xmlx.Node, n string) (obj *nga.ControllerInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ControllerInst)
	}
	return
}

func objs_ControllerInst(xn *xmlx.Node, n string) (objs []*nga.ControllerInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.ControllerInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerInst(xn, "")
	}
	return
}

func obj_NodeDef(xn *xmlx.Node, n string) (obj *nga.NodeDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.NodeDef)
	}
	return
}

func objs_NodeDef(xn *xmlx.Node, n string) (objs []*nga.NodeDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.NodeDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_NodeDef(xn, "")
	}
	return
}

func obj_FxProfileGlSlCodeInclude(xn *xmlx.Node, n string) (obj *nga.FxProfileGlSlCodeInclude) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxProfileGlSlCodeInclude)
	}
	return
}

func objs_FxProfileGlSlCodeInclude(xn *xmlx.Node, n string) (objs []*nga.FxProfileGlSlCodeInclude) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxProfileGlSlCodeInclude, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileGlSlCodeInclude(xn, "")
	}
	return
}

func obj_GeometryBrepCurve(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCurve) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCurve)
	}
	return
}

func objs_GeometryBrepCurve(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCurve) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCurve, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCurve(xn, "")
	}
	return
}

func obj_AssetGeographicLocation(xn *xmlx.Node, n string) (obj *nga.AssetGeographicLocation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AssetGeographicLocation)
	}
	return
}

func objs_AssetGeographicLocation(xn *xmlx.Node, n string) (objs []*nga.AssetGeographicLocation) {
	xns := xsns(xn, n)
	objs = make([]*nga.AssetGeographicLocation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AssetGeographicLocation(xn, "")
	}
	return
}

func obj_FxPassEvaluationTarget(xn *xmlx.Node, n string) (obj *nga.FxPassEvaluationTarget) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassEvaluationTarget)
	}
	return
}

func objs_FxPassEvaluationTarget(xn *xmlx.Node, n string) (objs []*nga.FxPassEvaluationTarget) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassEvaluationTarget, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationTarget(xn, "")
	}
	return
}

func obj_KxArticulatedSystemMotionAxis(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemMotionAxis) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewKxArticulatedSystemMotionAxis()
	}
	return
}

func objs_KxArticulatedSystemMotionAxis(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemMotionAxis) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemMotionAxis, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemMotionAxis(xn, "")
	}
	return
}

func obj_FxTexture(xn *xmlx.Node, n string) (obj *nga.FxTexture) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTexture)
	}
	return
}

func objs_FxTexture(xn *xmlx.Node, n string) (objs []*nga.FxTexture) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTexture, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTexture(xn, "")
	}
	return
}

func obj_GeometryBrepCircle(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCircle) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCircle)
	}
	return
}

func objs_GeometryBrepCircle(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCircle) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCircle, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCircle(xn, "")
	}
	return
}

func obj_GeometryBrepSolids(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSolids) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSolids)
	}
	return
}

func objs_GeometryBrepSolids(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSolids) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSolids, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSolids(xn, "")
	}
	return
}

func obj_KxSceneDef(xn *xmlx.Node, n string) (obj *nga.KxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxSceneDef)
	}
	return
}

func objs_KxSceneDef(xn *xmlx.Node, n string) (objs []*nga.KxSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneDef(xn, "")
	}
	return
}

func obj_FxCreateInitFrom(xn *xmlx.Node, n string) (obj *nga.FxCreateInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateInitFrom)
	}
	return
}

func objs_FxCreateInitFrom(xn *xmlx.Node, n string) (objs []*nga.FxCreateInitFrom) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateInitFrom(xn, "")
	}
	return
}

func obj_KxJointInst(xn *xmlx.Node, n string) (obj *nga.KxJointInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxJointInst)
	}
	return
}

func objs_KxJointInst(xn *xmlx.Node, n string) (objs []*nga.KxJointInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxJointInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointInst(xn, "")
	}
	return
}

func obj_FxProfileGlSl(xn *xmlx.Node, n string) (obj *nga.FxProfileGlSl) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewFxProfileGlSl()
	}
	return
}

func objs_FxProfileGlSl(xn *xmlx.Node, n string) (objs []*nga.FxProfileGlSl) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxProfileGlSl, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileGlSl(xn, "")
	}
	return
}

func obj_GeometryBrep(xn *xmlx.Node, n string) (obj *nga.GeometryBrep) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewGeometryBrep()
	}
	return
}

func objs_GeometryBrep(xn *xmlx.Node, n string) (objs []*nga.GeometryBrep) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrep, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrep(xn, "")
	}
	return
}

func obj_GeometryDef(xn *xmlx.Node, n string) (obj *nga.GeometryDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryDef)
	}
	return
}

func objs_GeometryDef(xn *xmlx.Node, n string) (objs []*nga.GeometryDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryDef(xn, "")
	}
	return
}

func obj_ScopedBool(xn *xmlx.Node, n string) (obj *nga.ScopedBool) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ScopedBool)
	}
	return
}

func objs_ScopedBool(xn *xmlx.Node, n string) (objs []*nga.ScopedBool) {
	xns := xsns(xn, n)
	objs = make([]*nga.ScopedBool, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ScopedBool(xn, "")
	}
	return
}

func obj_KxArticulatedSystemDef(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemDef)
	}
	return
}

func objs_KxArticulatedSystemDef(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemDef(xn, "")
	}
	return
}

func obj_ControllerMorph(xn *xmlx.Node, n string) (obj *nga.ControllerMorph) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewControllerMorph()
	}
	return
}

func objs_ControllerMorph(xn *xmlx.Node, n string) (objs []*nga.ControllerMorph) {
	xns := xsns(xn, n)
	objs = make([]*nga.ControllerMorph, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerMorph(xn, "")
	}
	return
}

func obj_KxArticulatedSystemMotion(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemMotion) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemMotion)
	}
	return
}

func objs_KxArticulatedSystemMotion(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemMotion) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemMotion, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemMotion(xn, "")
	}
	return
}

func obj_ParamScopedFloat(xn *xmlx.Node, n string) (obj *nga.ParamScopedFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamScopedFloat)
	}
	return
}

func objs_ParamScopedFloat(xn *xmlx.Node, n string) (objs []*nga.ParamScopedFloat) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamScopedFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamScopedFloat(xn, "")
	}
	return
}

func obj_KxArticulatedSystemInst(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemInst)
	}
	return
}

func objs_KxArticulatedSystemInst(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemInst(xn, "")
	}
	return
}

func obj_MeshVert(xn *xmlx.Node, n string) (obj *nga.MeshVert) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshVert)
	}
	return
}

func objs_MeshVert(xn *xmlx.Node, n string) (objs []*nga.MeshVert) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshVert, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshVert(xn, "")
	}
	return
}

func obj_ListStrings(xn *xmlx.Node, n string) (obj *nga.ListStrings) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ListStrings)
	}
	return
}

func objs_ListStrings(xn *xmlx.Node, n string) (objs []*nga.ListStrings) {
	xns := xsns(xn, n)
	objs = make([]*nga.ListStrings, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ListStrings(xn, "")
	}
	return
}

func obj_AnimationInst(xn *xmlx.Node, n string) (obj *nga.AnimationInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AnimationInst)
	}
	return
}

func objs_AnimationInst(xn *xmlx.Node, n string) (objs []*nga.AnimationInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.AnimationInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationInst(xn, "")
	}
	return
}

func obj_IndexedInputs(xn *xmlx.Node, n string) (obj *nga.IndexedInputs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.IndexedInputs)
	}
	return
}

func objs_IndexedInputs(xn *xmlx.Node, n string) (objs []*nga.IndexedInputs) {
	xns := xsns(xn, n)
	objs = make([]*nga.IndexedInputs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_IndexedInputs(xn, "")
	}
	return
}

func obj_LibsVisualSceneDef(xn *xmlx.Node, n string) (obj *nga.LibsVisualSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsVisualSceneDef)
	}
	return
}

func objs_LibsVisualSceneDef(xn *xmlx.Node, n string) (objs []*nga.LibsVisualSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsVisualSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsVisualSceneDef(xn, "")
	}
	return
}

func obj_FxSamplerWrapping(xn *xmlx.Node, n string) (obj *nga.FxSamplerWrapping) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxSamplerWrapping)
	}
	return
}

func objs_FxSamplerWrapping(xn *xmlx.Node, n string) (objs []*nga.FxSamplerWrapping) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxSamplerWrapping, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerWrapping(xn, "")
	}
	return
}

func obj_LibFormulaDefs(xn *xmlx.Node, n string) (obj *nga.LibFormulaDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibFormulaDefs)
	}
	return
}

func objs_LibFormulaDefs(xn *xmlx.Node, n string) (objs []*nga.LibFormulaDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibFormulaDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibFormulaDefs(xn, "")
	}
	return
}

func obj_LibsFxMaterialDef(xn *xmlx.Node, n string) (obj *nga.LibsFxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsFxMaterialDef)
	}
	return
}

func objs_LibsFxMaterialDef(xn *xmlx.Node, n string) (objs []*nga.LibsFxMaterialDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsFxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsFxMaterialDef(xn, "")
	}
	return
}

func obj_GeometryBrepNurbs(xn *xmlx.Node, n string) (obj *nga.GeometryBrepNurbs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewGeometryBrepNurbs()
	}
	return
}

func objs_GeometryBrepNurbs(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepNurbs) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepNurbs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepNurbs(xn, "")
	}
	return
}

func obj_LightAmbient(xn *xmlx.Node, n string) (obj *nga.LightAmbient) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightAmbient)
	}
	return
}

func objs_LightAmbient(xn *xmlx.Node, n string) (objs []*nga.LightAmbient) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightAmbient, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightAmbient(xn, "")
	}
	return
}

func obj_Int4(xn *xmlx.Node, n string) (obj *nga.Int4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int4)
	}
	return
}

func objs_Int4(xn *xmlx.Node, n string) (objs []*nga.Int4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int4(xn, "")
	}
	return
}

func obj_LibVisualSceneDefs(xn *xmlx.Node, n string) (obj *nga.LibVisualSceneDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibVisualSceneDefs)
	}
	return
}

func objs_LibVisualSceneDefs(xn *xmlx.Node, n string) (objs []*nga.LibVisualSceneDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibVisualSceneDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibVisualSceneDefs(xn, "")
	}
	return
}

func obj_GeometryBrepSweptSurface(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSweptSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSweptSurface)
	}
	return
}

func objs_GeometryBrepSweptSurface(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSweptSurface) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSweptSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSweptSurface(xn, "")
	}
	return
}

func obj_GeometryBrepLine(xn *xmlx.Node, n string) (obj *nga.GeometryBrepLine) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepLine)
	}
	return
}

func objs_GeometryBrepLine(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepLine) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepLine, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepLine(xn, "")
	}
	return
}

func obj_FxTechniqueCommon(xn *xmlx.Node, n string) (obj *nga.FxTechniqueCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueCommon)
	}
	return
}

func objs_FxTechniqueCommon(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueCommon) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommon(xn, "")
	}
	return
}

func obj_PxRigidConstraintSpring(xn *xmlx.Node, n string) (obj *nga.PxRigidConstraintSpring) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidConstraintSpring)
	}
	return
}

func objs_PxRigidConstraintSpring(xn *xmlx.Node, n string) (objs []*nga.PxRigidConstraintSpring) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidConstraintSpring, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintSpring(xn, "")
	}
	return
}

func obj_GeometryBrepSphere(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSphere) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSphere)
	}
	return
}

func objs_GeometryBrepSphere(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSphere) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSphere, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSphere(xn, "")
	}
	return
}

func obj_KxModelDef(xn *xmlx.Node, n string) (obj *nga.KxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxModelDef)
	}
	return
}

func objs_KxModelDef(xn *xmlx.Node, n string) (objs []*nga.KxModelDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxModelDef(xn, "")
	}
	return
}

func obj_GeometryBrepSurfaces(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSurfaces) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSurfaces)
	}
	return
}

func objs_GeometryBrepSurfaces(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSurfaces) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSurfaces, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurfaces(xn, "")
	}
	return
}

func obj_FxImageDef(xn *xmlx.Node, n string) (obj *nga.FxImageDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxImageDef)
	}
	return
}

func objs_FxImageDef(xn *xmlx.Node, n string) (objs []*nga.FxImageDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxImageDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageDef(xn, "")
	}
	return
}

func obj_MeshRawFace(xn *xmlx.Node, n string) (obj *nga.MeshRawFace) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewMeshRawFace()
	}
	return
}

func objs_MeshRawFace(xn *xmlx.Node, n string) (objs []*nga.MeshRawFace) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshRawFace, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshRawFace(xn, "")
	}
	return
}

func obj_VisualSceneDef(xn *xmlx.Node, n string) (obj *nga.VisualSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.VisualSceneDef)
	}
	return
}

func objs_VisualSceneDef(xn *xmlx.Node, n string) (objs []*nga.VisualSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.VisualSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneDef(xn, "")
	}
	return
}

func obj_Float4x3(xn *xmlx.Node, n string) (obj *nga.Float4x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float4x3)
	}
	return
}

func objs_Float4x3(xn *xmlx.Node, n string) (objs []*nga.Float4x3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float4x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x3(xn, "")
	}
	return
}

func obj_LibsKxSceneDef(xn *xmlx.Node, n string) (obj *nga.LibsKxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsKxSceneDef)
	}
	return
}

func objs_LibsKxSceneDef(xn *xmlx.Node, n string) (objs []*nga.LibsKxSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsKxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsKxSceneDef(xn, "")
	}
	return
}

func obj_Transform(xn *xmlx.Node, n string) (obj *nga.Transform) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Transform)
	}
	return
}

func objs_Transform(xn *xmlx.Node, n string) (objs []*nga.Transform) {
	xns := xsns(xn, n)
	objs = make([]*nga.Transform, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Transform(xn, "")
	}
	return
}

func obj_GeometryBrepCylinder(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCylinder) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCylinder)
	}
	return
}

func objs_GeometryBrepCylinder(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCylinder) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCylinder, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCylinder(xn, "")
	}
	return
}

func obj_LibPxMaterialDefs(xn *xmlx.Node, n string) (obj *nga.LibPxMaterialDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibPxMaterialDefs)
	}
	return
}

func objs_LibPxMaterialDefs(xn *xmlx.Node, n string) (objs []*nga.LibPxMaterialDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibPxMaterialDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibPxMaterialDefs(xn, "")
	}
	return
}

func obj_PxSceneInst(xn *xmlx.Node, n string) (obj *nga.PxSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxSceneInst)
	}
	return
}

func objs_PxSceneInst(xn *xmlx.Node, n string) (objs []*nga.PxSceneInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxSceneInst(xn, "")
	}
	return
}

func obj_LibsKxArticulatedSystemDef(xn *xmlx.Node, n string) (obj *nga.LibsKxArticulatedSystemDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsKxArticulatedSystemDef)
	}
	return
}

func objs_LibsKxArticulatedSystemDef(xn *xmlx.Node, n string) (objs []*nga.LibsKxArticulatedSystemDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsKxArticulatedSystemDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsKxArticulatedSystemDef(xn, "")
	}
	return
}

func obj_GeometryBrepFaces(xn *xmlx.Node, n string) (obj *nga.GeometryBrepFaces) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepFaces)
	}
	return
}

func objs_GeometryBrepFaces(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepFaces) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepFaces, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepFaces(xn, "")
	}
	return
}

func obj_LibsAnimationClipDef(xn *xmlx.Node, n string) (obj *nga.LibsAnimationClipDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsAnimationClipDef)
	}
	return
}

func objs_LibsAnimationClipDef(xn *xmlx.Node, n string) (objs []*nga.LibsAnimationClipDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsAnimationClipDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsAnimationClipDef(xn, "")
	}
	return
}

func obj_FxTechniqueCommonPhong(xn *xmlx.Node, n string) (obj *nga.FxTechniqueCommonPhong) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueCommonPhong)
	}
	return
}

func objs_FxTechniqueCommonPhong(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueCommonPhong) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueCommonPhong, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommonPhong(xn, "")
	}
	return
}

func obj_LibKxArticulatedSystemDefs(xn *xmlx.Node, n string) (obj *nga.LibKxArticulatedSystemDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibKxArticulatedSystemDefs)
	}
	return
}

func objs_LibKxArticulatedSystemDefs(xn *xmlx.Node, n string) (objs []*nga.LibKxArticulatedSystemDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibKxArticulatedSystemDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibKxArticulatedSystemDefs(xn, "")
	}
	return
}

func obj_ListBools(xn *xmlx.Node, n string) (obj *nga.ListBools) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ListBools)
	}
	return
}

func objs_ListBools(xn *xmlx.Node, n string) (objs []*nga.ListBools) {
	xns := xsns(xn, n)
	objs = make([]*nga.ListBools, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ListBools(xn, "")
	}
	return
}

func obj_FxMaterialInst(xn *xmlx.Node, n string) (obj *nga.FxMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxMaterialInst)
	}
	return
}

func objs_FxMaterialInst(xn *xmlx.Node, n string) (objs []*nga.FxMaterialInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialInst(xn, "")
	}
	return
}

func obj_GeometryBrepPcurves(xn *xmlx.Node, n string) (obj *nga.GeometryBrepPcurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepPcurves)
	}
	return
}

func objs_GeometryBrepPcurves(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepPcurves) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepPcurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepPcurves(xn, "")
	}
	return
}

func obj_FxImageInst(xn *xmlx.Node, n string) (obj *nga.FxImageInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxImageInst)
	}
	return
}

func objs_FxImageInst(xn *xmlx.Node, n string) (objs []*nga.FxImageInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxImageInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageInst(xn, "")
	}
	return
}

func obj_KxLink(xn *xmlx.Node, n string) (obj *nga.KxLink) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxLink)
	}
	return
}

func objs_KxLink(xn *xmlx.Node, n string) (objs []*nga.KxLink) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxLink, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxLink(xn, "")
	}
	return
}

func obj_LibKxModelDefs(xn *xmlx.Node, n string) (obj *nga.LibKxModelDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibKxModelDefs)
	}
	return
}

func objs_LibKxModelDefs(xn *xmlx.Node, n string) (objs []*nga.LibKxModelDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibKxModelDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibKxModelDefs(xn, "")
	}
	return
}

func obj_ParamInst(xn *xmlx.Node, n string) (obj *nga.ParamInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamInst)
	}
	return
}

func objs_ParamInst(xn *xmlx.Node, n string) (objs []*nga.ParamInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamInst(xn, "")
	}
	return
}

func obj_GeometryBrepShells(xn *xmlx.Node, n string) (obj *nga.GeometryBrepShells) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepShells)
	}
	return
}

func objs_GeometryBrepShells(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepShells) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepShells, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepShells(xn, "")
	}
	return
}

func obj_CameraOptics(xn *xmlx.Node, n string) (obj *nga.CameraOptics) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraOptics)
	}
	return
}

func objs_CameraOptics(xn *xmlx.Node, n string) (objs []*nga.CameraOptics) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraOptics, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraOptics(xn, "")
	}
	return
}

func obj_GeometryBrepTorus(xn *xmlx.Node, n string) (obj *nga.GeometryBrepTorus) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepTorus)
	}
	return
}

func objs_GeometryBrepTorus(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepTorus) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepTorus, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepTorus(xn, "")
	}
	return
}

func obj_FxColorOrTexture(xn *xmlx.Node, n string) (obj *nga.FxColorOrTexture) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxColorOrTexture)
	}
	return
}

func objs_FxColorOrTexture(xn *xmlx.Node, n string) (objs []*nga.FxColorOrTexture) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxColorOrTexture, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxColorOrTexture(xn, "")
	}
	return
}

func obj_Float3x3(xn *xmlx.Node, n string) (obj *nga.Float3x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float3x3)
	}
	return
}

func objs_Float3x3(xn *xmlx.Node, n string) (objs []*nga.Float3x3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float3x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x3(xn, "")
	}
	return
}

func obj_FxCreate2D(xn *xmlx.Node, n string) (obj *nga.FxCreate2D) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreate2D)
	}
	return
}

func objs_FxCreate2D(xn *xmlx.Node, n string) (objs []*nga.FxCreate2D) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreate2D, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2D(xn, "")
	}
	return
}

func obj_ParamInt(xn *xmlx.Node, n string) (obj *nga.ParamInt) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamInt)
	}
	return
}

func objs_ParamInt(xn *xmlx.Node, n string) (objs []*nga.ParamInt) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamInt, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamInt(xn, "")
	}
	return
}

func obj_PxRigidBodyCommon(xn *xmlx.Node, n string) (obj *nga.PxRigidBodyCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidBodyCommon)
	}
	return
}

func objs_PxRigidBodyCommon(xn *xmlx.Node, n string) (objs []*nga.PxRigidBodyCommon) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidBodyCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyCommon(xn, "")
	}
	return
}

func obj_FxSampler(xn *xmlx.Node, n string) (obj *nga.FxSampler) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewFxSampler()
	}
	return
}

func objs_FxSampler(xn *xmlx.Node, n string) (objs []*nga.FxSampler) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxSampler, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSampler(xn, "")
	}
	return
}

func obj_MeshRaw(xn *xmlx.Node, n string) (obj *nga.MeshRaw) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshRaw)
	}
	return
}

func objs_MeshRaw(xn *xmlx.Node, n string) (objs []*nga.MeshRaw) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshRaw, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshRaw(xn, "")
	}
	return
}

func obj_FxPassEvaluation(xn *xmlx.Node, n string) (obj *nga.FxPassEvaluation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassEvaluation)
	}
	return
}

func objs_FxPassEvaluation(xn *xmlx.Node, n string) (objs []*nga.FxPassEvaluation) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassEvaluation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluation(xn, "")
	}
	return
}

func obj_LightInst(xn *xmlx.Node, n string) (obj *nga.LightInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightInst)
	}
	return
}

func objs_LightInst(xn *xmlx.Node, n string) (objs []*nga.LightInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightInst(xn, "")
	}
	return
}

func obj_MeshData(xn *xmlx.Node, n string) (obj *nga.MeshData) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewMeshData()
	}
	return
}

func objs_MeshData(xn *xmlx.Node, n string) (objs []*nga.MeshData) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshData, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshData(xn, "")
	}
	return
}

func obj_GeometryBrepOrientation(xn *xmlx.Node, n string) (obj *nga.GeometryBrepOrientation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepOrientation)
	}
	return
}

func objs_GeometryBrepOrientation(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepOrientation) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepOrientation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepOrientation(xn, "")
	}
	return
}

func obj_Float4x2(xn *xmlx.Node, n string) (obj *nga.Float4x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float4x2)
	}
	return
}

func objs_Float4x2(xn *xmlx.Node, n string) (objs []*nga.Float4x2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float4x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x2(xn, "")
	}
	return
}

func obj_LibControllerDefs(xn *xmlx.Node, n string) (obj *nga.LibControllerDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibControllerDefs)
	}
	return
}

func objs_LibControllerDefs(xn *xmlx.Node, n string) (objs []*nga.LibControllerDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibControllerDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibControllerDefs(xn, "")
	}
	return
}

func obj_PxCylinder(xn *xmlx.Node, n string) (obj *nga.PxCylinder) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxCylinder)
	}
	return
}

func objs_PxCylinder(xn *xmlx.Node, n string) (objs []*nga.PxCylinder) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxCylinder, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxCylinder(xn, "")
	}
	return
}

func obj_GeometryBrepSurface(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSurface)
	}
	return
}

func objs_GeometryBrepSurface(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSurface) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurface(xn, "")
	}
	return
}

func obj_AnimationDef(xn *xmlx.Node, n string) (obj *nga.AnimationDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AnimationDef)
	}
	return
}

func objs_AnimationDef(xn *xmlx.Node, n string) (objs []*nga.AnimationDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.AnimationDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationDef(xn, "")
	}
	return
}

func obj_LibsLightDef(xn *xmlx.Node, n string) (obj *nga.LibsLightDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsLightDef)
	}
	return
}

func objs_LibsLightDef(xn *xmlx.Node, n string) (objs []*nga.LibsLightDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsLightDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsLightDef(xn, "")
	}
	return
}

func obj_GeometrySpline(xn *xmlx.Node, n string) (obj *nga.GeometrySpline) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewGeometrySpline()
	}
	return
}

func objs_GeometrySpline(xn *xmlx.Node, n string) (objs []*nga.GeometrySpline) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometrySpline, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometrySpline(xn, "")
	}
	return
}

func obj_GeometryPolygonHole(xn *xmlx.Node, n string) (obj *nga.GeometryPolygonHole) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryPolygonHole)
	}
	return
}

func objs_GeometryPolygonHole(xn *xmlx.Node, n string) (objs []*nga.GeometryPolygonHole) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryPolygonHole, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPolygonHole(xn, "")
	}
	return
}

func obj_Float3x2(xn *xmlx.Node, n string) (obj *nga.Float3x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float3x2)
	}
	return
}

func objs_Float3x2(xn *xmlx.Node, n string) (objs []*nga.Float3x2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float3x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x2(xn, "")
	}
	return
}

func obj_Bool3(xn *xmlx.Node, n string) (obj *nga.Bool3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Bool3)
	}
	return
}

func objs_Bool3(xn *xmlx.Node, n string) (objs []*nga.Bool3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Bool3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool3(xn, "")
	}
	return
}

func obj_Int2x2(xn *xmlx.Node, n string) (obj *nga.Int2x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int2x2)
	}
	return
}

func objs_Int2x2(xn *xmlx.Node, n string) (objs []*nga.Int2x2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int2x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int2x2(xn, "")
	}
	return
}

func obj_KxBind(xn *xmlx.Node, n string) (obj *nga.KxBind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxBind)
	}
	return
}

func objs_KxBind(xn *xmlx.Node, n string) (objs []*nga.KxBind) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxBind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxBind(xn, "")
	}
	return
}

func obj_FxAnnotation(xn *xmlx.Node, n string) (obj *nga.FxAnnotation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxAnnotation)
	}
	return
}

func objs_FxAnnotation(xn *xmlx.Node, n string) (objs []*nga.FxAnnotation) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxAnnotation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxAnnotation(xn, "")
	}
	return
}

func obj_FxCreateCubeInitFrom(xn *xmlx.Node, n string) (obj *nga.FxCreateCubeInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateCubeInitFrom)
	}
	return
}

func objs_FxCreateCubeInitFrom(xn *xmlx.Node, n string) (objs []*nga.FxCreateCubeInitFrom) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateCubeInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateCubeInitFrom(xn, "")
	}
	return
}

func obj_FxPassProgramBindAttribute(xn *xmlx.Node, n string) (obj *nga.FxPassProgramBindAttribute) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassProgramBindAttribute)
	}
	return
}

func objs_FxPassProgramBindAttribute(xn *xmlx.Node, n string) (objs []*nga.FxPassProgramBindAttribute) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassProgramBindAttribute, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramBindAttribute(xn, "")
	}
	return
}

func obj_FxMaterialInstBindVertexInput(xn *xmlx.Node, n string) (obj *nga.FxMaterialInstBindVertexInput) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxMaterialInstBindVertexInput)
	}
	return
}

func objs_FxMaterialInstBindVertexInput(xn *xmlx.Node, n string) (objs []*nga.FxMaterialInstBindVertexInput) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxMaterialInstBindVertexInput, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialInstBindVertexInput(xn, "")
	}
	return
}

func obj_FxPassProgramShader(xn *xmlx.Node, n string) (obj *nga.FxPassProgramShader) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassProgramShader)
	}
	return
}

func objs_FxPassProgramShader(xn *xmlx.Node, n string) (objs []*nga.FxPassProgramShader) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassProgramShader, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramShader(xn, "")
	}
	return
}

func obj_PxRigidBodyInst(xn *xmlx.Node, n string) (obj *nga.PxRigidBodyInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidBodyInst)
	}
	return
}

func objs_PxRigidBodyInst(xn *xmlx.Node, n string) (objs []*nga.PxRigidBodyInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidBodyInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyInst(xn, "")
	}
	return
}

func obj_GeometryBrepCone(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCone) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCone)
	}
	return
}

func objs_GeometryBrepCone(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCone) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCone, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCone(xn, "")
	}
	return
}

func obj_Float4x4(xn *xmlx.Node, n string) (obj *nga.Float4x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float4x4)
	}
	return
}

func objs_Float4x4(xn *xmlx.Node, n string) (objs []*nga.Float4x4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float4x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4x4(xn, "")
	}
	return
}

func obj_Int4x4(xn *xmlx.Node, n string) (obj *nga.Int4x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int4x4)
	}
	return
}

func objs_Int4x4(xn *xmlx.Node, n string) (objs []*nga.Int4x4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int4x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int4x4(xn, "")
	}
	return
}

func obj_CameraPerspective(xn *xmlx.Node, n string) (obj *nga.CameraPerspective) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraPerspective)
	}
	return
}

func objs_CameraPerspective(xn *xmlx.Node, n string) (objs []*nga.CameraPerspective) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraPerspective, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraPerspective(xn, "")
	}
	return
}

func obj_KxArticulatedSystemAxisIndex(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemAxisIndex) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemAxisIndex)
	}
	return
}

func objs_KxArticulatedSystemAxisIndex(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemAxisIndex) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemAxisIndex, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemAxisIndex(xn, "")
	}
	return
}

func obj_KxSceneInst(xn *xmlx.Node, n string) (obj *nga.KxSceneInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxSceneInst)
	}
	return
}

func objs_KxSceneInst(xn *xmlx.Node, n string) (objs []*nga.KxSceneInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxSceneInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneInst(xn, "")
	}
	return
}

func obj_LibNodeDefs(xn *xmlx.Node, n string) (obj *nga.LibNodeDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibNodeDefs)
	}
	return
}

func objs_LibNodeDefs(xn *xmlx.Node, n string) (objs []*nga.LibNodeDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibNodeDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibNodeDefs(xn, "")
	}
	return
}

func obj_ParamBool(xn *xmlx.Node, n string) (obj *nga.ParamBool) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamBool)
	}
	return
}

func objs_ParamBool(xn *xmlx.Node, n string) (objs []*nga.ParamBool) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamBool, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamBool(xn, "")
	}
	return
}

func obj_GeometryVertices(xn *xmlx.Node, n string) (obj *nga.GeometryVertices) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryVertices)
	}
	return
}

func objs_GeometryVertices(xn *xmlx.Node, n string) (objs []*nga.GeometryVertices) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryVertices, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryVertices(xn, "")
	}
	return
}

func obj_LibPxModelDefs(xn *xmlx.Node, n string) (obj *nga.LibPxModelDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibPxModelDefs)
	}
	return
}

func objs_LibPxModelDefs(xn *xmlx.Node, n string) (objs []*nga.LibPxModelDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibPxModelDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibPxModelDefs(xn, "")
	}
	return
}

func obj_FxCreateCube(xn *xmlx.Node, n string) (obj *nga.FxCreateCube) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateCube)
	}
	return
}

func objs_FxCreateCube(xn *xmlx.Node, n string) (objs []*nga.FxCreateCube) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateCube, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateCube(xn, "")
	}
	return
}

func obj_VisualSceneEvaluation(xn *xmlx.Node, n string) (obj *nga.VisualSceneEvaluation) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.VisualSceneEvaluation)
	}
	return
}

func objs_VisualSceneEvaluation(xn *xmlx.Node, n string) (objs []*nga.VisualSceneEvaluation) {
	xns := xsns(xn, n)
	objs = make([]*nga.VisualSceneEvaluation, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneEvaluation(xn, "")
	}
	return
}

func obj_KxArticulatedSystemKinematicsFrame(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemKinematicsFrame) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxArticulatedSystemKinematicsFrame)
	}
	return
}

func objs_KxArticulatedSystemKinematicsFrame(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemKinematicsFrame) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemKinematicsFrame, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemKinematicsFrame(xn, "")
	}
	return
}

func obj_PxRigidConstraintLimit(xn *xmlx.Node, n string) (obj *nga.PxRigidConstraintLimit) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidConstraintLimit)
	}
	return
}

func objs_PxRigidConstraintLimit(xn *xmlx.Node, n string) (objs []*nga.PxRigidConstraintLimit) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidConstraintLimit, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintLimit(xn, "")
	}
	return
}

func obj_MeshVertAtt2(xn *xmlx.Node, n string) (obj *nga.MeshVertAtt2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshVertAtt2)
	}
	return
}

func objs_MeshVertAtt2(xn *xmlx.Node, n string) (objs []*nga.MeshVertAtt2) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshVertAtt2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshVertAtt2(xn, "")
	}
	return
}

func obj_GeometryBrepWires(xn *xmlx.Node, n string) (obj *nga.GeometryBrepWires) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepWires)
	}
	return
}

func objs_GeometryBrepWires(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepWires) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepWires, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepWires(xn, "")
	}
	return
}

func obj_GeometryBrepSurfaceCurves(xn *xmlx.Node, n string) (obj *nga.GeometryBrepSurfaceCurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepSurfaceCurves)
	}
	return
}

func objs_GeometryBrepSurfaceCurves(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepSurfaceCurves) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepSurfaceCurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepSurfaceCurves(xn, "")
	}
	return
}

func obj_MeshFace3(xn *xmlx.Node, n string) (obj *nga.MeshFace3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshFace3)
	}
	return
}

func objs_MeshFace3(xn *xmlx.Node, n string) (objs []*nga.MeshFace3) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshFace3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshFace3(xn, "")
	}
	return
}

func obj_KxArticulatedSystemEffector(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemEffector) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewKxArticulatedSystemEffector()
	}
	return
}

func objs_KxArticulatedSystemEffector(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemEffector) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemEffector, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemEffector(xn, "")
	}
	return
}

func obj_LibPxSceneDefs(xn *xmlx.Node, n string) (obj *nga.LibPxSceneDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibPxSceneDefs)
	}
	return
}

func objs_LibPxSceneDefs(xn *xmlx.Node, n string) (objs []*nga.LibPxSceneDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibPxSceneDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibPxSceneDefs(xn, "")
	}
	return
}

func obj_GeometryBrepNurbsSurface(xn *xmlx.Node, n string) (obj *nga.GeometryBrepNurbsSurface) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewGeometryBrepNurbsSurface()
	}
	return
}

func objs_GeometryBrepNurbsSurface(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepNurbsSurface) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepNurbsSurface, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepNurbsSurface(xn, "")
	}
	return
}

func obj_GeometryControlVertices(xn *xmlx.Node, n string) (obj *nga.GeometryControlVertices) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryControlVertices)
	}
	return
}

func objs_GeometryControlVertices(xn *xmlx.Node, n string) (objs []*nga.GeometryControlVertices) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryControlVertices, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryControlVertices(xn, "")
	}
	return
}

func obj_InputShared(xn *xmlx.Node, n string) (obj *nga.InputShared) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.InputShared)
	}
	return
}

func objs_InputShared(xn *xmlx.Node, n string) (objs []*nga.InputShared) {
	xns := xsns(xn, n)
	objs = make([]*nga.InputShared, len(xns))
	for i, xn := range xns {
		objs[i] = obj_InputShared(xn, "")
	}
	return
}

func obj_LibFxEffectDefs(xn *xmlx.Node, n string) (obj *nga.LibFxEffectDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibFxEffectDefs)
	}
	return
}

func objs_LibFxEffectDefs(xn *xmlx.Node, n string) (objs []*nga.LibFxEffectDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibFxEffectDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibFxEffectDefs(xn, "")
	}
	return
}

func obj_Param(xn *xmlx.Node, n string) (obj *nga.Param) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Param)
	}
	return
}

func objs_Param(xn *xmlx.Node, n string) (objs []*nga.Param) {
	xns := xsns(xn, n)
	objs = make([]*nga.Param, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Param(xn, "")
	}
	return
}

func obj_MeshVertAtt3(xn *xmlx.Node, n string) (obj *nga.MeshVertAtt3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshVertAtt3)
	}
	return
}

func objs_MeshVertAtt3(xn *xmlx.Node, n string) (objs []*nga.MeshVertAtt3) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshVertAtt3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshVertAtt3(xn, "")
	}
	return
}

func obj_PxSceneDef(xn *xmlx.Node, n string) (obj *nga.PxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxSceneDef)
	}
	return
}

func objs_PxSceneDef(xn *xmlx.Node, n string) (objs []*nga.PxSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxSceneDef(xn, "")
	}
	return
}

func obj_AnimationClipDef(xn *xmlx.Node, n string) (obj *nga.AnimationClipDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AnimationClipDef)
	}
	return
}

func objs_AnimationClipDef(xn *xmlx.Node, n string) (objs []*nga.AnimationClipDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.AnimationClipDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationClipDef(xn, "")
	}
	return
}

func obj_Scene(xn *xmlx.Node, n string) (obj *nga.Scene) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Scene)
	}
	return
}

func objs_Scene(xn *xmlx.Node, n string) (objs []*nga.Scene) {
	xns := xsns(xn, n)
	objs = make([]*nga.Scene, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Scene(xn, "")
	}
	return
}

func obj_LibLightDefs(xn *xmlx.Node, n string) (obj *nga.LibLightDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibLightDefs)
	}
	return
}

func objs_LibLightDefs(xn *xmlx.Node, n string) (objs []*nga.LibLightDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibLightDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibLightDefs(xn, "")
	}
	return
}

func obj_Float4(xn *xmlx.Node, n string) (obj *nga.Float4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float4)
	}
	return
}

func objs_Float4(xn *xmlx.Node, n string) (objs []*nga.Float4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float4(xn, "")
	}
	return
}

func obj_Float2(xn *xmlx.Node, n string) (obj *nga.Float2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float2)
	}
	return
}

func objs_Float2(xn *xmlx.Node, n string) (objs []*nga.Float2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2(xn, "")
	}
	return
}

func obj_LibFxImageDefs(xn *xmlx.Node, n string) (obj *nga.LibFxImageDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibFxImageDefs)
	}
	return
}

func objs_LibFxImageDefs(xn *xmlx.Node, n string) (objs []*nga.LibFxImageDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibFxImageDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibFxImageDefs(xn, "")
	}
	return
}

func obj_ParamDefs(xn *xmlx.Node, n string) (obj *nga.ParamDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamDefs)
	}
	return
}

func objs_ParamDefs(xn *xmlx.Node, n string) (objs []*nga.ParamDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamDefs(xn, "")
	}
	return
}

func obj_LibsPxMaterialDef(xn *xmlx.Node, n string) (obj *nga.LibsPxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsPxMaterialDef)
	}
	return
}

func objs_LibsPxMaterialDef(xn *xmlx.Node, n string) (objs []*nga.LibsPxMaterialDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsPxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsPxMaterialDef(xn, "")
	}
	return
}

func obj_LibsNodeDef(xn *xmlx.Node, n string) (obj *nga.LibsNodeDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsNodeDef)
	}
	return
}

func objs_LibsNodeDef(xn *xmlx.Node, n string) (objs []*nga.LibsNodeDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsNodeDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsNodeDef(xn, "")
	}
	return
}

func obj_ControllerSkin(xn *xmlx.Node, n string) (obj *nga.ControllerSkin) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewControllerSkin()
	}
	return
}

func objs_ControllerSkin(xn *xmlx.Node, n string) (objs []*nga.ControllerSkin) {
	xns := xsns(xn, n)
	objs = make([]*nga.ControllerSkin, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerSkin(xn, "")
	}
	return
}

func obj_ListInts(xn *xmlx.Node, n string) (obj *nga.ListInts) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ListInts)
	}
	return
}

func objs_ListInts(xn *xmlx.Node, n string) (objs []*nga.ListInts) {
	xns := xsns(xn, n)
	objs = make([]*nga.ListInts, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ListInts(xn, "")
	}
	return
}

func obj_PxShape(xn *xmlx.Node, n string) (obj *nga.PxShape) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxShape)
	}
	return
}

func objs_PxShape(xn *xmlx.Node, n string) (objs []*nga.PxShape) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxShape, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxShape(xn, "")
	}
	return
}

func obj_Float2x4(xn *xmlx.Node, n string) (obj *nga.Float2x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float2x4)
	}
	return
}

func objs_Float2x4(xn *xmlx.Node, n string) (objs []*nga.Float2x4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float2x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x4(xn, "")
	}
	return
}

func obj_FxCreate3DInitFrom(xn *xmlx.Node, n string) (obj *nga.FxCreate3DInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreate3DInitFrom)
	}
	return
}

func objs_FxCreate3DInitFrom(xn *xmlx.Node, n string) (objs []*nga.FxCreate3DInitFrom) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreate3DInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate3DInitFrom(xn, "")
	}
	return
}

func obj_FxParamDef(xn *xmlx.Node, n string) (obj *nga.FxParamDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxParamDef)
	}
	return
}

func objs_FxParamDef(xn *xmlx.Node, n string) (objs []*nga.FxParamDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxParamDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxParamDef(xn, "")
	}
	return
}

func obj_LibKxJointDefs(xn *xmlx.Node, n string) (obj *nga.LibKxJointDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibKxJointDefs)
	}
	return
}

func objs_LibKxJointDefs(xn *xmlx.Node, n string) (objs []*nga.LibKxJointDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibKxJointDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibKxJointDefs(xn, "")
	}
	return
}

func obj_LibPxForceFieldDefs(xn *xmlx.Node, n string) (obj *nga.LibPxForceFieldDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibPxForceFieldDefs)
	}
	return
}

func objs_LibPxForceFieldDefs(xn *xmlx.Node, n string) (objs []*nga.LibPxForceFieldDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibPxForceFieldDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibPxForceFieldDefs(xn, "")
	}
	return
}

func obj_Float2x2(xn *xmlx.Node, n string) (obj *nga.Float2x2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float2x2)
	}
	return
}

func objs_Float2x2(xn *xmlx.Node, n string) (objs []*nga.Float2x2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float2x2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x2(xn, "")
	}
	return
}

func obj_CameraOrthographic(xn *xmlx.Node, n string) (obj *nga.CameraOrthographic) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraOrthographic)
	}
	return
}

func objs_CameraOrthographic(xn *xmlx.Node, n string) (objs []*nga.CameraOrthographic) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraOrthographic, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraOrthographic(xn, "")
	}
	return
}

func obj_Bool2(xn *xmlx.Node, n string) (obj *nga.Bool2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Bool2)
	}
	return
}

func objs_Bool2(xn *xmlx.Node, n string) (objs []*nga.Bool2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Bool2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool2(xn, "")
	}
	return
}

func obj_ListFloats(xn *xmlx.Node, n string) (obj *nga.ListFloats) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ListFloats)
	}
	return
}

func objs_ListFloats(xn *xmlx.Node, n string) (objs []*nga.ListFloats) {
	xns := xsns(xn, n)
	objs = make([]*nga.ListFloats, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ListFloats(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearColor(xn *xmlx.Node, n string) (obj *nga.FxPassEvaluationClearColor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassEvaluationClearColor)
	}
	return
}

func objs_FxPassEvaluationClearColor(xn *xmlx.Node, n string) (objs []*nga.FxPassEvaluationClearColor) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassEvaluationClearColor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearColor(xn, "")
	}
	return
}

func obj_PxModelDef(xn *xmlx.Node, n string) (obj *nga.PxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxModelDef)
	}
	return
}

func objs_PxModelDef(xn *xmlx.Node, n string) (objs []*nga.PxModelDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxModelDef(xn, "")
	}
	return
}

func obj_ParamUint(xn *xmlx.Node, n string) (obj *nga.ParamUint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamUint)
	}
	return
}

func objs_ParamUint(xn *xmlx.Node, n string) (objs []*nga.ParamUint) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamUint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamUint(xn, "")
	}
	return
}

func obj_LibsGeometryDef(xn *xmlx.Node, n string) (obj *nga.LibsGeometryDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsGeometryDef)
	}
	return
}

func objs_LibsGeometryDef(xn *xmlx.Node, n string) (objs []*nga.LibsGeometryDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsGeometryDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsGeometryDef(xn, "")
	}
	return
}

func obj_FormulaDef(xn *xmlx.Node, n string) (obj *nga.FormulaDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FormulaDef)
	}
	return
}

func objs_FormulaDef(xn *xmlx.Node, n string) (objs []*nga.FormulaDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.FormulaDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FormulaDef(xn, "")
	}
	return
}

func obj_FxSamplerFiltering(xn *xmlx.Node, n string) (obj *nga.FxSamplerFiltering) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxSamplerFiltering)
	}
	return
}

func objs_FxSamplerFiltering(xn *xmlx.Node, n string) (objs []*nga.FxSamplerFiltering) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxSamplerFiltering, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxSamplerFiltering(xn, "")
	}
	return
}

func obj_ScopedVec3(xn *xmlx.Node, n string) (obj *nga.ScopedVec3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ScopedVec3)
	}
	return
}

func objs_ScopedVec3(xn *xmlx.Node, n string) (objs []*nga.ScopedVec3) {
	xns := xsns(xn, n)
	objs = make([]*nga.ScopedVec3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ScopedVec3(xn, "")
	}
	return
}

func obj_FxCreate3D(xn *xmlx.Node, n string) (obj *nga.FxCreate3D) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreate3D)
	}
	return
}

func objs_FxCreate3D(xn *xmlx.Node, n string) (objs []*nga.FxCreate3D) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreate3D, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate3D(xn, "")
	}
	return
}

func obj_Float3(xn *xmlx.Node, n string) (obj *nga.Float3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float3)
	}
	return
}

func objs_Float3(xn *xmlx.Node, n string) (objs []*nga.Float3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearDepth(xn *xmlx.Node, n string) (obj *nga.FxPassEvaluationClearDepth) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassEvaluationClearDepth)
	}
	return
}

func objs_FxPassEvaluationClearDepth(xn *xmlx.Node, n string) (objs []*nga.FxPassEvaluationClearDepth) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassEvaluationClearDepth, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearDepth(xn, "")
	}
	return
}

func obj_Float3x4(xn *xmlx.Node, n string) (obj *nga.Float3x4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float3x4)
	}
	return
}

func objs_Float3x4(xn *xmlx.Node, n string) (objs []*nga.Float3x4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float3x4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float3x4(xn, "")
	}
	return
}

func obj_LibsCameraDef(xn *xmlx.Node, n string) (obj *nga.LibsCameraDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsCameraDef)
	}
	return
}

func objs_LibsCameraDef(xn *xmlx.Node, n string) (objs []*nga.LibsCameraDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsCameraDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsCameraDef(xn, "")
	}
	return
}

func obj_LightDirectional(xn *xmlx.Node, n string) (obj *nga.LightDirectional) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightDirectional)
	}
	return
}

func objs_LightDirectional(xn *xmlx.Node, n string) (objs []*nga.LightDirectional) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightDirectional, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightDirectional(xn, "")
	}
	return
}

func obj_FxPassProgramBindUniform(xn *xmlx.Node, n string) (obj *nga.FxPassProgramBindUniform) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassProgramBindUniform)
	}
	return
}

func objs_FxPassProgramBindUniform(xn *xmlx.Node, n string) (objs []*nga.FxPassProgramBindUniform) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassProgramBindUniform, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramBindUniform(xn, "")
	}
	return
}

func obj_GeometryBrepHyperbola(xn *xmlx.Node, n string) (obj *nga.GeometryBrepHyperbola) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepHyperbola)
	}
	return
}

func objs_GeometryBrepHyperbola(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepHyperbola) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepHyperbola, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepHyperbola(xn, "")
	}
	return
}

func obj_PxRigidConstraintDef(xn *xmlx.Node, n string) (obj *nga.PxRigidConstraintDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidConstraintDef)
	}
	return
}

func objs_PxRigidConstraintDef(xn *xmlx.Node, n string) (objs []*nga.PxRigidConstraintDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidConstraintDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintDef(xn, "")
	}
	return
}

func obj_FxTechnique(xn *xmlx.Node, n string) (obj *nga.FxTechnique) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechnique)
	}
	return
}

func objs_FxTechnique(xn *xmlx.Node, n string) (objs []*nga.FxTechnique) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechnique, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechnique(xn, "")
	}
	return
}

func obj_PxRigidConstraintAttachment(xn *xmlx.Node, n string) (obj *nga.PxRigidConstraintAttachment) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidConstraintAttachment)
	}
	return
}

func objs_PxRigidConstraintAttachment(xn *xmlx.Node, n string) (objs []*nga.PxRigidConstraintAttachment) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidConstraintAttachment, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidConstraintAttachment(xn, "")
	}
	return
}

func obj_GeometryBrepCurves(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCurves) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCurves)
	}
	return
}

func objs_GeometryBrepCurves(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCurves) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCurves, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCurves(xn, "")
	}
	return
}

func obj_Int2(xn *xmlx.Node, n string) (obj *nga.Int2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int2)
	}
	return
}

func objs_Int2(xn *xmlx.Node, n string) (objs []*nga.Int2) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int2(xn, "")
	}
	return
}

func obj_FxTechniqueCommonLambert(xn *xmlx.Node, n string) (obj *nga.FxTechniqueCommonLambert) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueCommonLambert)
	}
	return
}

func objs_FxTechniqueCommonLambert(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueCommonLambert) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueCommonLambert, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommonLambert(xn, "")
	}
	return
}

func obj_KxJointLimits(xn *xmlx.Node, n string) (obj *nga.KxJointLimits) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxJointLimits)
	}
	return
}

func objs_KxJointLimits(xn *xmlx.Node, n string) (objs []*nga.KxJointLimits) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxJointLimits, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointLimits(xn, "")
	}
	return
}

func obj_FxMaterialDef(xn *xmlx.Node, n string) (obj *nga.FxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxMaterialDef)
	}
	return
}

func objs_FxMaterialDef(xn *xmlx.Node, n string) (objs []*nga.FxMaterialDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialDef(xn, "")
	}
	return
}

func obj_CameraImager(xn *xmlx.Node, n string) (obj *nga.CameraImager) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraImager)
	}
	return
}

func objs_CameraImager(xn *xmlx.Node, n string) (objs []*nga.CameraImager) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraImager, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraImager(xn, "")
	}
	return
}

func obj_KxJointDef(xn *xmlx.Node, n string) (obj *nga.KxJointDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxJointDef)
	}
	return
}

func objs_KxJointDef(xn *xmlx.Node, n string) (objs []*nga.KxJointDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxJointDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJointDef(xn, "")
	}
	return
}

func obj_ScopedFloat3(xn *xmlx.Node, n string) (obj *nga.ScopedFloat3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ScopedFloat3)
	}
	return
}

func objs_ScopedFloat3(xn *xmlx.Node, n string) (objs []*nga.ScopedFloat3) {
	xns := xsns(xn, n)
	objs = make([]*nga.ScopedFloat3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ScopedFloat3(xn, "")
	}
	return
}

func obj_LibsKxModelDef(xn *xmlx.Node, n string) (obj *nga.LibsKxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsKxModelDef)
	}
	return
}

func objs_LibsKxModelDef(xn *xmlx.Node, n string) (objs []*nga.LibsKxModelDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsKxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsKxModelDef(xn, "")
	}
	return
}

func obj_PxMaterialDef(xn *xmlx.Node, n string) (obj *nga.PxMaterialDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxMaterialDef)
	}
	return
}

func objs_PxMaterialDef(xn *xmlx.Node, n string) (objs []*nga.PxMaterialDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxMaterialDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxMaterialDef(xn, "")
	}
	return
}

func obj_Input(xn *xmlx.Node, n string) (obj *nga.Input) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Input)
	}
	return
}

func objs_Input(xn *xmlx.Node, n string) (objs []*nga.Input) {
	xns := xsns(xn, n)
	objs = make([]*nga.Input, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Input(xn, "")
	}
	return
}

func obj_FxMaterialInstBind(xn *xmlx.Node, n string) (obj *nga.FxMaterialInstBind) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxMaterialInstBind)
	}
	return
}

func objs_FxMaterialInstBind(xn *xmlx.Node, n string) (objs []*nga.FxMaterialInstBind) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxMaterialInstBind, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxMaterialInstBind(xn, "")
	}
	return
}

func obj_FormulaInst(xn *xmlx.Node, n string) (obj *nga.FormulaInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FormulaInst)
	}
	return
}

func objs_FormulaInst(xn *xmlx.Node, n string) (objs []*nga.FormulaInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.FormulaInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FormulaInst(xn, "")
	}
	return
}

func obj_FxCreateFormatHint(xn *xmlx.Node, n string) (obj *nga.FxCreateFormatHint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateFormatHint)
	}
	return
}

func objs_FxCreateFormatHint(xn *xmlx.Node, n string) (objs []*nga.FxCreateFormatHint) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateFormatHint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateFormatHint(xn, "")
	}
	return
}

func obj_VisualSceneRenderingMaterialInst(xn *xmlx.Node, n string) (obj *nga.VisualSceneRenderingMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.VisualSceneRenderingMaterialInst)
	}
	return
}

func objs_VisualSceneRenderingMaterialInst(xn *xmlx.Node, n string) (objs []*nga.VisualSceneRenderingMaterialInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.VisualSceneRenderingMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_VisualSceneRenderingMaterialInst(xn, "")
	}
	return
}

func obj_FxTechniqueCommonBlinn(xn *xmlx.Node, n string) (obj *nga.FxTechniqueCommonBlinn) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueCommonBlinn)
	}
	return
}

func objs_FxTechniqueCommonBlinn(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueCommonBlinn) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueCommonBlinn, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueCommonBlinn(xn, "")
	}
	return
}

func obj_LibAnimationDefs(xn *xmlx.Node, n string) (obj *nga.LibAnimationDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibAnimationDefs)
	}
	return
}

func objs_LibAnimationDefs(xn *xmlx.Node, n string) (objs []*nga.LibAnimationDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibAnimationDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibAnimationDefs(xn, "")
	}
	return
}

func obj_FxProfileCommon(xn *xmlx.Node, n string) (obj *nga.FxProfileCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewFxProfileCommon()
	}
	return
}

func objs_FxProfileCommon(xn *xmlx.Node, n string) (objs []*nga.FxProfileCommon) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxProfileCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfileCommon(xn, "")
	}
	return
}

func obj_Float2x3(xn *xmlx.Node, n string) (obj *nga.Float2x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float2x3)
	}
	return
}

func objs_Float2x3(xn *xmlx.Node, n string) (objs []*nga.Float2x3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float2x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float2x3(xn, "")
	}
	return
}

func obj_LibsPxForceFieldDef(xn *xmlx.Node, n string) (obj *nga.LibsPxForceFieldDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsPxForceFieldDef)
	}
	return
}

func objs_LibsPxForceFieldDef(xn *xmlx.Node, n string) (objs []*nga.LibsPxForceFieldDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsPxForceFieldDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsPxForceFieldDef(xn, "")
	}
	return
}

func obj_LibsPxSceneDef(xn *xmlx.Node, n string) (obj *nga.LibsPxSceneDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsPxSceneDef)
	}
	return
}

func objs_LibsPxSceneDef(xn *xmlx.Node, n string) (objs []*nga.LibsPxSceneDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsPxSceneDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsPxSceneDef(xn, "")
	}
	return
}

func obj_Int3x3(xn *xmlx.Node, n string) (obj *nga.Int3x3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int3x3)
	}
	return
}

func objs_Int3x3(xn *xmlx.Node, n string) (objs []*nga.Int3x3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int3x3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int3x3(xn, "")
	}
	return
}

func obj_AnimationChannel(xn *xmlx.Node, n string) (obj *nga.AnimationChannel) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AnimationChannel)
	}
	return
}

func objs_AnimationChannel(xn *xmlx.Node, n string) (objs []*nga.AnimationChannel) {
	xns := xsns(xn, n)
	objs = make([]*nga.AnimationChannel, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationChannel(xn, "")
	}
	return
}

func obj_FxEffectInst(xn *xmlx.Node, n string) (obj *nga.FxEffectInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxEffectInst)
	}
	return
}

func objs_FxEffectInst(xn *xmlx.Node, n string) (objs []*nga.FxEffectInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxEffectInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectInst(xn, "")
	}
	return
}

func obj_LightDef(xn *xmlx.Node, n string) (obj *nga.LightDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightDef)
	}
	return
}

func objs_LightDef(xn *xmlx.Node, n string) (objs []*nga.LightDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightDef(xn, "")
	}
	return
}

func obj_Source(xn *xmlx.Node, n string) (obj *nga.Source) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Source)
	}
	return
}

func objs_Source(xn *xmlx.Node, n string) (objs []*nga.Source) {
	xns := xsns(xn, n)
	objs = make([]*nga.Source, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Source(xn, "")
	}
	return
}

func obj_LibAnimationClipDefs(xn *xmlx.Node, n string) (obj *nga.LibAnimationClipDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibAnimationClipDefs)
	}
	return
}

func objs_LibAnimationClipDefs(xn *xmlx.Node, n string) (objs []*nga.LibAnimationClipDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibAnimationClipDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibAnimationClipDefs(xn, "")
	}
	return
}

func obj_FxPass(xn *xmlx.Node, n string) (obj *nga.FxPass) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPass)
	}
	return
}

func objs_FxPass(xn *xmlx.Node, n string) (objs []*nga.FxPass) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPass, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPass(xn, "")
	}
	return
}

func obj_MeshProvider(xn *xmlx.Node, n string) (obj *nga.MeshProvider) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.MeshProvider)
	}
	return
}

func objs_MeshProvider(xn *xmlx.Node, n string) (objs []*nga.MeshProvider) {
	xns := xsns(xn, n)
	objs = make([]*nga.MeshProvider, len(xns))
	for i, xn := range xns {
		objs[i] = obj_MeshProvider(xn, "")
	}
	return
}

func obj_AnimationSampler(xn *xmlx.Node, n string) (obj *nga.AnimationSampler) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AnimationSampler)
	}
	return
}

func objs_AnimationSampler(xn *xmlx.Node, n string) (objs []*nga.AnimationSampler) {
	xns := xsns(xn, n)
	objs = make([]*nga.AnimationSampler, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AnimationSampler(xn, "")
	}
	return
}

func obj_Extra(xn *xmlx.Node, n string) (obj *nga.Extra) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Extra)
	}
	return
}

func objs_Extra(xn *xmlx.Node, n string) (objs []*nga.Extra) {
	xns := xsns(xn, n)
	objs = make([]*nga.Extra, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Extra(xn, "")
	}
	return
}

func obj_LightPoint(xn *xmlx.Node, n string) (obj *nga.LightPoint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightPoint)
	}
	return
}

func objs_LightPoint(xn *xmlx.Node, n string) (objs []*nga.LightPoint) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightPoint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightPoint(xn, "")
	}
	return
}

func obj_FxInitFrom(xn *xmlx.Node, n string) (obj *nga.FxInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxInitFrom)
	}
	return
}

func objs_FxInitFrom(xn *xmlx.Node, n string) (objs []*nga.FxInitFrom) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxInitFrom(xn, "")
	}
	return
}

func obj_AssetContributor(xn *xmlx.Node, n string) (obj *nga.AssetContributor) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.AssetContributor)
	}
	return
}

func objs_AssetContributor(xn *xmlx.Node, n string) (objs []*nga.AssetContributor) {
	xns := xsns(xn, n)
	objs = make([]*nga.AssetContributor, len(xns))
	for i, xn := range xns {
		objs[i] = obj_AssetContributor(xn, "")
	}
	return
}

func obj_PxRigidBodyDef(xn *xmlx.Node, n string) (obj *nga.PxRigidBodyDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxRigidBodyDef)
	}
	return
}

func objs_PxRigidBodyDef(xn *xmlx.Node, n string) (objs []*nga.PxRigidBodyDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxRigidBodyDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxRigidBodyDef(xn, "")
	}
	return
}

func obj_Asset(xn *xmlx.Node, n string) (obj *nga.Asset) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Asset)
	}
	return
}

func objs_Asset(xn *xmlx.Node, n string) (objs []*nga.Asset) {
	xns := xsns(xn, n)
	objs = make([]*nga.Asset, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Asset(xn, "")
	}
	return
}

func obj_FxEffectDef(xn *xmlx.Node, n string) (obj *nga.FxEffectDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxEffectDef)
	}
	return
}

func objs_FxEffectDef(xn *xmlx.Node, n string) (objs []*nga.FxEffectDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxEffectDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectDef(xn, "")
	}
	return
}

func obj_GeometryBrepEdges(xn *xmlx.Node, n string) (obj *nga.GeometryBrepEdges) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepEdges)
	}
	return
}

func objs_GeometryBrepEdges(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepEdges) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepEdges, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepEdges(xn, "")
	}
	return
}

func obj_KxJoint(xn *xmlx.Node, n string) (obj *nga.KxJoint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxJoint)
	}
	return
}

func objs_KxJoint(xn *xmlx.Node, n string) (objs []*nga.KxJoint) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxJoint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxJoint(xn, "")
	}
	return
}

func obj_ControllerDef(xn *xmlx.Node, n string) (obj *nga.ControllerDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ControllerDef)
	}
	return
}

func objs_ControllerDef(xn *xmlx.Node, n string) (objs []*nga.ControllerDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.ControllerDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ControllerDef(xn, "")
	}
	return
}

func obj_FxCreateCommon(xn *xmlx.Node, n string) (obj *nga.FxCreateCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreateCommon)
	}
	return
}

func objs_FxCreateCommon(xn *xmlx.Node, n string) (objs []*nga.FxCreateCommon) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreateCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreateCommon(xn, "")
	}
	return
}

func obj_FxParamDefs(xn *xmlx.Node, n string) (obj *nga.FxParamDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxParamDefs)
	}
	return
}

func objs_FxParamDefs(xn *xmlx.Node, n string) (objs []*nga.FxParamDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxParamDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxParamDefs(xn, "")
	}
	return
}

func obj_CameraCommon(xn *xmlx.Node, n string) (obj *nga.CameraCommon) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraCommon)
	}
	return
}

func objs_CameraCommon(xn *xmlx.Node, n string) (objs []*nga.CameraCommon) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraCommon, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraCommon(xn, "")
	}
	return
}

func obj_PxForceFieldInst(xn *xmlx.Node, n string) (obj *nga.PxForceFieldInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxForceFieldInst)
	}
	return
}

func objs_PxForceFieldInst(xn *xmlx.Node, n string) (objs []*nga.PxForceFieldInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxForceFieldInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxForceFieldInst(xn, "")
	}
	return
}

func obj_FxCreate2DSizeRatio(xn *xmlx.Node, n string) (obj *nga.FxCreate2DSizeRatio) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreate2DSizeRatio)
	}
	return
}

func objs_FxCreate2DSizeRatio(xn *xmlx.Node, n string) (objs []*nga.FxCreate2DSizeRatio) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreate2DSizeRatio, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2DSizeRatio(xn, "")
	}
	return
}

func obj_FxEffectInstTechniqueHint(xn *xmlx.Node, n string) (obj *nga.FxEffectInstTechniqueHint) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxEffectInstTechniqueHint)
	}
	return
}

func objs_FxEffectInstTechniqueHint(xn *xmlx.Node, n string) (objs []*nga.FxEffectInstTechniqueHint) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxEffectInstTechniqueHint, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxEffectInstTechniqueHint(xn, "")
	}
	return
}

func obj_PxMaterialInst(xn *xmlx.Node, n string) (obj *nga.PxMaterialInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.PxMaterialInst)
	}
	return
}

func objs_PxMaterialInst(xn *xmlx.Node, n string) (objs []*nga.PxMaterialInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.PxMaterialInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_PxMaterialInst(xn, "")
	}
	return
}

func obj_GeometryMesh(xn *xmlx.Node, n string) (obj *nga.GeometryMesh) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewGeometryMesh()
	}
	return
}

func objs_GeometryMesh(xn *xmlx.Node, n string) (objs []*nga.GeometryMesh) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryMesh, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryMesh(xn, "")
	}
	return
}

func obj_GeometryBrepCapsule(xn *xmlx.Node, n string) (obj *nga.GeometryBrepCapsule) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepCapsule)
	}
	return
}

func objs_GeometryBrepCapsule(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepCapsule) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepCapsule, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepCapsule(xn, "")
	}
	return
}

func obj_FxPassProgramShaderSources(xn *xmlx.Node, n string) (obj *nga.FxPassProgramShaderSources) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassProgramShaderSources)
	}
	return
}

func objs_FxPassProgramShaderSources(xn *xmlx.Node, n string) (objs []*nga.FxPassProgramShaderSources) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassProgramShaderSources, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassProgramShaderSources(xn, "")
	}
	return
}

func obj_BindMaterial(xn *xmlx.Node, n string) (obj *nga.BindMaterial) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.BindMaterial)
	}
	return
}

func objs_BindMaterial(xn *xmlx.Node, n string) (objs []*nga.BindMaterial) {
	xns := xsns(xn, n)
	objs = make([]*nga.BindMaterial, len(xns))
	for i, xn := range xns {
		objs[i] = obj_BindMaterial(xn, "")
	}
	return
}

func obj_Bool4(xn *xmlx.Node, n string) (obj *nga.Bool4) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Bool4)
	}
	return
}

func objs_Bool4(xn *xmlx.Node, n string) (objs []*nga.Bool4) {
	xns := xsns(xn, n)
	objs = make([]*nga.Bool4, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Bool4(xn, "")
	}
	return
}

func obj_Int3(xn *xmlx.Node, n string) (obj *nga.Int3) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Int3)
	}
	return
}

func objs_Int3(xn *xmlx.Node, n string) (objs []*nga.Int3) {
	xns := xsns(xn, n)
	objs = make([]*nga.Int3, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Int3(xn, "")
	}
	return
}

func obj_ParamFloat2(xn *xmlx.Node, n string) (obj *nga.ParamFloat2) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamFloat2)
	}
	return
}

func objs_ParamFloat2(xn *xmlx.Node, n string) (objs []*nga.ParamFloat2) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamFloat2, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamFloat2(xn, "")
	}
	return
}

func obj_CameraInst(xn *xmlx.Node, n string) (obj *nga.CameraInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraInst)
	}
	return
}

func objs_CameraInst(xn *xmlx.Node, n string) (objs []*nga.CameraInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraInst(xn, "")
	}
	return
}

func obj_LibCameraDefs(xn *xmlx.Node, n string) (obj *nga.LibCameraDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibCameraDefs)
	}
	return
}

func objs_LibCameraDefs(xn *xmlx.Node, n string) (objs []*nga.LibCameraDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibCameraDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibCameraDefs(xn, "")
	}
	return
}

func obj_CameraDef(xn *xmlx.Node, n string) (obj *nga.CameraDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.CameraDef)
	}
	return
}

func objs_CameraDef(xn *xmlx.Node, n string) (objs []*nga.CameraDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.CameraDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_CameraDef(xn, "")
	}
	return
}

func obj_FxCreate2DSizeExact(xn *xmlx.Node, n string) (obj *nga.FxCreate2DSizeExact) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxCreate2DSizeExact)
	}
	return
}

func objs_FxCreate2DSizeExact(xn *xmlx.Node, n string) (objs []*nga.FxCreate2DSizeExact) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxCreate2DSizeExact, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxCreate2DSizeExact(xn, "")
	}
	return
}

func obj_LightSpot(xn *xmlx.Node, n string) (obj *nga.LightSpot) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LightSpot)
	}
	return
}

func objs_LightSpot(xn *xmlx.Node, n string) (objs []*nga.LightSpot) {
	xns := xsns(xn, n)
	objs = make([]*nga.LightSpot, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LightSpot(xn, "")
	}
	return
}

func obj_LibsFxImageDef(xn *xmlx.Node, n string) (obj *nga.LibsFxImageDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsFxImageDef)
	}
	return
}

func objs_LibsFxImageDef(xn *xmlx.Node, n string) (objs []*nga.LibsFxImageDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsFxImageDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsFxImageDef(xn, "")
	}
	return
}

func obj_FxPassEvaluationClearStencil(xn *xmlx.Node, n string) (obj *nga.FxPassEvaluationClearStencil) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassEvaluationClearStencil)
	}
	return
}

func objs_FxPassEvaluationClearStencil(xn *xmlx.Node, n string) (objs []*nga.FxPassEvaluationClearStencil) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassEvaluationClearStencil, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassEvaluationClearStencil(xn, "")
	}
	return
}

func obj_GeometryBrepParabola(xn *xmlx.Node, n string) (obj *nga.GeometryBrepParabola) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepParabola)
	}
	return
}

func objs_GeometryBrepParabola(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepParabola) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepParabola, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepParabola(xn, "")
	}
	return
}

func obj_FxImageInitFrom(xn *xmlx.Node, n string) (obj *nga.FxImageInitFrom) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxImageInitFrom)
	}
	return
}

func objs_FxImageInitFrom(xn *xmlx.Node, n string) (objs []*nga.FxImageInitFrom) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxImageInitFrom, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxImageInitFrom(xn, "")
	}
	return
}

func obj_FxProfile(xn *xmlx.Node, n string) (obj *nga.FxProfile) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxProfile)
	}
	return
}

func objs_FxProfile(xn *xmlx.Node, n string) (objs []*nga.FxProfile) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxProfile, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxProfile(xn, "")
	}
	return
}

func obj_SceneGraph(xn *xmlx.Node, n string) (obj *nga.SceneGraph) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.SceneGraph)
	}
	return
}

func objs_SceneGraph(xn *xmlx.Node, n string) (objs []*nga.SceneGraph) {
	xns := xsns(xn, n)
	objs = make([]*nga.SceneGraph, len(xns))
	for i, xn := range xns {
		objs[i] = obj_SceneGraph(xn, "")
	}
	return
}

func obj_LibKxSceneDefs(xn *xmlx.Node, n string) (obj *nga.LibKxSceneDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibKxSceneDefs)
	}
	return
}

func objs_LibKxSceneDefs(xn *xmlx.Node, n string) (objs []*nga.LibKxSceneDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibKxSceneDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibKxSceneDefs(xn, "")
	}
	return
}

func obj_GeometryBrepEllipse(xn *xmlx.Node, n string) (obj *nga.GeometryBrepEllipse) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepEllipse)
	}
	return
}

func objs_GeometryBrepEllipse(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepEllipse) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepEllipse, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepEllipse(xn, "")
	}
	return
}

func obj_LibsKxJointDef(xn *xmlx.Node, n string) (obj *nga.LibsKxJointDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsKxJointDef)
	}
	return
}

func objs_LibsKxJointDef(xn *xmlx.Node, n string) (objs []*nga.LibsKxJointDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsKxJointDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsKxJointDef(xn, "")
	}
	return
}

func obj_KxArticulatedSystemKinematicsAxis(xn *xmlx.Node, n string) (obj *nga.KxArticulatedSystemKinematicsAxis) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = nga.NewKxArticulatedSystemKinematicsAxis()
	}
	return
}

func objs_KxArticulatedSystemKinematicsAxis(xn *xmlx.Node, n string) (objs []*nga.KxArticulatedSystemKinematicsAxis) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxArticulatedSystemKinematicsAxis, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxArticulatedSystemKinematicsAxis(xn, "")
	}
	return
}

func obj_LibGeometryDefs(xn *xmlx.Node, n string) (obj *nga.LibGeometryDefs) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibGeometryDefs)
	}
	return
}

func objs_LibGeometryDefs(xn *xmlx.Node, n string) (objs []*nga.LibGeometryDefs) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibGeometryDefs, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibGeometryDefs(xn, "")
	}
	return
}

func obj_ParamDef(xn *xmlx.Node, n string) (obj *nga.ParamDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamDef)
	}
	return
}

func objs_ParamDef(xn *xmlx.Node, n string) (objs []*nga.ParamDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamDef(xn, "")
	}
	return
}

func obj_FxPassState(xn *xmlx.Node, n string) (obj *nga.FxPassState) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxPassState)
	}
	return
}

func objs_FxPassState(xn *xmlx.Node, n string) (objs []*nga.FxPassState) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxPassState, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxPassState(xn, "")
	}
	return
}

func obj_ScopedFloat(xn *xmlx.Node, n string) (obj *nga.ScopedFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ScopedFloat)
	}
	return
}

func objs_ScopedFloat(xn *xmlx.Node, n string) (objs []*nga.ScopedFloat) {
	xns := xsns(xn, n)
	objs = make([]*nga.ScopedFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ScopedFloat(xn, "")
	}
	return
}

func obj_FxTechniqueGlsl(xn *xmlx.Node, n string) (obj *nga.FxTechniqueGlsl) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.FxTechniqueGlsl)
	}
	return
}

func objs_FxTechniqueGlsl(xn *xmlx.Node, n string) (objs []*nga.FxTechniqueGlsl) {
	xns := xsns(xn, n)
	objs = make([]*nga.FxTechniqueGlsl, len(xns))
	for i, xn := range xns {
		objs[i] = obj_FxTechniqueGlsl(xn, "")
	}
	return
}

func obj_LibsPxModelDef(xn *xmlx.Node, n string) (obj *nga.LibsPxModelDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsPxModelDef)
	}
	return
}

func objs_LibsPxModelDef(xn *xmlx.Node, n string) (objs []*nga.LibsPxModelDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsPxModelDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsPxModelDef(xn, "")
	}
	return
}

func obj_ParamFloat(xn *xmlx.Node, n string) (obj *nga.ParamFloat) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.ParamFloat)
	}
	return
}

func objs_ParamFloat(xn *xmlx.Node, n string) (objs []*nga.ParamFloat) {
	xns := xsns(xn, n)
	objs = make([]*nga.ParamFloat, len(xns))
	for i, xn := range xns {
		objs[i] = obj_ParamFloat(xn, "")
	}
	return
}

func obj_KxSceneInstBindModel(xn *xmlx.Node, n string) (obj *nga.KxSceneInstBindModel) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.KxSceneInstBindModel)
	}
	return
}

func objs_KxSceneInstBindModel(xn *xmlx.Node, n string) (objs []*nga.KxSceneInstBindModel) {
	xns := xsns(xn, n)
	objs = make([]*nga.KxSceneInstBindModel, len(xns))
	for i, xn := range xns {
		objs[i] = obj_KxSceneInstBindModel(xn, "")
	}
	return
}

func obj_LibsFormulaDef(xn *xmlx.Node, n string) (obj *nga.LibsFormulaDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsFormulaDef)
	}
	return
}

func objs_LibsFormulaDef(xn *xmlx.Node, n string) (objs []*nga.LibsFormulaDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsFormulaDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsFormulaDef(xn, "")
	}
	return
}

func obj_LibsAnimationDef(xn *xmlx.Node, n string) (obj *nga.LibsAnimationDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsAnimationDef)
	}
	return
}

func objs_LibsAnimationDef(xn *xmlx.Node, n string) (objs []*nga.LibsAnimationDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsAnimationDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsAnimationDef(xn, "")
	}
	return
}

func obj_GeometryBrepBox(xn *xmlx.Node, n string) (obj *nga.GeometryBrepBox) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryBrepBox)
	}
	return
}

func objs_GeometryBrepBox(xn *xmlx.Node, n string) (objs []*nga.GeometryBrepBox) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryBrepBox, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryBrepBox(xn, "")
	}
	return
}

func obj_LibsFxEffectDef(xn *xmlx.Node, n string) (obj *nga.LibsFxEffectDef) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.LibsFxEffectDef)
	}
	return
}

func objs_LibsFxEffectDef(xn *xmlx.Node, n string) (objs []*nga.LibsFxEffectDef) {
	xns := xsns(xn, n)
	objs = make([]*nga.LibsFxEffectDef, len(xns))
	for i, xn := range xns {
		objs[i] = obj_LibsFxEffectDef(xn, "")
	}
	return
}

func obj_Layers(xn *xmlx.Node, n string) (obj *nga.Layers) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Layers)
	}
	return
}

func objs_Layers(xn *xmlx.Node, n string) (objs []*nga.Layers) {
	xns := xsns(xn, n)
	objs = make([]*nga.Layers, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Layers(xn, "")
	}
	return
}

func obj_Float7(xn *xmlx.Node, n string) (obj *nga.Float7) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.Float7)
	}
	return
}

func objs_Float7(xn *xmlx.Node, n string) (objs []*nga.Float7) {
	xns := xsns(xn, n)
	objs = make([]*nga.Float7, len(xns))
	for i, xn := range xns {
		objs[i] = obj_Float7(xn, "")
	}
	return
}

func obj_GeometryPrimitives(xn *xmlx.Node, n string) (obj *nga.GeometryPrimitives) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.GeometryPrimitives)
	}
	return
}

func objs_GeometryPrimitives(xn *xmlx.Node, n string) (objs []*nga.GeometryPrimitives) {
	xns := xsns(xn, n)
	objs = make([]*nga.GeometryPrimitives, len(xns))
	for i, xn := range xns {
		objs[i] = obj_GeometryPrimitives(xn, "")
	}
	return
}

func obj_NodeInst(xn *xmlx.Node, n string) (obj *nga.NodeInst) {
	if (xn != nil) && (len(n) > 0) {
		xn = xsn(xn, n)
	}
	if xn != nil {
		obj = new(nga.NodeInst)
	}
	return
}

func objs_NodeInst(xn *xmlx.Node, n string) (objs []*nga.NodeInst) {
	xns := xsns(xn, n)
	objs = make([]*nga.NodeInst, len(xns))
	for i, xn := range xns {
		objs[i] = obj_NodeInst(xn, "")
	}
	return
}
