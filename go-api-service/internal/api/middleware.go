package api

import (
	"encoding/json"
	"marksheets-pipeline/internal/logging"
	"net/http"
)

type ExtractionResponse struct {
	Filename string 	`json:"filename"`
	ContentType string 	`json:"content_type"`
	ExtractedInformation map[string]interface{} `json:"extracted_information"`
	ProcessingStatus string `json:"processing_status"`
	Error string 		`json:"error"`
}


func RespondWithJSON(w http.ResponseWriter,r *http.Request,statusCode int , payload interface{},logger logging.Logger){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	err :=json.NewEncoder(w).Encode(payload)
	if err !=nil {
		logger.Errorf("Error encoding JSON response %v",err)
		http.Error(w,"Internal Server Errror",http.StatusInternalServerError)
		return 
	}

}

func RespondWithError(w http.ResponseWriter,r*http.Request,statusCode int, message string , logger logging.Logger){
	logger.Errorf("API error : Status Code : %d,Message : %s",statusCode,message)
	response :=ExtractionResponse {
		ProcessingStatus: "error",
		Error: message,
	}
	RespondWithJSON(w,r,statusCode,response,logger)


}
