package admin

// tokenize finds the columns to be tokenized and tokenizes their values
//
// Example:
//
//	config.Tokenized = []string{"first_name", "last_name", "email"}
//
//	inputMap = {
//		"first_name": "John",
//		"last_name":  "Doe",
//		"status":     "active",
//		"email":      "4l8Hw@example.com",
//	}
//
//	outputMap = {
//		"first_name": "tk_a1h1t2h3o4n5",
//		"last_name":  "tk_z7d8e9f0g1h2",
//		"status":     "active",
//		"email":      "tk_4l8sUzs2h3s4",
//	}
// func tokenize(config shared.Config, inputMap map[string]string) (outputMap map[string]string, err error) {
// 	// are there tokenizable columns to tokenize?
// 	if len(config.Tokenized) < 1 {
// 		return inputMap, nil
// 	}

// 	columnsToTokenize := lo.Intersect(lo.Keys(inputMap), config.Tokenized)

// 	// create a subset map containing only the columns to be tokenized
// 	mapToTokenize := make(map[string]string)

// 	for _, key := range columnsToTokenize {
// 		mapToTokenize[key] = inputMap[key]
// 	}

// 	tokenizedMap, err := config.Tokenize(mapToTokenize)

// 	if err != nil {
// 		return map[string]string{}, err
// 	}

// 	// replace the original values with the tokenized values
// 	outputMap = inputMap

// 	for k, v := range tokenizedMap {
// 		outputMap[k] = v
// 	}

// 	return outputMap, nil
// }

// untokenize finds the columns to be untokenized and untokenizes their values
//
// Example:
//
//	config.Tokenized = []string{"first_name", "last_name", "email"}
//
//	inputMap = {
//		"first_name": "tk_a1h1t2h3o4n5",
//		"last_name":  "tk_z7d8e9f0g1h2",
//		"status":     "active",
//		"email":      "tk_4l8sUzs2h3s4",
//	}
//
//	outputMap = {
//		"first_name": "John",
//		"last_name":  "Doe",
//		"status":     "active",
//		"email":      "4l8Hw@example.com",
//	}
// func untokenize(config shared.Config, inputMap map[string]string) (outputMap map[string]string, err error) {
// 	// are there tokenizable columns to untokenize?
// 	if len(config.Tokenized) < 1 {
// 		return inputMap, nil
// 	}

// 	columnsToUntokenize := lo.Intersect(lo.Keys(inputMap), config.Tokenized)

// 	// create a subset map containing only the columns to be untokenized
// 	mapToUntokenize := make(map[string]string)

// 	for _, key := range columnsToUntokenize {
// 		mapToUntokenize[key] = inputMap[key]
// 	}

// 	untokenizedMap, err := config.Untokenize(mapToUntokenize)

// 	if err != nil {
// 		return map[string]string{}, err
// 	}

// 	// replace the original values with the untokenized values
// 	outputMap = inputMap

// 	for k, v := range untokenizedMap {
// 		outputMap[k] = v
// 	}

// 	return outputMap, nil
// }

// tokensUpdate updates the tokens in the vault
//
// Example:
//
//	tokenValueMap = {
//		"tk_a1h1t2h3o4n5": "Jane",
//		"tk_z7d8e9f0g1h2": "Air",
//		"tk_4l8sUzs2h3s4": "jane@example.com",
//	}
// func tokensUpdate(config shared.Config, tokenValueMap map[string]string) (err error) {
// 	for k, v := range tokenValueMap {
// 		err = config.TokenUpdate(k, v)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
