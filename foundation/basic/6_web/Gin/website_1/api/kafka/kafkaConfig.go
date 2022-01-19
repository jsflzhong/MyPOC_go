package kafka

func GetBrokers() []string {
	// return os.Getenv("CLOUDKARAFKA_BROKERS")
	return []string{
		"sulky-01.srvs.cloudkafka.com:9094",
		"sulky-02.srvs.cloudkafka.com:9094",
		"sulky-03.srvs.cloudkafka.com:9094"}
}

func GetBrokersString() string {
	// return os.Getenv("CLOUDKARAFKA_BROKERS")
	return "sulky-01.srvs.cloudkafka.com:9094,sulky-02.srvs.cloudkafka.com:9094,sulky-03.srvs.cloudkafka.com:9094"
}

func GetUsername() string {
	//return  os.Getenv("CLOUDKARAFKA_USERNAME")
	return "ptxqyt3u"
}

func GetPassword() string {
	//return  os.Getenv("CLOUDKARAFKA_PASSWORD")
	return ""
}

func GetGroupId() string {
	return "cloudkarafka-example"
}

func GetTopicPrefix() string {
	return "ptxqyt3u-"
}
