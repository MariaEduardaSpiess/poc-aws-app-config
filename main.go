package main

import (
	"net/http"
	"regexp"

	"github.com/MariaEduardaSpiess/aws_app_config_test/feature_flags"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma nova instância do Gin
	r := gin.Default()

	// Define uma rota para o endpoint "/"
	r.GET("/:cpf", func(c *gin.Context) {
		cpf := c.Param("cpf")
		flags, err := feature_flags.GetFeatureFlags()
		if err != nil {
			c.Error(err)
		}

		// Valida se o a flag está ativa e se o CPF é válido pro teste
		if flags.FeatureXpto.Enabled {
			cpfPattern := regexp.MustCompile(flags.FeatureXpto.RegexCpfsRollout)
			if cpfPattern.MatchString(cpf) {
				c.JSON(http.StatusOK, gin.H{"featureEnabled": true})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"featureEnabled": false})
	})

	// Inicia o servidor
	r.Run()
}
