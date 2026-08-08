package mongopkg

const (
	// Comparison query operators.
	OperationEq  = "$eq"
	OperationGt  = "$gt"
	OperationGte = "$gte"
	OperationIn  = "$in"
	OperationLt  = "$lt"
	OperationLte = "$lte"
	OperationNe  = "$ne"
	OperationNin = "$nin"

	// Logical query operators.
	OperationAnd = "$and"
	OperationNot = "$not"
	OperationNor = "$nor"
	OperationOr  = "$or"

	// Element and evaluation query operators.
	OperationExists     = "$exists"
	OperationType       = "$type"
	OperationExpr       = "$expr"
	OperationJSONSchema = "$jsonSchema"
	OperationMod        = "$mod"
	OperationRegex      = "$regex"
	OperationText       = "$text"

	// Array query operators.
	OperationAll       = "$all"
	OperationElemMatch = "$elemMatch"
	OperationSize      = "$size"

	// Bitwise query operators.
	OperationBitsAllClear = "$bitsAllClear"
	OperationBitsAllSet   = "$bitsAllSet"
	OperationBitsAnyClear = "$bitsAnyClear"
	OperationBitsAnySet   = "$bitsAnySet"

	// Geospatial query operators.
	OperationGeoIntersects = "$geoIntersects"
	OperationGeoWithin     = "$geoWithin"
	OperationNear          = "$near"
	OperationNearSphere    = "$nearSphere"
)

const (
	// Field update operators.
	OperationCurrentDate = "$currentDate"
	OperationInc         = "$inc"
	OperationMin         = "$min"
	OperationMax         = "$max"
	OperationMul         = "$mul"
	OperationRename      = "$rename"
	OperationSet         = "$set"
	OperationSetOnInsert = "$setOnInsert"
	OperationUnset       = "$unset"

	// Array update operators.
	OperationAddToSet = "$addToSet"
	OperationPop      = "$pop"
	OperationPull     = "$pull"
	OperationPullAll  = "$pullAll"
	OperationPush     = "$push"
	OperationEach     = "$each"
	OperationPosition = "$position"
	OperationSlice    = "$slice"

	// Bitwise update operator.
	OperationBit = "$bit"
)

const (
	// Aggregation pipeline stages.
	OperationAddFields       = "$addFields"
	OperationBucket          = "$bucket"
	OperationBucketAuto      = "$bucketAuto"
	OperationChangeStream    = "$changeStream"
	OperationCollStats       = "$collStats"
	OperationCount           = "$count"
	OperationDensify         = "$densify"
	OperationFacet           = "$facet"
	OperationFill            = "$fill"
	OperationGeoNear         = "$geoNear"
	OperationGraphLookup     = "$graphLookup"
	OperationMatch           = "$match"
	OperationGroup           = "$group"
	OperationIndexStats      = "$indexStats"
	OperationLimit           = "$limit"
	OperationLookup          = "$lookup"
	OperationMerge           = "$merge"
	OperationOut             = "$out"
	OperationPlanCacheStats  = "$planCacheStats"
	OperationProject         = "$project"
	OperationRedact          = "$redact"
	OperationReplaceRoot     = "$replaceRoot"
	OperationReplaceWith     = "$replaceWith"
	OperationSample          = "$sample"
	OperationSetWindowFields = "$setWindowFields"
	OperationSkip            = "$skip"
	OperationSort            = "$sort"
	OperationSortByCount     = "$sortByCount"
	OperationUnionWith       = "$unionWith"
	OperationUnwind          = "$unwind"
)

const (
	// Common aggregation expressions and accumulators.
	OperationAdd          = "$add"
	OperationAvg          = "$avg"
	OperationCond         = "$cond"
	OperationConcat       = "$concat"
	OperationConcatArrays = "$concatArrays"
	OperationDivide       = "$divide"
	OperationFilter       = "$filter"
	OperationFirst        = "$first"
	OperationIfNull       = "$ifNull"
	OperationLast         = "$last"
	OperationLiteral      = "$literal"
	OperationMap          = "$map"
	OperationMultiply     = "$multiply"
	OperationReduce       = "$reduce"
	OperationSubtract     = "$subtract"
	OperationSum          = "$sum"
	OperationSwitch       = "$switch"
	OperationToString     = "$toString"
)

const (
	// OperationNotIn is kept for compatibility.
	// Deprecated: use OperationNin.
	OperationNotIn = OperationNin

	// Common $lookup fields.
	FieldFrom         = "from"
	FieldLocalField   = "localField"
	FieldForeignField = "foreignField"
	FieldAs           = "as"
	FieldLet          = "let"
	FieldPipeline     = "pipeline"
)
