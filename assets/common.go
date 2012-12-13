package assets

import (
	xmlx "github.com/jteeuwen/go-pkg-xmlx"
)

const (
	//	A position and orientation transformation suitable for aiming a camera.
	TRANSFORM_TYPE_LOOKAT = 1
	//	A transformation that embodies mathematical changes to points within a coordinate system or the coordinate system itself.
	TRANSFORM_TYPE_MATRIX = iota
	//	A transformation that specifies how to rotate an object around an axis.
	TRANSFORM_TYPE_ROTATE = iota
	//	A transformation that specifies how to deform an object along one axis.
	TRANSFORM_TYPE_SKEW = iota
	//	A transformation that specifies how to change an object's size.
	TRANSFORM_TYPE_SCALE = iota
	//	A transformation that changes the position of an object in a local coordinate system.
	TRANSFORM_TYPE_TRANSLATE = iota
)

//	Used in all resources that require asset-management information.
type HasAsset struct {
	//	Resource-specific asset-management information and meta-data.
	Asset *Asset
}

//	Used in all resources that support custom techniques / foreign profiles.
type HasExtras struct {
	//	Custom-technique/foreign-profile meta-data.
	Extras []*Extra
}

//	Used in all FX resources that declare their own parameters.
type HasFxParamDefs struct {
	//	A hash-table containing all parameter declarations of this resource.
	NewParams FxParamDefs
}

//	Used in all resources that declare their own unique identifier.
type HasId struct {
	//	The unique identifier of this resource.
	Id string
}

//	Used in all data consumers that require input connections into a data Source.
type HasInputs struct {
	//	Declares the input semantics of a data Source and connects a consumer to that Source.
	Inputs []*Input
}

//	Used in all resources that support arbitrary pretty-print names/titles.
type HasName struct {
	//	The optional pretty-print name/title of this *Def*, *Inst* or *Lib*.
	Name string
}

//	Used in all resources that declare their own parameters.
type HasParamDefs struct {
	//	A hash-table containing all parameter declarations of this resource.
	NewParams ParamDefs
}

//	Used in all resources that assign values to other parameters.
type HasParamInsts struct {
	//	All parameter values assigned by this resource.
	SetParams []*ParamInst
}

//	Used in all resources that declare their own scoped identifier.
type HasSid struct {
	//	The Scoped identifier of this resource.
	Sid string
}

//	Used in all resources that provide data arrays.
type HasSources struct {
	//	Provides the bulk of this resource's data.
	Sources Sources
}

//	Used in all resources that support custom techniques / foreign profiles.
type HasTechniques struct {
	//	Custom-technique/foreign-profile content or data.
	Techniques []*Technique
}

//	Resource-specific asset-management information and meta-data.
type Asset struct {
	//	Custom-technique/foreign-profile meta-data.
	HasExtras
	//	Contains the date and time that the parent element was created.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Created string
	//	Contains the date and time that the parent element was last modified.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Modified string
	//	Contains a list of words used as search criteria.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Keywords string
	//	Contains revision information.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Revision string
	//	Contains a description of the topical subject.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Subject string
	//	Contains title information.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Title string
	//	Contains descriptive information about the coordinate system of the geometric data. All
	//	coordinates are right-handed by definition. Valid values are "X", "Y" (the default), or "Z".
	UpAxis string
	//	The unit of distance that applies to all spatial measurements within the scope of this resource.
	Unit struct {
		//	How many real-world meters in one distance unit as a floating-point number. 1.0 for meter, 0.01 for centimeter, 1000 for kilometer etc.
		Meter float64
		//	Name of the distance unit, such as "centimeter", "kilometer", "meter", "inch". Default is "meter".
		Name string
	}
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Contributors []*AssetContributor
	//	Provides information about the location of the visual scene in physical space.
	//	This is only set-and-retained for imported Collada assets that provide this field, and is not otherwise used.
	Coverage *AssetGeographicLocation
}

//	Constructor
func NewAsset() (me *Asset) {
	me = &Asset{}
	me.Unit.Meter, me.Unit.Name = 1, "meter"
	return
}

//	Defines authoring information for asset management.
//	ALL fields are only set-and-retained for imported Collada assets that provide those fields, and are not otherwise written, read or used.
type AssetContributor struct {
	Author        string
	AuthorEmail   string
	AuthorWebsite string
	AuthoringTool string
	Comments      string
	Copyright     string
	SourceData    string
}

//	Provides information about the location of the visual scene in physical space.
//	ALL fields are only set-and-retained for imported Collada assets that provide those fields, and are not otherwise written, read or used.
type AssetGeographicLocation struct {
	Longitude        float64
	Latitude         float64
	Altitude         float64
	AltitudeAbsolute bool
}

//	Provides arbitrary additional information about or related to its parent resource.
type Extra struct {
	//	Unique identifier
	HasId
	//	Pretty-print name/title
	HasName
	//	Resource-specific asset-management meta-data
	HasAsset
	//	Custom-technique/foreign-profile data.
	HasTechniques
	//	A hint as to the type of information that this particular Extra represents.
	Type string
}

//	Used in various geometry primitives and b-rep resources.
type IndexedInputs struct {
	//	Number of primitives
	Count uint64
	//	Inputs specify how to read data from Sources.
	Inputs []*InputShared
	//	Indices that describe the attributes for a number of primitives. The indices reference into the Sources that are referenced by the Inputs.
	Indices []uint64
	//	Number of sub-primitives, if used.
	Vcount []int64
}

//	Declares unshared input semantics of a data source and connects a consumer to that source.
type Input struct {
	//	The user-defined meaning of the input connection.
	Semantic string
	//	The location of the data source.
	Source string
}

//	Declares shared input semantics of a data source and connects a consumer to that source.
type InputShared struct {
	//	Semantic and Source
	Input
	//	The offset into the list of indices.
	Offset uint64
	//	Which inputs to group as a single set. This is helpful when multiple inputs share the same semantics.
	Set *uint64
}

//	Allows simple association of resources with custom named layers.
type Layers map[string]bool

//	Binds a specific material to a piece of geometry, binding varying and uniform parameters at the same time.
type MaterialBinding struct {
	//	Custom-technique/foreign-profile meta-data.
	HasExtras
	//	Custom-technique/foreign-profile data.
	HasTechniques
	//	Targets for animation
	Params []*Param
	//	Common-technique profile.
	TC struct {
		//	References to the materials included in this material binding.
		Materials []*FxMaterialInst
	}
}

//	Declares parametric information for its parent resource.
type Param struct {
	//	Pretty-print name/title
	HasName
	//	Scoped identifier
	HasSid
	//	The user-defined meaning of the parameter.
	Semantic string
	//	The type of the value data. This text string must be understood by the application.
	Type string
}

//	Declares a new parameter for its parent resource, and assigns it an initial value.
type ParamDef struct {
	//	Scoped identifier
	HasSid
	//	Initial value for this parameter
	Value interface{}
}

//	A hash-table containing parameter declarations of this resource.
type ParamDefs map[string]*ParamDef

//	Assigns a new value to a previously defined parameter.
type ParamInst struct {
	//	References the identifier of the pre-defined parameter (ParamDef) that will have its value set.
	Ref string
	//	Indicates if the Value is a string referencing the identifier of a connected parameter.
	IsConnectParamRef bool
	//	The new value for the referenced parameter.
	Value interface{}
}

//	References a resource by its unique identifier (Id).
type RefId string

//	Returns its current value.
func (me RefId) S() string {
	return string(me)
}

//	Modifies its current value.
func (me *RefId) Set(v string) {
	*me = RefId(v)
}

//	References a resource by its scoped identifier (Sid).
type RefSid string

//	Returns its current value.
func (me RefSid) S() string {
	return string(me)
}

//	Modifies its current value.
func (me *RefSid) Set(v string) {
	*me = RefSid(v)
}

//	Declares platform-specific or program-specific information used to process some portion of the content.
type Technique struct {
	//	The type of profile. This is a vendor-defined character string that indicates the platform or capability target for the technique.
	Profile string
	//	Arbitrary content or meta-data for this Technique.
	Data []*xmlx.Node
}

//	Represents a single transformation of a specific kind.
type Transform struct {
	//	Scoped identifier
	HasSid
	//	The type of this transformation (rotation, skewing, scaling, translation, "look-at", or matrix).
	//	The only valid values are the TRANSFORM_TYPE_* enumerated constants.
	Type int
	//	Contains one or more vectors and values representing this transformation.
	//	If type is "look-at", contains 9 values representing 3 vectors (eye position, interest point, up-axis).
	//	If type is matrix, contains 16 values representing a column-major 4x4 matrix.
	//	If type is skew, contains 7 values: one angle in degrees, then 2 vectors specifying the axes of rotation and translation.
	//	If type is rotate, contains 4 values: one vector specifying the axis of rotation, then 1 value for the angle in degrees.
	//	If type is translate or scale, contains 3 values representing a single column vector.
	F []float64
}
