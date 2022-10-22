package platformcontracts

func GetRequestConsumerName(platformName string) string {
	return platformName + PLATFORM_URL_REQUESTS_CONSUMER_NAME_TEMPLATE
}

func GetRequestStreamSubject(platformName string) string {
	return PLATFORM_URL_REQUESTS_STREAM_SUBJECT_TEMPLATE + platformName
}

const PLATFORM_URL_REQUESTS_STREAM_NAME = "PLATFORM-URL-REQUESTS"
const PLATFORM_URL_REQUESTS_STREAM_SUBJECT_TEMPLATE = "PLATFORM-URL-REQUESTS."
const PLATFORM_URL_REQUESTS_STREAM_SUBJECTS = "PLATFORM-URL-REQUESTS.*"
const PLATFORM_URL_REQUESTS_CONSUMER_NAME_TEMPLATE = "-platform-request-consumer"

const PLATFORM_URL_RESPONSE_STREAM_NAME = "PLATFORM-URL-RESPONSES"
const PLATFORM_URL_RESPONSE_STREAM_SUBJECT = "PLATFORM-URL-RESPONSES.app"
const PLATFORM_URL_RESPONSE_CONSUMER_NAME = "app-platform-response-consumer"
