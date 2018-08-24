package core

type Data struct {
	class string
	sentence string
}

func CreateDataTraining() *[]Data {
	training_data := make([]Data, 0)
	training_data = append(training_data, Data{class:"greeting", sentence:"how are you?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"how is your day?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"how is it going today?"})
	training_data = append(training_data, Data{class:"greeting", sentence:"good day"})

	training_data = append(training_data, Data{class:"goodbye", sentence:"have a nice day"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"goodbye"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"have a nice day"})
	training_data = append(training_data, Data{class:"goodbye", sentence:"talk to you soon"})
	return &training_data
}