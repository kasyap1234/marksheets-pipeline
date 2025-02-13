package handlers

import (
	"fmt"
	"io"
	"marksheets-pipeline/internal/api"
	"marksheets-pipeline/internal/logging"
	"net/http"
	"strings"
)

type APIHandler struct {
	pythonProcessorClient *client.PythonProcesorClient 
	logger logging.Logger 
}

func NewAPIHandler(pythonProcessorClient *client.PythonProcessorClient,logger logging.Logger)*APIHandler{
	return &APIHandler{
		pythonProcessorClient : pythonProcessorClient, 
		logger : logger ,
	}

}

func (h*APIHandler)ExtractMarksheetsInfo(w http.ResponseWriter, r*http.Request){
	if r.Method !=http.MethodPost{
		api.RespondWithError(w,r,http.StatusMethodNotAllowed,"Method Not Allowed",h.logger)
		return 
	}

	if err :=r.ParseMultipartForm(10 << 20); err !=nil {
		api.RespondWithError(w,r,http.StatusInternalServerError,"Error parsing multipart form",h.logger)
		http.Error(w,"unable to parse multipart form",http.StatusInternalServerError)
	} // left shift and 10MB limit 
	file,fileHeader,err :=r.FormFile("file")
	if err !=nil {
		http.Error(w,"Unable to parse file",http.StatusBadRequest)
		return 
	}
	defer file.Close()
	contentType := fileHeader.Header.Get("Content-Type")
	allowedTypes :=[]string{"image/jpg","image/png","image/jpeg"}
	isValidType :=false
	for _,allowedTypes :=range allowedTypes{
		if strings.ToLower(contentType)==strings.ToLower(allowedTypes){
				isValidType=true 
		break 
			}
			if !isValidType{
				 api.RespondWithError(w,r,http.StatusBadRequest,"Invalid file type",h.logger)
return 
			}
			imageBytes,err :=io.ReadAll(file)
			if err !=nil {
				http.Error(w,"Unable to parse bytes",http.StatusInternalServerError)

			return 
			}
			pythonResponse, err := h.pythonProcessorClient.ProcessImage(r.Context(),imageBytes,contentType)
			if err !=nil {
				api.RespondWithError(w,r,http.StatusInternalServerError,"invalid response",h.logger)

			}
			response :=api.ExtractionResponse{
				Filename: fileHeader.Filename,
				ContentType: contentType,
				ExtractedInformation: pythonResponse,
				ProcessingStatus: "success",
			}
			api.RespondWithJSON(w,r,http.StatusOK,response,h.logger)

		}



	


}