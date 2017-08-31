package pipeline

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"
)

func loadTrainCsv(dirName string) ([][]string, error) {
	// load training data from the supplied directory
	f, err := os.Open(path.Join(dirName, "trainData.csv"))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func writeResultCsv(resultPath string, data [][]string) error {
	// create result directory if necessary
	err := os.MkdirAll(path.Dir(resultPath), 0700)
	if err != nil {
		return err
	}

	// create the result file
	file, err := os.Create(resultPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// write teh result into it and close
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateResultCsv(
	pipelineID string,
	seqNum int,
	trainDirName string,
	resultDirName string,
	targetFeature string,
	resultGenerator func() string,
) (string, error) {
	// load training data - just use it to get count for now
	records, err := loadTrainCsv(trainDirName)
	if err != nil {
		return "", err
	}

	// generate mock results
	result := [][]string{{"d3m_index", targetFeature}}
	for i := 0; i < len(records); i++ {
		result = append(result, []string{strconv.Itoa(i), resultGenerator()})
	}

	// write results out to disk
	path := path.Join(resultDirName, fmt.Sprintf("%s-%d.csv", pipelineID, seqNum))
	return path, writeResultCsv(path, result)
}
