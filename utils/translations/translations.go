package translations

var ErrorMessages = map[any]map[string]string{
	"BadRequest": {
		"en-US": "Bad request",
		"pt-BR": "Requisição inválida",
	},
	"InternalServer": {
		"en-US": "Internal server error",
		"pt-BR": "Erro interno do servidor",
	},
	"MethodNotAllowed": {
		"en-US": "Method not allowed",
		"pt-BR": "Método não permitido",
	},
	"NotFound": {
		"en-US": "Not found",
		"pt-BR": "Não encontrado",
	},
	"FailedToGenerateUUID": {
		"en-US": "Failed to generate UUID",
		"pt-BR": "Falha ao gerar UUID",
	},
	"FailedToEncryptPassword": {
		"en-US": "Failed to encrypt password",
		"pt-BR": "Falha ao criptografar senha",
	},
	"FailedToReadUsersFile": {
		"en-US": "Failed to read users file",
		"pt-BR": "Falha ao ler o arquivo de usuários",
	},
	"FailedToWriteUsersFile": {
		"en-US": "Failed to write to users file",
		"pt-BR": "Falha ao escrever no arquivo de usuários",
	},
	"InvalidUserID": {
		"en-US": "Invalid user ID",
		"pt-BR": "ID de usuário inválido",
	},
	"UserNotFound": {
		"en-US": "User not found",
		"pt-BR": "Usuário não encontrado",
	},
}
