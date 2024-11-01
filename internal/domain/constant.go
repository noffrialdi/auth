package constants

const (
	EmptyString = ""
)

const (
	CallTypeAPI   = "API"
	CallTypeKafka = "KAFKA"

	Controller         = "controller"
	Usecase            = "usecase"
	RepositoryMariaDB  = "repository.mariadb"
	RepositoryRedis    = "repository.redis"
	RepositoryBag      = "repository.bag"
	RepositoryDupsvc   = "repository.dupsvc"
	RepositoryRedshift = "repository.redshift"
)

const (
	CrudService         = "crud-service"
	CrudServiceConsumer = "crud-svc-consumer"
)

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

const (
	mariadb = "mariadb"
)
