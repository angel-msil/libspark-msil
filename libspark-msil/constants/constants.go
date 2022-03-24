package constants
import "time"

// config names
const (
	DatabaseConfig          = "database"
	ApplicationConfig       = "application"
	LoggerConfig            = "logger"
	LogLevelConfigKey       = "level"
)

const(
	Listen          = "listen"
)
//database
const(
	Host="host"
	Port = "port"
	User = "user"
	Password = "password"
	Database = "database"
)
//Error Code
const(
	KeyNotFound = "ELOO1"
	Exception = "ELOO2"
)
//common constants
const(
	ErrorCode = "errorCode"
	ErrorMessage = "errorMessage"
	Message = "message"
	Status = "status"
	UnableToInitializeURL =  "unable to initialize presigned url"
	EndDateVerificationMessage = "EndTime should be greater than current time"
	StatusCode = "statusCode"
	Msg = "msg"
	SuccessResponse = "success"
	Body = "body"
	Location = "Asia/Kolkata"
	TimeFormat = "2006-01-02T15:04"
)
//Device Values

const(
        SparkMobile = "spark_mobile"
        OncePerDayWeb = "once_per_day"
        OncePerWeekWeb = "once_per_week"
)
//freq constants

const (
        OncePerDay = "1"
        OncePerWeek = "2"
        Default = "0"
)


//http constants
const(
	FailureCode = "502"
	SuccessCode = "200"
)

//kafka constants

const(
	KakfaServer = "localhost:9092"
	Channel = "channel"
	RequestReceived = "Request"
	ResponseSend = "Response"
	DB_Error = "DynamoDB error"
	UnmarshalFailed = "Unmarshal Failed"
	MarshalFailed = "Marshal Failed"
	DBResult = "DynamoDB Result"
	DefaultValues = "Default Values"
	EmptyRes = "Empty Response from DynamoDB"
	KafkaCloseProducerErrorMsg  = "Kakfa: No asyncproducer available"
	KafkaProducerErrorMsg  = "Kafka: Error sending to kafka, kafka produce error"
	PostedToKafka  = "Successfully posted to kafka"
	KafkaConnectErrorMsg = "Kafka: Error initializing asyncProducer"
	KafkaConnect  = "Initialized kafka asyncproducer"
	BadRequest="Bad Request"
        DBUpdationError="DB Updation Error"
        DateVarification="Date Verification"
        UnableToInitialize="Initialise URL Error"
        ExceptionError="Exception case"
)
//timeout

const(
	Timeout = 100 * time.Millisecond
)
//Device Key
const(
	CTAWeb = "cta"
	CTA = "CTA"
        Title = "title"
        Type = "typ"
        Link = "link"
        PlaceHolderId = "plcHoldrID"
        Start = "start"
        End = "end"
        Frequency = "freq"
        NudgeId = "nudgeId"
        HeadLine = "hdlne"
        HeadLineWeb = "headline"
        Desc = "desc"
        CardLines = "cardLines"
        Image = "img"
        BgImage = "bgImage"
        BgColor = "bgC"
        BgColorWeb = "bgColor"
	Label = "label"
        HowOften = "howOften"
        CTAType = "cta_type"
)
