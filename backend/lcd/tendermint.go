package lcd

import (
	"encoding/json"
	"fmt"

	"github.com/irisnet/explorer/backend/conf"
	"github.com/irisnet/explorer/backend/logger"
	"github.com/irisnet/explorer/backend/utils"
)

func NodeInfo() (result NodeInfoVo, err error) {
	url := fmt.Sprintf(UrlNodeInfo, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}

func Genesis() (result GenesisVo, err error) {
	url := fmt.Sprintf(UrlGenesis, conf.Get().Hub.NodeUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("get account error", logger.String("err", err.Error()))
		return result, err
	}
	return result, nil
}

func GetGenesisGovModuleParamMap() (map[string]interface{}, error) {
	url := fmt.Sprintf(UrlGenesis, conf.Get().Hub.NodeUrl)
	resBytes, err := utils.Get(url)

	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(resBytes), &m)
	if err != nil {
		return nil, err
	}

	resultMap := m["result"].(map[string]interface{})
	genesisMap := resultMap["genesis"].(map[string]interface{})
	appStateMap := genesisMap["app_state"].(map[string]interface{})

	authMap := appStateMap[GovModuleAuth].(map[string]interface{})
	authParamMap := authMap["params"].(map[string]interface{})

	stakeMap := appStateMap[GovModuleStake].(map[string]interface{})
	stakeParamMap := stakeMap["params"].(map[string]interface{})

	mintMap := appStateMap[GovModuleMint].(map[string]interface{})
	mintParamMap := mintMap["params"].(map[string]interface{})

	distrMap := appStateMap[GovModuleDistr].(map[string]interface{})
	distrParamMap := distrMap["params"].(map[string]interface{})

	slashingMap := appStateMap[GovModuleSlashing].(map[string]interface{})
	slashingParamMap := slashingMap["params"].(map[string]interface{})

	for k, v := range distrParamMap {
		slashingParamMap[k] = v
	}

	for k, v := range mintParamMap {
		slashingParamMap[k] = v
	}

	for k, v := range stakeParamMap {
		slashingParamMap[k] = v
	}
	for k, v := range authParamMap {
		slashingParamMap[k] = v
	}

	return slashingParamMap, nil
}

func Block(height int64) (result BlockVo) {
	url := fmt.Sprintf(UrlBlock, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Block error", logger.Int64("height", height))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Block error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func BlockLatest() (result BlockVo) {
	url := fmt.Sprintf(UrlBlockLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func ValidatorSet(height int64) (result ValidatorSetVo) {
	url := fmt.Sprintf(UrlValidatorSet, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func LatestValidatorSet() (result ValidatorSetVo) {
	url := fmt.Sprintf(UrlValidatorSetLatest, conf.Get().Hub.LcdUrl)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockLatest error", logger.String("err", err.Error()))
		return result
	}
	return result
}

func BlockResult(height int64) (result BlockResultVo) {

	url := fmt.Sprintf(UrlBlocksResult, conf.Get().Hub.LcdUrl, height)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("BlockResult error", logger.String("err", err.Error()))
		return result
	}

	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("BlockResult error", logger.String("err", err.Error()))
		return result
	}
	return result

}
