package common

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"errors"
	"Libraries/libspark-msil/constants"
	"Libraries/libspark-msil/kafka"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

func SendBadRequest(c *gin.Context,err error,messageID string){

	log.Println("send bad request : ",err)
	log.Println("--------------", err, "---------------")

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var out string
		for _, fe := range ve {
			out = MsgFormat(fe.Field(),fe.Tag())
		}
		c.JSON(http.StatusOK, gin.H{constants.ErrorCode:constants.KeyNotFound,constants.Message:out})
		//		go kafka.SendLogsToKafka(constants.BadRequest , messageID ,c.Request.URL.Path , gin.H{constants.ErrorCode:constants.KeyNotFound, constants.Message:out})

	}
}

func SendDBUpdationError(c *gin.Context, msg string,messageID string){

	c.JSON(http.StatusOK, gin.H{
		constants.Status:  http.StatusInternalServerError,
		constants.Message: msg,
	})
	//	go kafka.SendLogsToKafka(constants.DBUpdationError , messageID ,c.Request.URL.Path , gin.H{constants.Status:  http.StatusInternalServerError,constants.Message: msg})

}

func SendDateVarificationMsg(c *gin.Context,messageID string){

	c.JSON(http.StatusOK, gin.H{
		constants.Status: http.StatusOK,
		constants.Message: constants.EndDateVerificationMessage,
	})
	//	go kafka.SendLogsToKafka(constants.DateVarification , messageID ,c.Request.URL.Path , gin.H{constants.Status:  http.StatusOK, constants.Message:constants.EndDateVerificationMessage})

}

func SendUnableToInitializeURL(c *gin.Context,err error,messageID string){

	log.Println("unable to initialize presigned url",err)

	c.JSON(http.StatusOK, gin.H{
		constants.StatusCode : constants.FailureCode,
		constants.Msg  : constants.UnableToInitializeURL,
	})
	//	go kafka.SendLogsToKafka(constants.UnableToInitialize , messageID ,c.Request.URL.Path , gin.H{ constants.StatusCode : constants.FailureCode,constants.Msg  : constants.UnableToInitializeURL})


}

func MsgFormat(field string, tag string) string{
	switch tag {
	case "required":
		return field + " is required"
	}
	return ""
}

func ExceptionHandling(c *gin.Context,messageID string){

	if r := recover(); r != nil {
		log.Printf("Exception: %v \n", r)
		msg := fmt.Sprintf("%v", r)
		c.IndentedJSON(http.StatusOK,gin.H{
			constants.ErrorCode: constants.Exception,
			constants.ErrorMessage: msg,
		})
		//		go kafka.SendLogsToKafka(constants.ExceptionError , messageID ,c.Request.URL.Path , gin.H{ constants.ErrorCode: constants.Exception, constants.ErrorMessage: msg})

	}
}




/////////////////////////////////////////////////////////////////////////////////////////////////



func SendOperationError(c *gin.Context,topic string, msg string,messageID string, err error){

	c.IndentedJSON(http.StatusBadRequest,gin.H{constants.Status:http.StatusBadRequest,constants.Message:msg})
	fmt.Println("Got error calling Operation: %s", err)
	err_msg := fmt.Sprintf("%v", err)
	go kafka.SendLogsToKafka(topic,constants.DB_Error , messageID ,c.Request.URL.Path , gin.H{constants.Status:http.StatusBadRequest,constants.Message:err_msg})
}


func UnableToUnmarshal(c *gin.Context, topic string,msg string,messageID string, err error){
	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	err_msg := fmt.Sprintf("%v", err)
	c.IndentedJSON(http.StatusBadRequest,gin.H{constants.Status:http.StatusBadRequest,constants.Message:msg})
	go kafka.SendLogsToKafka(topic,constants.UnmarshalFailed , messageID ,c.Request.URL.Path , gin.H{constants.Status:  http.StatusBadRequest,constants.Message: err_msg})
}


func UnableToMarshal(c *gin.Context, topic string,msg string,messageID string, err error){
	log.Println("Got error marshalling new item: %s", err)
	err_msg := fmt.Sprintf("%v", err)

	go kafka.SendLogsToKafka(topic,msg, messageID ,c.Request.URL.Path , gin.H{constants.Status:  http.StatusBadRequest,constants.Message: err_msg})

}
