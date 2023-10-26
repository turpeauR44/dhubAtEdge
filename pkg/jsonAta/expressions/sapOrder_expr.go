package jsonAtaTransform

func SAP_Order()(string){
	return `$.data.{
	"operationRequest":{
	"id":manufacturingOrder,
"requestState":"ToBemapped",
	"scheduledStartDate": mfgOrderScheduledStartDate &"T"& mfgOrderScheduledStartTime,
	"scheduledEndDate": mfgOrderScheduledEndDate &"T"& mfgOrderScheduledEndTime,
	"operationsType":"ToBemapped",
	"segmentRequirment":{
		"id":manufacturingOrder,
		"earliestStartDate":"ToBeDefined",
		"materialRequirements":{
			"id":"SR."& manufacturingOrder &"."& productionOrderItems.material &".0?",
			"quantity":mfgOrderPlannedTotalQty,
			"materialLots":{
				"materialDefinition":{
					"id":material,
					"label":materialName
				}
			}
		}
	}
}
`
}