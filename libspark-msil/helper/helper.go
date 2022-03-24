package helper

import (
	"github.com/rs/xid"
	"net/url"
	"strings"
	"libspark-msil/database/postgresql"
	"libspark-msil/constants"
	"libspark-msil/entities"
	"log"
	"time"
	"fmt"
	"context"

)

func GenerateUniqueID () string{
	uuid := xid.New()
	return uuid.String()
}
var bucket string
var key string

func S3URLtoURI(s3Url string)(string ,string)  {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Exception while parsing S3 url: %v \n", r)
		}
	}()


	u, err := url.Parse(s3Url)
	if err != nil {
		panic("issue occurring on S3URLtoURI function")
	}

	if u.Scheme == "s3" {
		//s3: //bucket/key
		bucket = u.Host
		key = strings.TrimLeft(u.Path, "/")

	} else if u.Scheme == "https" {
		host := strings.SplitN(u.Host, ".", 2)
		if host[0] == "s3" {
			// No bucket name in the host;
			path := strings.SplitN(u.Path, "/", 3)
			bucket = path[1]
			key = path[2]

		} else { //bucket name in host
			bucket = host[0]
			key = strings.TrimLeft(u.Path, "/")
		}

	}
	log.Println("buc",bucket)
	log.Println("key",key)

	return bucket ,key

}


//function to delete from user nudge master table

func DeleteUserNudgeMasterTble (sQuery string, sNudgeID string){

	log.Println("Started Deleting")

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Exception in DeleteUserNudgeMasterTble : %v \n", r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), constants.Timeout)
	defer cancel()

	rows := database.GetResultsExec(ctx, sQuery,sNudgeID)

	log.Println("Record deleted from user nudgemaster..!" , rows.RowsAffected())
	log.Println("Finished Deleting")

}

func VerifyTime(endTime string) bool {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Exception in VerifyTime : %v \n", r)
		}
	}()


	loc, _ := time.LoadLocation(constants.Location)

	log.Println(endTime)

	endDate, err := time.ParseInLocation(constants.TimeFormat, endTime , loc)

	if err != nil {
		log.Println("Error parsing end Time")
	}

	today := time.Now()
	flag := today.Before(endDate)
	log.Println(endDate,"Today:",today)
	log.Println(flag)
	return flag
}


func MsgFormat(field string, tag string) string{
	switch tag {
	case "required":
		return field + " is required"
	}
	return ""
}

func GetFrequency(freq string)string{
	var res string

	if strings.EqualFold(freq, constants.OncePerDayWeb) {
		res = constants.OncePerDay
	}else if strings.EqualFold(freq, constants.OncePerWeekWeb) {
		res = constants.OncePerWeek
	}else {
		res = constants.Default
	}

	return res
}

func AddCTAArr(Contents map[string]interface{})  []entities.CTAResponse {
	CTAArr:= []entities.CTAResponse{}
	for k,v := range Contents{
		if k == constants.CTAWeb {
			if(v == nil){
				log.Println("CTA Array is NULL")
				//      c.IndentedJSON(http.StatusOK,gin.H{"Status":"Not Alive"})
				break
			}
			CTA := v.([]interface{})
			for _, cta_v := range CTA {
				CTAResp :=entities.CTAResponse{}
				CTAResp.Text = cta_v.(map[string]interface{})[constants.Title]
				CTAResp.Typ = cta_v.(map[string]interface{})[constants.Type]
				CTAResp.Link = cta_v.(map[string]interface{})[constants.Link]
				CTAResp.CTA_Type = cta_v.(map[string]interface{})[constants.CTAType]
				CTAResp.Label = cta_v.(map[string]interface{})[constants.Label]
				CTAResp.IsEnabled = true

				CTAArr = append(CTAArr, CTAResp)
			}
		}

	}


	return  CTAArr
}


func AddCommonNudgeDetails(values []interface{}) (map[string]interface{}){
	Contents := values[3].(map[string]interface{})

	CTAArr:=AddCTAArr(Contents)

	Contents[constants.PlaceHolderId] = values[1]
	Contents[constants.Type] = values[2]
	Contents[constants.Start] = values[4]
	Contents[constants.End] = values[5]

	freq := fmt.Sprint(values[6])
	freq = GetFrequency(freq)
	Contents[constants.Frequency] = freq

	Contents[constants.HowOften] = values[6]
	Contents[constants.NudgeId] = values[7]
	Contents[constants.HeadLine] = Contents[constants.HeadLineWeb]
	Contents[constants.Desc] = Contents[constants.CardLines]
	Contents[constants.Image] = Contents[constants.BgImage]
	Contents[constants.CTA] = CTAArr
	Contents[constants.BgColor] = Contents[constants.BgColorWeb]
	delete(Contents,constants.BgImage)
	delete(Contents,constants.CardLines)
	delete(Contents,constants.HeadLineWeb)
	delete(Contents,constants.CTAWeb)
	delete(Contents,constants.BgColorWeb)

	return Contents
}


