package handlers

// @Summary Adiciona item à fila
// @Description Adiciona um item a uma fila nomeada. Cria a fila se não existir.
// @Tags fila
// @Accept json
// @Produce plain
// @Param data body EnqueueRequest true "Dados do enqueue"
// @Success 201 {string} string "adicionado"
// @Failure 400 {string} string "JSON inválido ou campos obrigatórios ausentes"
// @Router /enqueue [post]
func _Swagger_EnqueueHandler() {}

// @Summary Remove item da fila
// @Description Remove e retorna o próximo item de uma fila nomeada.
// @Tags fila
// @Accept json
// @Produce json
// @Param data body DequeueRequest true "Dados do dequeue"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "JSON inválido ou campo 'queue' ausente"
// @Failure 404 {string} string "Fila não encontrada ou vazia"
// @Router /dequeue [post]
func _Swagger_DequeueHandler() {}

// @Summary Health check
// @Description Verifica se o serviço está online.
// @Tags health
// @Success 200 {string} string "ok"
// @Router /health [get]
func _Swagger_HealthCheckHandler() {}
