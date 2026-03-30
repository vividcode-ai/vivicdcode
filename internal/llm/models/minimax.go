package models

const (
	MiniMaxM2_7 ModelID = "minimax-m2.7"
	MiniMaxM2_5 ModelID = "minimax-m2.5"
)

var MiniMaxModels = map[ModelID]Model{
	MiniMaxM2_7: {
		ID:                  MiniMaxM2_7,
		Name:                "MiniMax M2.7",
		Provider:            ProviderMiniMax,
		APIModel:            "MiniMax-M2.7",
		CostPer1MIn:         1.0,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        10.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: false,
	},
	MiniMaxM2_5: {
		ID:                  MiniMaxM2_5,
		Name:                "MiniMax M2.5",
		Provider:            ProviderMiniMax,
		APIModel:            "MiniMax-M2.5",
		CostPer1MIn:         0.8,
		CostPer1MInCached:   0,
		CostPer1MOutCached:  0,
		CostPer1MOut:        8.0,
		ContextWindow:       1_000_000,
		DefaultMaxTokens:    32_768,
		SupportsAttachments: false,
	},
}
