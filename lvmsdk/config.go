package lvmsdk

type Config struct {
	Addr             string
	Token            string
	Header           Header
	Parallelism      int
	MemoryLimitPages uint32
}
